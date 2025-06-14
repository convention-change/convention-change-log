package command

import (
	"fmt"
	"github.com/convention-change/convention-change-log/cmd/kit/constant"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/urfave/cli/v2"
	"strings"
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
		&cli.BoolFlag{
			Name:  "change-version",
			Usage: "only change version, by versionrc settings (1.10+)",
			Value: false,
		},

		&cli.StringSliceFlag{
			Name:  "append-monorepo",
			Usage: "append changelog to monorepo path, like: --append-monorepo=packages/a --append-monorepo=packages/b, this must define at config file .monorepo-pkg-path, (v1.11.+)",
		},

		&cli.BoolFlag{
			Name:  "append-monorepo-all",
			Usage: "append changelog to all monorepo path, this will ignore --append-monorepo config, (1.13.+)",
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
		&cli.StringFlag{
			Name:    "git-info-scheme",
			Usage:   fmt.Sprintf("git info scheme, only support: %s", strings.Join(gitInfoSchemeSupport, ", ")),
			Value:   "https",
			EnvVars: []string{constant.EnvKeyGitInfoScheme},
		},
		&cli.BoolFlag{
			Name:    "skip-worktree-check",
			Usage:   "skip git worktree dirty check",
			EnvVars: []string{constant.EnvKeySkipWorktreeCheck},
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
