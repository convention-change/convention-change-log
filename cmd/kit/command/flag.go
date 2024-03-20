package command

import (
	"fmt"
	"github.com/convention-change/convention-change-log/cmd/kit/constant"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/urfave/cli/v2"
)

// MainFlag
// main flags
func MainFlag() []cli.Flag {
	return []cli.Flag{
		//&cli.StringFlag{
		//	Name:  "clone-url",
		//	Usage: "Set git url to use by clone, if not set will use local repository",
		//	Value: "",
		//},
		&cli.StringFlag{
			Name:    "release-as",
			Usage:   fmt.Sprintf("Specify the release type manually (like npm version <major|minor|patch>) if not setting will use semver by history, if first release will change to %s", convention.DefaultSemverVersion),
			Aliases: []string{"r"},
		},
		&cli.StringFlag{
			Name:    "tag-prefix",
			Aliases: []string{"t"},
			Usage:   "Set a custom prefix for the git tag to be created",
			Value:   "v",
		},
		&cli.StringFlag{
			Name:    "infile",
			Aliases: []string{"i"},
			Usage:   "Read the CHANGELOG from this file",
			Value:   constant.DefaultChangelogMarkdownFile,
		},
		&cli.StringFlag{
			Name:    "outfile",
			Aliases: []string{"o"},
			Usage:   "Write the CHANGELOG to this file",
			Value:   constant.DefaultChangelogMarkdownFile,
		},
		&cli.StringFlag{
			Name:  "from-commit",
			Usage: "Generate the changelog from a specific tag commit full code. If not specified will use latest releaseCommitMessageFormat to find",
			Value: "",
		},
		&cli.BoolFlag{
			Name:  "auto-push",
			Usage: "enable auto git push after generating changelog, and will ignore --dry-run",
		},
	}
}

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Usage:   "open cli verbose mode",
			Value:   false,
			EnvVars: []string{constant.EnvKeyCliVerbose},
		},
		&cli.BoolFlag{
			Name:  "dry-run",
			Usage: "enable dry-run mode. this will not change any file and git",
			Value: true,
		},
		&cli.BoolFlag{
			Name:    "dry-run-disable",
			Usage:   "disable dry-run mode",
			Value:   false,
			EnvVars: []string{constant.EnvKeyDryRunDisable},
		},
	}
}

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    "config.timeout_second",
			Usage:   "command timeout setting second. default 10",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{constant.EnvKeyCliTimeoutSecond},
		},
		&cli.StringFlag{
			Name:   "git-remote",
			Usage:  "change git remote name. default origin",
			Value:  "origin",
			Hidden: true,
		},
	}
}
