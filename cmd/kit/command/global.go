package command

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/convention-change/convention-change-log/internal/log"
	"github.com/convention-change/convention-change-log/internal/pkgJson"
	"github.com/urfave/cli/v2"
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
		RootCfg GlobalConfig
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

// GlobalAction
// do command Action flag.
func GlobalAction(c *cli.Context) error {
	if cmdGlobalEntry == nil {
		panic(fmt.Errorf("not init GlobalBeforeAction success to new cmdGlobalEntry"))
	}

	slog.Debug("-> start GlobalAction")

	err := cmdGlobalEntry.globalExec()
	if err != nil {
		return err
	}
	return nil
}

func (c *GlobalCommand) globalExec() error {

	return nil
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
	cmdGlobalEntry = withGlobalFlag(c, cliVersion, appName)

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

func withGlobalFlag(c *cli.Context, cliVersion, cliName string) *GlobalCommand {
	isVerbose := c.Bool("verbose")

	config := GlobalConfig{
		LogLevel:      c.String("config.log_level"),
		TimeoutSecond: c.Uint("config.timeout_second"),
	}

	p := GlobalCommand{
		Name:    cliName,
		Version: cliVersion,
		Verbose: isVerbose,
		DryRun:  c.Bool("dry-run"),
		RootCfg: config,
	}
	return &p
}
