package command

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/convention-change/convention-change-log/internal/log"
	"github.com/convention-change/convention-change-log/internal/pkgJson"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/urfave/cli/v2"
	"os"
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

		GitRootPath string

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

	return nil
}

type GenerateConfig struct {
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
		return nil, fmt.Errorf("can not get target foler err: %v", err)
	}
	gitRootFolder := dir
	_, err = git_info.IsPathGitManagementRoot(gitRootFolder)
	if err != nil {
		return nil, err
	}

	config := GlobalConfig{
		LogLevel:      c.String("config.log_level"),
		TimeoutSecond: c.Uint("config.timeout_second"),
	}

	generateConfig := GenerateConfig{
		ReleaseAs: c.String("release-as"),
		TagPrefix: c.String("tag-prefix"),

		Infile:  c.String("infile"),
		Outfile: c.String("outfile"),

		FromCommit: c.String("from-commit"),

		AutoPush: c.Bool("auto-push"),
	}

	p := GlobalCommand{
		Name:    cliName,
		Version: cliVersion,
		Verbose: isVerbose,
		DryRun:  c.Bool("dry-run"),

		GitRootPath: gitRootFolder,

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
