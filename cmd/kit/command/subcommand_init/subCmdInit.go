package subcommand_init

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bar-counter/slog"
	convention_change_log "github.com/convention-change/convention-change-log"
	"github.com/convention-change/convention-change-log/cmd/kit/command"
	"github.com/convention-change/convention-change-log/cmd/kit/command/exit_cli"
	"github.com/convention-change/convention-change-log/cmd/kit/constant"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/convention-change/convention-change-log/internal/tools"
	"github.com/convention-change/convention-change-log/internal/urfave_cli"
	"github.com/gookit/color"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/urfave/cli/v2"
)

const (
	commandName = "init"
)

var commandEntry *InitCommand

type InitCommand struct {
	isDebug   bool
	isDruInit bool

	TargetFile string
	MoreConfig bool

	SkipBranchCheck     bool
	DryRun              bool
	BranchCheckPatterns []string
	GitRootPath         string
}

func (n *InitCommand) Exec() error {
	slog.Debugf("-> Exec subCommand [ %s ]", commandName)

	var branchCheckWarning string
	// branch check
	if !n.SkipBranchCheck {
		branchName, matched, err := command.CheckBranchAtRoot(n.GitRootPath, n.BranchCheckPatterns)
		if err != nil {
			slog.Warnf("branch check error: %v", err)
		} else if !matched {
			if n.DryRun {
				branchCheckWarning = fmt.Sprintf(
					"current branch %q does not match branch-check patterns %v, this may not be the intended branch",
					branchName,
					n.BranchCheckPatterns,
				)
			} else {
				return exit_cli.Format(
					"current branch %q does not match branch-check patterns %v, please switch to a matching branch or use --skip-branch",
					branchName,
					n.BranchCheckPatterns,
				)
			}
		}
	}

	defer func() {
		if branchCheckWarning != "" {
			slog.Warnf("%s", branchCheckWarning)
		}
	}()

	if filepath_plus.PathExistsFast(n.TargetFile) {
		color.Yellowf("init versionrc file is exists, file: %s\n", n.TargetFile)
		return nil
	}

	var spec *convention.ConventionalChangeLogSpec
	if !n.MoreConfig {

		if n.isDruInit {
			color.Greenln("will init at: %s", n.TargetFile)
			color.Grayf("%s", convention_change_log.ResVersionRcBeautyJson)
			return nil
		}

		err := filepath_plus.WriteFileByByte(
			n.TargetFile,
			[]byte(convention_change_log.ResVersionRcBeautyJson),
			os.FileMode(0o666),
			false,
		)
		if err != nil {
			slog.Error("init .versionrc file err: %v", err)
			return exit_cli.Err(err)
		}
		return nil
	}
	logSpec := convention.DefaultConventionalChangeLogSpec()
	spec = &logSpec

	if n.isDruInit {
		color.Greenln("will init at: %s", n.TargetFile)
		res, errJsonMarshalBeauty := tools.JsonMarshalBeauty(spec)
		if errJsonMarshalBeauty != nil {
			return errJsonMarshalBeauty
		}
		color.Grayf("%s", res)

		return nil
	}

	err := filepath_plus.WriteFileAsJsonBeauty(n.TargetFile, spec, false)
	if err != nil {
		slog.Error("write .versionrc file err: %v", err)

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
		&cli.BoolFlag{
			Name:  "dry-init",
			Usage: "only show init config file content (1.9.+)",
		},
	}
}

func withEntry(c *cli.Context) (*InitCommand, error) {
	slog.Debugf("-> withEntry subCommand [ %s ]", commandName)

	globalEntry := command.CmdGlobalEntry()

	targetFile := filepath.Join(globalEntry.GitRootPath, constant.VersionRcFileName)

	return &InitCommand{
		isDebug:   globalEntry.Verbose,
		isDruInit: c.Bool("dry-init"),

		TargetFile: targetFile,
		MoreConfig: c.Bool("more"),

		SkipBranchCheck: globalEntry.GenerateConfig.SkipBranchCheck,
		DryRun:          globalEntry.DryRun,
		BranchCheckPatterns: command.ResolveBranchCheckPatterns(
			globalEntry.ChangeLogSpec.BranchCheck,
		),
		GitRootPath: globalEntry.GitRootPath,
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
			Usage:  "init convention change log config, this cli must run in git root folder\ncan use --dry-run",
			Action: action,
			Flags:  flag(),
		},
	}
}
