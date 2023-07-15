package urfave_cli

import "github.com/urfave/cli/v2"

func UrfaveCliAppendCliCommand(target []*cli.Command, elem []*cli.Command) []*cli.Command {
	if len(elem) == 0 {
		return target
	}
	return append(target, elem...)
}
