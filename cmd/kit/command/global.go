package command

import (
	"encoding/json"
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/convention-change/convention-change-log/cmd/kit/command/exit_cli"
	"github.com/convention-change/convention-change-log/cmd/kit/constant"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/convention-change/convention-change-log/internal/log"
	"github.com/convention-change/convention-change-log/internal/pkgJson"
	goGit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/sinlov-go/go-git-tools/git"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
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

	if c.GenerateConfig.ReleaseAs == "" {
		return exit_cli.Format("args --release-as is empty")
	}

	_, err := git_info.IsPathGitManagementRoot(c.GitRootPath)
	if err != nil {
		return exit_cli.Format("cli run path not git repository root, please check path at: %s", c.GitRootPath)
	}
	fistRemoteInfo, err := git_info.RepositoryFistRemoteInfo(c.GitRootPath, c.GitRemote)
	if err != nil {
		return exit_cli.Err(err)
	}
	if c.Verbose {
		bytes, errJson := json.Marshal(fistRemoteInfo)
		if errJson != nil {
			return exit_cli.Err(errJson)
		}
		slog.Debugf("fistRemoteInfo:\n%s", string(bytes))
	}

	var repository *git.Repository
	if c.GenerateConfig.GitCloneUrl == "" {

		repositoryOpen, errOpen := git.NewRepositoryRemoteByPath(c.GitRemote, c.GitRootPath)
		if errOpen != nil {
			return exit_cli.Format("load local git repository error: %s", errOpen)
		}
		repository = &repositoryOpen

	} else {
		repositoryClone, errClone := git.NewRepositoryRemoteClone(c.GitRemote, memory.NewStorage(), nil, &goGit.CloneOptions{
			URL:        c.GenerateConfig.GitCloneUrl,
			RemoteName: c.GitRemote,
		})
		if errClone != nil {
			return exit_cli.Format("clone git repository %s \nerror: %s", c.GenerateConfig.GitCloneUrl, errClone)
		}
		repository = &repositoryClone
	}
	if repository == nil {
		return exit_cli.Format("can not load git repository")
	}

	return nil
}

type GenerateConfig struct {
	GitCloneUrl string

	ReleaseAs string
	TagPrefix string

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

	generateConfig := GenerateConfig{
		GitCloneUrl: c.String("clone-url"),
		ReleaseAs:   c.String("release-as"),
		TagPrefix:   c.String("tag-prefix"),

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
