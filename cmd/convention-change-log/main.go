//go:build !test

package main

import (
	"fmt"
	"github.com/convention-change/convention-change-log"
	"github.com/convention-change/convention-change-log/cmd/kit/cli"
	"github.com/convention-change/convention-change-log/internal/pkgJson"
	"github.com/gookit/color"
	"os"
)

func main() {
	pkgJson.InitPkgJsonContent(convention_change_log.PackageJson)
	app := cli.NewCliApp()
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
		os.Exit(2)
	}
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
