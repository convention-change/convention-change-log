package cli

import (
	"fmt"
	"github.com/convention-change/convention-change-log/cmd/kit/command"
	"github.com/convention-change/convention-change-log/cmd/kit/command/subcommand_init"
	"github.com/convention-change/convention-change-log/cmd/kit/command/subcommand_read_latest"
	"github.com/convention-change/convention-change-log/internal/pkg_kit"
	"github.com/convention-change/convention-change-log/internal/urfave_cli"
	"github.com/urfave/cli/v2"
)

func NewCliApp(bdInfo pkg_kit.BuildInfo) *cli.App {

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = bdInfo.PgkNameString()
	app.Version = bdInfo.VersionString()
	if pkg_kit.GetPackageJsonHomepage() != "" {
		app.Usage = fmt.Sprintf("see: %s", pkg_kit.GetPackageJsonHomepage())
	}
	app.Description = pkg_kit.GetPackageJsonDescription()
	jsonAuthor := pkg_kit.GetPackageJsonAuthor()
	app.Copyright = bdInfo.String()
	author := &cli.Author{
		Name:  jsonAuthor.Name,
		Email: jsonAuthor.Email,
	}
	app.Authors = []*cli.Author{
		author,
	}
	app.UsageText = fmt.Sprintf("%s --dry-run -r <release-as>", app.Name)

	flags := urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	flags = urfave_cli.UrfaveCliAppendCliFlag(flags, command.MainFlag())

	app.Flags = flags
	app.Before = command.GlobalBeforeAction
	app.Action = command.GlobalAction
	app.After = command.GlobalAfterAction

	var appCommands []*cli.Command
	appCommands = urfave_cli.UrfaveCliAppendCliCommand(appCommands, subcommand_init.Command())
	appCommands = urfave_cli.UrfaveCliAppendCliCommand(appCommands, subcommand_read_latest.Command())
	app.Commands = appCommands

	return app
}
