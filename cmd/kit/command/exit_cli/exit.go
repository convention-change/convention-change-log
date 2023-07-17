package exit_cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

const exitCodeDefault = 127

var exitCode = exitCodeDefault

func ChangeDefaultExitCode(code int) {
	exitCode = code
}

func Format(format string, a ...any) cli.ExitCoder {
	return cli.Exit(fmt.Sprintf(format, a...), exitCode)
}

func FormatCode(code int, format string, a ...any) cli.ExitCoder {
	return cli.Exit(fmt.Sprintf(format, a...), code)
}

func Err(err error) cli.ExitCoder {
	return cli.Exit(err.Error(), exitCode)
}

func ErrCode(code int, err error) cli.ExitCoder {
	return cli.Exit(err.Error(), code)
}

func ErrMsg(err error, msg string) cli.ExitCoder {
	return cli.Exit(fmt.Sprintf("%s err: %s", msg, err), exitCode)
}

func ErrMsgCode(code int, err error, msg string) cli.ExitCoder {
	return cli.Exit(fmt.Sprintf("%s err: %s", msg, err), code)
}
