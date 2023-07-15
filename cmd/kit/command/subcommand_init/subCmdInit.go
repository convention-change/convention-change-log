package subcommand_init

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/gookit/color"
	command2 "github.com/sinlov-go/convention-change-log/cmd/kit/command"
	"github.com/sinlov-go/convention-change-log/convention"
	"github.com/sinlov-go/convention-change-log/internal/urfave_cli"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

const (
	commandName = "init"

	versionRcFile = ".versionrc"
)

var commandEntry *NewCommand

type NewCommand struct {
	isDebug bool

	GitRootPath string

	TargetFile string
}

func (n *NewCommand) Exec() error {

	if filepath_plus.PathExistsFast(n.TargetFile) {
		color.Yellowf("init versionrc file is exists, file: %s", n.TargetFile)
	} else {
		err := filepath_plus.WriteFileAsJsonBeauty(n.TargetFile, convention.DefaultConventionalChangeLogSpec(), false)
		if err != nil {
			return fmt.Errorf("write file %s err: %v", n.TargetFile, err)
		}
		color.Greenf("init .versionrc file success, file: %s", n.TargetFile)
	}
	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{}
}

func withEntry(c *cli.Context) (*NewCommand, error) {
	globalEntry := command2.CmdGlobalEntry()

	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("can not get target foler err: %v", err)
	}
	gitRootFolder := dir
	_, err = git_info.IsPathGitManagementRoot(gitRootFolder)
	if err != nil {
		return nil, err
	}

	targetFile := filepath.Join(gitRootFolder, versionRcFile)

	return &NewCommand{
		isDebug:     globalEntry.Verbose,
		GitRootPath: gitRootFolder,
		TargetFile:  targetFile,
	}, nil
}

func action(c *cli.Context) error {
	slog.Debugf("SubCommand [ %s ] start", commandName)
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	commandEntry = entry
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command2.GlobalFlag(), command2.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "init convention change log config, this cli must run in git root folder",
			Action: action,
			Flags:  flag(),
		},
	}
}
