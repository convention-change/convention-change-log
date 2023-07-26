package cli

import (
	"fmt"
	"github.com/convention-change/convention-change-log/cmd/kit/command"
	"github.com/convention-change/convention-change-log/cmd/kit/command/subcommand_init"
	"github.com/convention-change/convention-change-log/cmd/kit/command/subcommand_read_latest"
	"github.com/convention-change/convention-change-log/cmd/kit/constant"
	"github.com/convention-change/convention-change-log/internal/pkgJson"
	"github.com/convention-change/convention-change-log/internal/urfave_cli"
	"github.com/urfave/cli/v2"
	"runtime"
	"time"
)

func NewCliApp() *cli.App {

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = pkgJson.GetPackageJsonVersionGoStyle(false)
	app.Name = pkgJson.GetPackageJsonName()
	if pkgJson.GetPackageJsonHomepage() != "" {
		app.Usage = fmt.Sprintf("see: %s", pkgJson.GetPackageJsonHomepage())
	}
	app.Description = pkgJson.GetPackageJsonDescription()

	year := time.Now().Year()
	jsonAuthor := pkgJson.GetPackageJsonAuthor()
	app.Copyright = fmt.Sprintf("© %s-%d %s by: %s, run on %s %s",
		constant.CopyrightStartYear, year, jsonAuthor.Name, runtime.Version(), runtime.GOOS, runtime.GOARCH)
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
