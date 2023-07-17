package subcommand_init

import (
	"github.com/convention-change/convention-change-log/cmd/kit/command"
	"github.com/convention-change/convention-change-log/cmd/kit/command/exit_cli"
	"github.com/convention-change/convention-change-log/convention"
	"path/filepath"

	"github.com/bar-counter/slog"
	"github.com/convention-change/convention-change-log/internal/urfave_cli"
	"github.com/gookit/color"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/urfave/cli/v2"
)

const (
	commandName = "init"

	versionRcFile = ".versionrc"
)

var commandEntry *NewCommand

type NewCommand struct {
	isDebug bool

	TargetFile string
	MoreConfig bool
}

func (n *NewCommand) Exec() error {
	slog.Debugf("-> Exec subCommand [ %s ]", commandName)

	if filepath_plus.PathExistsFast(n.TargetFile) {
		color.Yellowf("init versionrc file is exists, file: %s\n", n.TargetFile)
		return nil
	}

	var spec *convention.ConventionalChangeLogSpec
	if n.MoreConfig {
		logSpec := convention.DefaultConventionalChangeLogSpec()
		spec = &logSpec
	} else {
		spec = convention.SimplifyConventionalChangeLogSpec()
	}

	err := filepath_plus.WriteFileAsJsonBeauty(n.TargetFile, spec, false)
	if err != nil {
		return exit_cli.Format("write file %s err: %v", n.TargetFile, err)
	}
	color.Greenf("init .versionrc file success, file: %s", n.TargetFile)
	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "more",
			Usage: "more config at init",
		},
	}
}

func withEntry(c *cli.Context) (*NewCommand, error) {
	slog.Debugf("-> withEntry subCommand [ %s ]", commandName)

	globalEntry := command.CmdGlobalEntry()

	targetFile := filepath.Join(globalEntry.GitRootPath, versionRcFile)

	return &NewCommand{
		isDebug: globalEntry.Verbose,

		TargetFile: targetFile,

		MoreConfig: c.Bool("more"),
	}, nil
}

func action(c *cli.Context) error {
	slog.Debugf("-> Sub Command action [ %s ] start", commandName)
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	commandEntry = entry
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "init convention change log config, this cli must run in git root folder",
			Action: action,
			Flags:  flag(),
		},
	}
}
