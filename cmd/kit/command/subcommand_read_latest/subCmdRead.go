package subcommand_read_latest

import (
	"github.com/bar-counter/slog"
	"github.com/convention-change/convention-change-log/changelog"
	"github.com/convention-change/convention-change-log/cmd/kit/command"
	"github.com/convention-change/convention-change-log/cmd/kit/command/exit_cli"
	"github.com/convention-change/convention-change-log/cmd/kit/constant"
	"github.com/convention-change/convention-change-log/internal/urfave_cli"
	"github.com/gookit/color"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

const (
	commandName = "read-latest"
)

var commandEntry *ReadLatestCommand

type ReadLatestCommand struct {
	isDebug bool

	ReadChangeLogFile           string
	WriteLastChangeFileFullPath string
	isWriteLastChangeFile       bool
}

func (n *ReadLatestCommand) Exec() error {
	slog.Debugf("-> Exec subCommand [ %s ]", commandName)

	if !filepath_plus.PathExistsFast(n.ReadChangeLogFile) {
		return exit_cli.Format("can not find changelog at path: %s\n", n.ReadChangeLogFile)
	}

	changeLogSpec := command.CmdGlobalEntry().ChangeLogSpec
	reader, err := changelog.NewReader(n.ReadChangeLogFile, *changeLogSpec)
	if err != nil {
		return exit_cli.Err(err)
	}

	color.Greenf("=> Last change tag\n")
	color.Bluef("full tag: %s%s\n", changeLogSpec.TagPrefix, reader.HistoryFirstTagShort())
	color.Bluef("sort tag: %s\n", reader.HistoryFirstTagShort())
	color.Greenf("\n=> Last change title\n")
	color.Grayp(reader.HistoryFirstTitle())
	color.Greenf("\n\n=> Last change content\n")
	color.Grayp(reader.HistoryFirstContent())
	if reader.HistoryFirstChangeUrl() != "" {
		color.Greenf("\n\n=> Last change compare Url\n")
		color.Grayp(reader.HistoryFirstChangeUrl())
	}
	color.Println()

	if n.isWriteLastChangeFile {
		errWrite := filepath_plus.WriteFileByByte(n.WriteLastChangeFileFullPath, []byte(reader.HistoryFirstContent()), os.FileMode(0o666), true)
		if errWrite != nil {
			return exit_cli.Format("write last change to file: %s err: %v\n", n.WriteLastChangeFileFullPath, errWrite)
		}
		color.Greenf("\nWrite last change to file: %s\n", n.WriteLastChangeFileFullPath)
		color.Println()
	}

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "read-latest-file",
			Usage: "read change log file path",
			Value: constant.DefaultChangelogMarkdownFile,
		},
		&cli.StringFlag{
			Name:  "read-latest-out-path",
			Usage: "write last change file path",
			Value: constant.DefaultChangelogLastContentFile,
		},
		&cli.BoolFlag{
			Name:  "read-latest-out",
			Usage: "write last change to file at args --out-path",
		},
	}
}

func withEntry(c *cli.Context) (*ReadLatestCommand, error) {
	slog.Debugf("-> withEntry subCommand [ %s ]", commandName)

	globalEntry := command.CmdGlobalEntry()

	infile := c.String("read-latest-file")
	if infile == "" {
		return nil, exit_cli.Format("read change log file path is empty")
	}
	outfile := c.String("read-latest-out-path")
	if outfile == "" {
		return nil, exit_cli.Format("write last change file path is empty")
	}

	changeLogFile := filepath.Join(globalEntry.GitRootPath, infile)
	outLogFile := filepath.Join(globalEntry.GitRootPath, outfile)

	return &ReadLatestCommand{
		isDebug: globalEntry.Verbose,

		ReadChangeLogFile:           changeLogFile,
		WriteLastChangeFileFullPath: outLogFile,

		isWriteLastChangeFile: c.Bool("read-latest-out"),
	}, nil
}

func action(c *cli.Context) error {
	slog.Debugf("-> Sub Command action [ %s ] start", commandName)
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	commandEntry = entry
	return commandEntry.Exec()
}

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "read the latest change log or write latest change to file",
			Action: action,
			Flags:  flag(),
		},
	}
}
