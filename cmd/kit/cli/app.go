package cli

import (
	"fmt"
	"github.com/convention-change/convention-change-log"
	command2 "github.com/convention-change/convention-change-log/cmd/kit/command"
	"github.com/convention-change/convention-change-log/cmd/kit/command/subcommand_init"
	"github.com/convention-change/convention-change-log/cmd/kit/constant"
	"github.com/convention-change/convention-change-log/internal/pkgJson"
	"github.com/convention-change/convention-change-log/internal/urfave_cli"
	"github.com/urfave/cli/v2"
	"time"
)

func NewCliApp() *cli.App {
	pkgJson.InitPkgJsonContent(convention_change_log.PackageJson)
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
	app.Copyright = fmt.Sprintf("© %s-%d %s", constant.CopyrightStartYear, year, jsonAuthor.Name)
	author := &cli.Author{
		Name:  jsonAuthor.Name,
		Email: jsonAuthor.Email,
	}
	app.Authors = []*cli.Author{
		author,
	}

	flags := urfave_cli.UrfaveCliAppendCliFlag(command2.GlobalFlag(), command2.HideGlobalFlag())
	flags = urfave_cli.UrfaveCliAppendCliFlag(flags, command2.MainFlag())

	app.Flags = flags
	app.Before = command2.GlobalBeforeAction
	app.Action = command2.GlobalAction
	app.After = command2.GlobalAfterAction

	var appCommands []*cli.Command
	appCommands = urfave_cli.UrfaveCliAppendCliCommand(appCommands, subcommand_init.Command())
	app.Commands = appCommands

	return app
}
