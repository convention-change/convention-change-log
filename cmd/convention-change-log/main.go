//go:build !test

package main

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/sinlov-go/convention-change-log/cmd/kit/cli"
	"os"
)

func main() {
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
