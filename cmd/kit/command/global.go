package command

import (
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/bar-counter/slog"
	"github.com/convention-change/convention-change-log/changelog"
	"github.com/convention-change/convention-change-log/cmd/kit/command/exit_cli"
	"github.com/convention-change/convention-change-log/cmd/kit/constant"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/convention-change/convention-change-log/internal/log"
	"github.com/convention-change/convention-change-log/internal/pkgJson"
	goGit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/gookit/color"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/sinlov-go/go-git-tools/git"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/sinlov-go/sample-markdown/sample_mk"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type GlobalConfig struct {
	LogLevel      string
	TimeoutSecond uint
}

type (
	// GlobalCommand
	//	command root
	GlobalCommand struct {
		Name    string
		Version string
		Verbose bool
		DryRun  bool

		GitRootPath   string
		GitRemote     string
		ChangeLogSpec *convention.ConventionalChangeLogSpec

		RootCfg GlobalConfig

		GenerateConfig GenerateConfig
	}
)

var (
	cmdGlobalEntry *GlobalCommand
)

// CmdGlobalEntry
//
//	return global command entry
func CmdGlobalEntry() *GlobalCommand {
	return cmdGlobalEntry
}

// globalExec
//
//	do global command exec
func (c *GlobalCommand) globalExec() error {

	slog.Debug("-> start GlobalAction")

	_, err := git_info.IsPathGitManagementRoot(c.GitRootPath)
	if err != nil {
		return exit_cli.Format("cli run path not git repository root, please check path at: %s", c.GitRootPath)
	}

	var repository git.Repository
	if c.GenerateConfig.GitCloneUrl == "" {

		repositoryOpen, errOpen := git.NewRepositoryRemoteByPath(c.GitRemote, c.GitRootPath)
		if errOpen != nil {
			return exit_cli.Format("load local git repository error: %s", errOpen)
		}
		repository = repositoryOpen

	} else {
		repositoryClone, errClone := git.NewRepositoryRemoteClone(c.GitRemote, memory.NewStorage(), nil, &goGit.CloneOptions{
			URL:        c.GenerateConfig.GitCloneUrl,
			RemoteName: c.GitRemote,
		})
		if errClone != nil {
			return exit_cli.Format("clone git repository %s \nerror: %s", c.GenerateConfig.GitCloneUrl, errClone)
		}
		repository = repositoryClone
	}
	if repository == nil {
		return exit_cli.Format("can not load git repository")
	}
	headBranchName, err := repository.HeadBranchName()
	if err != nil {
		slog.Error("HeadBranchName err: %v", err)
		return err
	}

	gitRemoteInfo, err := repository.RemoteInfo(c.GitRemote, 0)
	if err != nil {
		slog.Error("RemoteInfo err: %v", err)
		return exit_cli.Err(err)
	}

	if c.Verbose {
		bytes, errJson := json.Marshal(gitRemoteInfo)
		if errJson != nil {
			slog.Error("Marshal err: %v", errJson)
			return exit_cli.Err(errJson)
		}
		slog.Debugf("gitRemoteInfo:\n%s", string(bytes))
	}

	changeLogNodes := make([]sample_mk.Node, 0)
	reader, errHistory := changelog.NewReader(c.GenerateConfig.Infile, *c.ChangeLogSpec)
	if errHistory != nil {
		slog.Debugf("can not read history changelog err: %v", errHistory)
		c.GenerateConfig.ReleaseAs = convention.DefaultSemverVersion
		c.GenerateConfig.ReleaseTag = fmt.Sprintf("%s%s", c.GenerateConfig.TagPrefix, c.GenerateConfig.ReleaseAs)
	} else {
		changeLogNodes = append(changeLogNodes, reader.HistoryNodes()...)
	}

	historyFirstTagName := ""
	if c.GenerateConfig.FromCommit == "" {
		if errHistory != nil {
			// this not find any tag and history
			c.GenerateConfig.FromCommit = ""
		} else {
			historyFirstTag := reader.HistoryFirstTag()
			tagSearchByName, errTagSearchByName := repository.CommitTagSearchByName(historyFirstTag)
			if errTagSearchByName != nil {
				slog.Debugf("errTagSearchByName err: %v", errTagSearchByName)
				c.GenerateConfig.FromCommit = ""
			} else {
				c.GenerateConfig.FromCommit = tagSearchByName.Hash.String()
				historyFirstTagName = historyFirstTag
			}
		}
	}
	slog.Debugf("historyFirstTagName: %s", historyFirstTagName)
	slog.Debugf("c.GenerateConfig.FromCommit: %s", c.GenerateConfig.FromCommit)

	latestCommits, errLatestCommits := repository.Log("", c.GenerateConfig.FromCommit)
	if errLatestCommits != nil {
		slog.Error("errLatestCommits err: %v", errLatestCommits)
		return exit_cli.Err(errLatestCommits)
	}

	slog.Debugf("latestCommits len %d", len(latestCommits))

	gitHttpHost := gitRemoteInfo.Host
	if c.ChangeLogSpec.CoverHttpHost != "" {
		gitHttpHost = c.ChangeLogSpec.CoverHttpHost
	}
	gitHttpInfoDefault := convention.GitRepositoryHttpInfo{
		Scheme:     "https",
		Host:       gitHttpHost,
		Owner:      gitRemoteInfo.User,
		Repository: gitRemoteInfo.Repo,
	}

	conventionCommits, errConvert := convention.ConvertGitCommits2ConventionCommits(latestCommits, *c.ChangeLogSpec, gitHttpInfoDefault)
	if errConvert != nil {
		slog.Error("ConvertGitCommits2ConventionCommits err: %v", errConvert)
		return exit_cli.Err(errConvert)
	}
	generateMarkdownNodes, featNodes, errConvert := changelog.GenerateMarkdownNodes(
		gitHttpInfoDefault,
		conventionCommits,
		*c.ChangeLogSpec,
	)
	if errConvert != nil {
		slog.Error("GenerateMarkdownNodes err: %v", errConvert)
		return exit_cli.Err(errConvert)
	}

	if c.GenerateConfig.ReleaseAs != "" {
		// check cli settings ReleaseAs tag is exists
		oldReleaseTag, errOldReleaseTag := repository.CommitTagSearchByName(c.GenerateConfig.ReleaseTag)
		if err != nil {
			slog.Debugf("not find tag: %s err: %v", c.GenerateConfig.ReleaseTag, errOldReleaseTag)
		}
		if oldReleaseTag != nil {
			errReleaseTagExist := fmt.Errorf("want release tag is exist: %s, tag message is %s", c.GenerateConfig.ReleaseTag, oldReleaseTag.Message)
			slog.Error("ReleaseTagExist err: %v", errReleaseTagExist)
			return exit_cli.Err(errReleaseTagExist)
		}
		// check version as semver
		_, errSemverCheck := semver.NewVersion(c.GenerateConfig.ReleaseAs)
		if errSemverCheck != nil {
			slog.Error("NewVersion err: %v", errSemverCheck)
			return exit_cli.ErrMsg(errSemverCheck, "release-as is not semver")
		}
	} else {
		// find new version as semver by historyFirstTagName
		historyVersion, errHistorySemver := semver.NewVersion(historyFirstTagName)
		if errHistorySemver != nil {
			slog.Error("find new version as semver by historyFirstTagName err: %v", errHistorySemver)
			return exit_cli.Err(errHistorySemver)
		}
		var version semver.Version
		if len(featNodes) > 0 {
			version = historyVersion.IncMinor()
		} else {
			version = historyVersion.IncPatch()
		}
		c.GenerateConfig.ReleaseAs = version.String()
		c.GenerateConfig.ReleaseTag = fmt.Sprintf("%s%s", c.GenerateConfig.TagPrefix, c.GenerateConfig.ReleaseAs)
	}

	changelogDesc := changelog.ConventionalChangeLogDesc{
		Version:      c.GenerateConfig.ReleaseTag,
		When:         time.Now(),
		Location:     time.Local,
		ToolsKitName: constant.KitName,
		ToolsKitURL:  constant.KitUrl,
	}
	if reader.HistoryFirstTagShort() != "" {
		changelogDesc.PreviousTag = reader.HistoryFirstTagShort()
	}

	nodesGenerateWithTitle, errAddTitle := changelog.AddMarkdownChangelogNodesTitle(
		generateMarkdownNodes,
		gitHttpInfoDefault,
		changelogDesc,
		*c.ChangeLogSpec,
	)

	if errAddTitle != nil {
		slog.Error("AddMarkdownChangelogNodesTitle err: %v", errAddTitle)
		return exit_cli.Err(errAddTitle)
	}

	err = pkgJson.ReplaceJsonVersionByLine(filepath.Join(c.GitRootPath, "package.json"), c.GenerateConfig.ReleaseAs)
	if err != nil {
		slog.Error("ReplaceJsonVersionByLine", err)
	}

	if c.DryRun {
		latestMarkdownContent := sample_mk.GenerateText(nodesGenerateWithTitle)
		color.Printf(constant.CmdHelpOutputting, c.GenerateConfig.Outfile)
		color.Println("")

		color.Println(constant.LogLineSpe)
		color.Grayf("%s\n", latestMarkdownContent)
		color.Println(constant.LogLineSpe)

		color.Printf(constant.CmdHelpCommitting, c.GenerateConfig.Infile)
		color.Println("")
		color.Printf(constant.CmdHelpTagRelease, c.GenerateConfig.ReleaseTag)
		color.Println("")
		return nil
	}

	// add history
	changeLogNodes = append(nodesGenerateWithTitle, changeLogNodes...)

	fullMarkdownContent := sample_mk.GenerateText(changeLogNodes)

	errChangeLocalFile := c.changeLocalFiles(fullMarkdownContent)
	if errChangeLocalFile != nil {
		slog.Error("WriteFileByByte err: %v", errChangeLocalFile)
		return exit_cli.Err(errChangeLocalFile)
	}

	errDoGit := c.doGit(headBranchName)
	if errDoGit != nil {
		slog.Error("doGit err: %v", errDoGit)
		return exit_cli.Err(errDoGit)
	}

	return nil
}

func (c *GlobalCommand) changeLocalFiles(fullChangeLogContent string) error {
	errWriteFile := filepath_plus.WriteFileByByte(c.GenerateConfig.Outfile, []byte(fullChangeLogContent), os.FileMode(0766), true)
	if errWriteFile != nil {
		return fmt.Errorf("WriteFileByByte err: %v", errWriteFile)
	}

	if c.GenerateConfig.ReleaseAs != "" {
		// try update node
		pkgJsonPath := filepath.Join(c.GitRootPath, "package.json")
		if filepath_plus.PathExistsFast(pkgJsonPath) {
			// replace file line by regexp
			slog.Debugf("try update node version in file: %s", pkgJsonPath)
			err := pkgJson.ReplaceJsonVersionByLine(pkgJsonPath, c.GenerateConfig.ReleaseAs)
			if err != nil {
				slog.Error("ReplaceJsonVersionByLine", err)
			}
			pkgJsonLockPath := filepath.Join(c.GitRootPath, "package-lock.json")
			if filepath_plus.PathExistsFast(pkgJsonLockPath) {
				output, errNpmInstall := exec.Command("npm", "install").CombinedOutput()
				if errNpmInstall != nil {
					slog.Error("do Command npm install error", errNpmInstall)
				}
				slog.Debugf("npm install output:\n%s", output)
			}
		}
	}

	return nil
}

func (c *GlobalCommand) doGit(branchName string) error {
	// disable git-go repository issues until https://github.com/go-git/go-git/issues/180 is fixed

	cmdOutput, err := exec.Command("git", "add", "--all").CombinedOutput()
	if err != nil {
		return err
	}
	slog.Debugf("git add output:\n%s", cmdOutput)

	releaseCommit := new(convention.ReleaseCommitMessageRenderTemplate)
	releaseCommit.CurrentTag = c.GenerateConfig.ReleaseAs
	releaseCommitMsg, errRender := convention.RaymondRender(c.ChangeLogSpec.ReleaseCommitMessageFormat, releaseCommit)
	if errRender != nil {
		return errRender
	}

	cmdOutput, err = exec.Command("git", "commit", "-m", releaseCommitMsg).CombinedOutput()
	if err != nil {
		return err
	}
	slog.Debugf("git commit output:\n%s", cmdOutput)

	cmdOutput, err = exec.Command("git", "tag", c.GenerateConfig.ReleaseTag, "-m", releaseCommitMsg).CombinedOutput()
	if err != nil {
		return err
	}
	slog.Debugf("git tag output:\n%s", cmdOutput)

	color.Printf(constant.CmdHelpOutputting, c.GenerateConfig.Outfile)
	color.Println("")
	color.Printf(constant.CmdHelpCommitting, c.GenerateConfig.Infile)
	color.Println("")
	color.Printf(constant.CmdHelpTagRelease, c.GenerateConfig.ReleaseTag)
	color.Println("")

	if c.GenerateConfig.AutoPush {
		cmdOutput, err = exec.Command("git", "push", "--follow-tags", "origin", branchName).CombinedOutput()
		if err != nil {
			return err
		}
		slog.Debugf("git push output:\n%s", cmdOutput)
		return nil
	}

	color.Printf(constant.CmdHelpGitPush, branchName)
	color.Println("")
	return nil
}

type GenerateConfig struct {
	GitCloneUrl string

	ReleaseAs  string
	TagPrefix  string
	ReleaseTag string

	Infile  string
	Outfile string

	FromCommit string

	AutoPush bool
}

// withGlobalFlag
//
// bind global flag to globalExec
func withGlobalFlag(c *cli.Context, cliVersion, cliName string) (*GlobalCommand, error) {
	slog.Debug("-> withGlobalFlag")

	isVerbose := c.Bool("verbose")

	dir, err := os.Getwd()
	if err != nil {
		return nil, exit_cli.Err(err)
	}
	gitRootFolder := dir

	config := GlobalConfig{
		LogLevel:      c.String("config.log_level"),
		TimeoutSecond: c.Uint("config.timeout_second"),
	}

	// c.String("clone-url") close this way to get latest tag

	tagPrefix := c.String("tag-prefix")
	cliReleaseAs := c.String("release-as")
	cliReleaseTag := ""
	if cliReleaseAs != "" {
		cliReleaseTag = fmt.Sprintf("%s%s", tagPrefix, cliReleaseAs)
	}

	generateConfig := GenerateConfig{
		GitCloneUrl: "",
		ReleaseAs:   cliReleaseAs,
		TagPrefix:   tagPrefix,
		ReleaseTag:  cliReleaseTag,

		Infile:  c.String("infile"),
		Outfile: c.String("outfile"),

		FromCommit: c.String("from-commit"),

		AutoPush: c.Bool("auto-push"),
	}

	var changeLogSpec *convention.ConventionalChangeLogSpec
	specFilePath := filepath.Join(gitRootFolder, constant.VersionRcFileName)
	if filepath_plus.PathExistsFast(specFilePath) {
		specByte, errReadSpec := filepath_plus.ReadFileAsByte(specFilePath)
		if errReadSpec != nil {
			return nil, exit_cli.Err(errReadSpec)
		}
		spec, errReadSpec := convention.LoadConventionalChangeLogSpecByData(specByte)
		if errReadSpec != nil {
			return nil, exit_cli.Err(errReadSpec)
		}
		changeLogSpec = spec
	} else {
		spec := convention.DefaultConventionalChangeLogSpec()
		changeLogSpec = &spec
	}

	p := GlobalCommand{
		Name:    cliName,
		Version: cliVersion,
		Verbose: isVerbose,
		DryRun:  c.Bool("dry-run"),

		GitRootPath:   gitRootFolder,
		GitRemote:     c.String("git-remote"),
		ChangeLogSpec: changeLogSpec,

		RootCfg:        config,
		GenerateConfig: generateConfig,
	}
	return &p, nil
}

// GlobalBeforeAction
// do command Action before flag global.
func GlobalBeforeAction(c *cli.Context) error {
	isVerbose := c.Bool("verbose")
	err := log.InitLog(isVerbose, !isVerbose)
	if err != nil {
		panic(err)
	}
	cliVersion := pkgJson.GetPackageJsonVersionGoStyle(false)
	if isVerbose {
		slog.Warnf("-> open verbose, and now command version is: %s", cliVersion)
	}
	appName := pkgJson.GetPackageJsonName()
	cmdGlobalEntry, err = withGlobalFlag(c, cliVersion, appName)
	if err != nil {
		return err
	}

	return nil
}

// GlobalAction
// do command Action flag.
func GlobalAction(c *cli.Context) error {
	if cmdGlobalEntry == nil {
		panic(fmt.Errorf("not init GlobalBeforeAction success to new cmdGlobalEntry"))
	}

	err := cmdGlobalEntry.globalExec()
	if err != nil {
		return err
	}
	return nil
}

// GlobalAfterAction
//
//	do command Action after flag global.
//
//nolint:golint,unused
func GlobalAfterAction(c *cli.Context) error {
	isVerbose := c.Bool("verbose")
	if isVerbose {
		slog.Infof("-> finish run command: %s, version %s", cmdGlobalEntry.Name, cmdGlobalEntry.Version)
	}
	return nil
}
