package urfave_cli

import "github.com/urfave/cli/v2"

// UrfaveCliAppendCliFlag
// append cli.Flag
func UrfaveCliAppendCliFlag(target []cli.Flag, elem []cli.Flag) []cli.Flag {
	if len(elem) == 0 {
		return target
	}

	return append(target, elem...)
}
