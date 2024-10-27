package pkg_kit

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

var (
	nowVersion = versionDevel
	nowBuildId = infoUnknown
)

func FetchNowVersion() string {
	return nowVersion
}
func FetchNowBuildId() string {
	return nowBuildId
}

const (
	infoUnknown  = "unknown"
	versionDevel = "devel"
)

type BuildInfo struct {
	PkgName string `json:"pkgName"`

	Version    string `json:"version"`
	RawVersion string `json:"rawVersion"`
	BuildId    string `json:"buildId"`

	GoVersion    string `json:"goVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	Date         string `json:"date"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`

	AuthorName         string `json:"authorName"`
	CopyrightStartYear string `json:"copyrightStartYear"`
	CopyrightNowYear   string `json:"copyrightNowYear"`
}

func (b BuildInfo) String() string {
	return fmt.Sprintf("%s has version %s, © %s-%s %s,  built with %s id: %s from %s on %s, run on %s",
		b.PkgName, b.Version, b.CopyrightStartYear, b.CopyrightNowYear, b.AuthorName, b.GoVersion, b.BuildId, b.GitCommit, b.Date, b.Platform)
}

func (b BuildInfo) Copyright() string {
	return fmt.Sprintf("© %s-%s by: %s  build with %s id: %s, run on %s",
		b.CopyrightStartYear, b.CopyrightNowYear, b.AuthorName, b.GoVersion, b.BuildId, b.Platform)
}

func (b BuildInfo) PgkNameString() string {
	return b.PkgName
}

func (b BuildInfo) VersionString() string {
	return b.Version
}

func (b BuildInfo) RawVersionString() string {
	return b.RawVersion
}

func NewBuildInfo(
	pkgName, version, rawVersion,
	buildId, commit, date,
	author, copyrightStartYear string,
) BuildInfo {
	info := BuildInfo{
		PkgName: pkgName,

		Version:    version,
		RawVersion: rawVersion,
		BuildId:    buildId,
		GitCommit:  commit,
		Date:       date,
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),

		AuthorName:         author,
		CopyrightStartYear: copyrightStartYear,
		CopyrightNowYear:   strconv.Itoa(time.Now().Year()),
	}

	bi, available := debug.ReadBuildInfo()
	if !available {
		return info
	}

	info.GoVersion = bi.GoVersion
	if info.GoVersion == "" {
		info.GoVersion = infoUnknown
	}

	if info.Version == "" || info.Version == infoUnknown {
		info.Version = firstNonEmpty(getGitVersion(bi), versionDevel)
	}

	nowVersion = info.Version
	nowBuildId = info.BuildId

	if date != "" {
		return info
	}

	var revision string
	var modified string
	for _, setting := range bi.Settings {
		// The `vcs.xxx` information is only available with `go build`.
		// This information is not available with `go install` or `go run`.
		switch setting.Key {
		case "vcs.time":
			info.Date = setting.Value
		case "vcs.revision":
			revision = setting.Value
		case "vcs.modified":
			modified = setting.Value
		}
	}

	if revision == "" {
		revision = infoUnknown
	}

	if modified == "" {
		modified = "?"
	}

	if info.Date == "" {
		info.Date = fmt.Sprintf("(%s)", infoUnknown)
	}

	if info.BuildId == "" {
		info.BuildId = fmt.Sprintf("(%s)", infoUnknown)
	}

	info.GitCommit = fmt.Sprintf("(%s, modified: %s, mod sum: %q)", revision, modified, bi.Main.Sum)

	return info
}

func getGitVersion(bi *debug.BuildInfo) string {
	if bi == nil {
		return ""
	}

	// remove this when the issue https://github.com/golang/go/issues/29228 is fixed
	if bi.Main.Version == "(devel)" || bi.Main.Version == "" {
		return ""
	}

	return bi.Main.Version
}

//nolint:golint,unused
func getCommit(bi *debug.BuildInfo) string {
	return getKey(bi, "vcs.revision")
}

//nolint:golint,unused
func getDirty(bi *debug.BuildInfo) string {
	modified := getKey(bi, "vcs.modified")
	if modified == "true" {
		return "dirty"
	}
	if modified == "false" {
		return "clean"
	}
	return ""
}

//nolint:golint,unused
func getBuildDate(bi *debug.BuildInfo) string {
	buildTime := getKey(bi, "vcs.time")
	t, err := time.Parse("2006-01-02T15:04:05Z", buildTime)
	if err != nil {
		return ""
	}
	return t.Format("2006-01-02T15:04:05")
}

//nolint:golint,unused
func getKey(bi *debug.BuildInfo, key string) string {
	if bi == nil {
		return ""
	}
	for _, iter := range bi.Settings {
		if iter.Key == key {
			return iter.Value
		}
	}
	return ""
}

//nolint:golint,unused
func firstNonEmpty(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}
