//go:build !test

package main

import (
	"fmt"
	"github.com/convention-change/convention-change-log"
	"github.com/convention-change/convention-change-log/cmd/kit/cli"
	"github.com/convention-change/convention-change-log/constant"
	"github.com/convention-change/convention-change-log/internal/pkg_kit"
	"github.com/gookit/color"
	"os"
)

const (
	// exitCodeCmdArgs SIGINT as 2
	exitCodeCmdArgs = 2
)

//nolint:gochecknoglobals
var (
	// Populated by goreleaser during build
	version    = "unknown"
	rawVersion = "unknown"
	buildID    string
	commit     = "?"
	date       = ""
)

func init() {
	if buildID == "" {
		buildID = "unknown"
	}
}

func main() {
	pkg_kit.InitPkgJsonContent(convention_change_log.PackageJson)

	bdInfo := pkg_kit.NewBuildInfo(
		pkg_kit.GetPackageJsonName(),
		version,
		rawVersion,
		buildID,
		commit,
		date,
		pkg_kit.GetPackageJsonAuthor().Name,
		constant.CopyrightStartYear,
	)

	app := cli.NewCliApp(bdInfo)
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
		os.Exit(exitCodeCmdArgs)
	}
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
		os.Exit(exitCodeCmdArgs)
	}
}
