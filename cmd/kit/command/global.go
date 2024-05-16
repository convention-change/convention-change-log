package command

import (
	"encoding/json"
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/convention-change/convention-change-log/cmd/kit/command/exit_cli"
	"github.com/convention-change/convention-change-log/cmd/kit/constant"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/convention-change/convention-change-log/internal/log"
	"github.com/convention-change/convention-change-log/internal/pkgJson"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strings"
)

type GlobalConfig struct {
	LogLevel      string
	TimeoutSecond uint
}

type (
	// GlobalCommand
	//	command root
	GlobalCommand struct {
		Name    string
		Version string
		Verbose bool
		DryRun  bool

		GitRootPath   string
		GitRemote     string
		ChangeLogSpec *convention.ConventionalChangeLogSpec

		RootCfg GlobalConfig

		GenerateConfig GenerateConfig
	}
)

var (
	cmdGlobalEntry *GlobalCommand
)

// CmdGlobalEntry
//
//	return global command entry
func CmdGlobalEntry() *GlobalCommand {
	return cmdGlobalEntry
}

// globalExec
//
//	do global command exec
func (c *GlobalCommand) globalExec() error {

	slog.Debug("-> start GlobalAction")

	if c.GenerateConfig.AutoPush {
		c.DryRun = false
		slog.Info("auto push is enable, will ignore --dry-run")
	}

	_, err := git_info.IsPathGitManagementRoot(c.GitRootPath)
	if err != nil {
		return exit_cli.Format("cli run path not git repository root, please check path at: %s", c.GitRootPath)
	}

	clGenerator := NewChangeLogGenerator(c.GitRootPath)

	errLoadRepository := clGenerator.LoadRepository(c.GenerateConfig.GitCloneUrl, c.GitRemote)
	if errLoadRepository != nil {
		return exit_cli.Err(errLoadRepository)
	}
	errCheckRepository := clGenerator.CheckRepository()
	if errCheckRepository != nil {
		return exit_cli.Err(errCheckRepository)
	}

	if c.Verbose {
		bytes, errJson := json.Marshal(clGenerator.GetGitRemoteInfo())
		if errJson != nil {
			slog.Error("git remote info Marshal err: %v", errJson)
			return exit_cli.Err(errJson)
		}
		slog.Debugf("gitRemoteInfo:\n%s", string(bytes))
	}

	errChangeLogInit := clGenerator.ChangeLogInit(c.GenerateConfig, c.ChangeLogSpec)
	if errChangeLogInit != nil {
		return exit_cli.Err(errChangeLogInit)
	}
	slog.Debugf("historyFirstTagName: %s", clGenerator.GetHistoryFirstTagName())
	slog.Debugf("c.GenerateConfig.FromCommit: %s", c.GenerateConfig.FromCommit)
	slog.Debugf("latestCommits len %d", len(clGenerator.GetLatestCommits()))

	errGenerateCommitNodes := clGenerator.GenerateCommitAsMdNodes()
	if errGenerateCommitNodes != nil {
		return exit_cli.Err(errGenerateCommitNodes)
	}

	if c.DryRun {
		clGenerator.DryRun()
		return nil
	}

	errDoChangeRepoFileByCommitLog := clGenerator.DoChangeRepoFileByCommitLog()
	if errDoChangeRepoFileByCommitLog != nil {
		return exit_cli.Err(errDoChangeRepoFileByCommitLog)
	}
	errDoGitOperator := clGenerator.DoGitOperator()
	if errDoGitOperator != nil {
		slog.Errorf(errDoGitOperator, "git operator is error")
		return errDoGitOperator
	}

	return nil
}

// withGlobalFlag
//
// bind global flag to globalExec
func withGlobalFlag(c *cli.Context, cliVersion, cliName string) (*GlobalCommand, error) {
	slog.Debug("-> withGlobalFlag")

	isVerbose := c.Bool("verbose")

	dir, err := os.Getwd()
	if err != nil {
		return nil, exit_cli.Err(err)
	}
	gitRootFolder := dir

	config := GlobalConfig{
		LogLevel:      c.String("config.log_level"),
		TimeoutSecond: c.Uint("config.timeout_second"),
	}

	// c.String("clone-url") close this way to get latest tag

	tagPrefix := c.String("tag-prefix")
	cliReleaseAs := c.String("release-as")
	cliReleaseTag := ""
	if cliReleaseAs != "" {
		cliReleaseTag = fmt.Sprintf("%s%s", tagPrefix, cliReleaseAs)
	}

	isAutoPush := c.Bool("auto-push")
	gitInfoScheme := c.String("git-info-scheme")

	if !string_tools.StringInArr(gitInfoScheme, gitInfoSchemeSupport) {
		return nil, exit_cli.Format("--git-info-scheme only support %s", strings.Join(gitInfoSchemeSupport, ", "))
	}

	generateConfig := GenerateConfig{
		GitCloneUrl:   "",
		GitInfoScheme: gitInfoScheme,
		ReleaseAs:     cliReleaseAs,
		TagPrefix:     tagPrefix,
		ReleaseTag:    cliReleaseTag,

		Infile:  c.String("infile"),
		Outfile: c.String("outfile"),

		FromCommit: c.String("from-commit"),

		AutoPush: isAutoPush,
	}

	specFilePath := filepath.Join(gitRootFolder, constant.VersionRcFileName)
	changeLogSpec, err := convention.LoadConventionalChangeLogSpecByPath(specFilePath)
	if err != nil {
		return nil, exit_cli.Err(err)
	}

	changeLogSpec.TagPrefix = tagPrefix

	isDryRun := c.Bool("dry-run")
	isDryRunDisable := c.Bool("dry-run-disable")
	if isDryRunDisable {
		isDryRun = false
	}

	p := GlobalCommand{
		Name:    cliName,
		Version: cliVersion,
		Verbose: isVerbose,
		DryRun:  isDryRun,

		GitRootPath:   gitRootFolder,
		GitRemote:     c.String("git-remote"),
		ChangeLogSpec: changeLogSpec,

		RootCfg:        config,
		GenerateConfig: generateConfig,
	}
	return &p, nil
}

// GlobalBeforeAction
// do command Action before flag global.
func GlobalBeforeAction(c *cli.Context) error {
	isVerbose := c.Bool("verbose")
	err := log.InitLog(isVerbose, !isVerbose)
	if err != nil {
		panic(err)
	}
	cliVersion := pkgJson.GetPackageJsonVersionGoStyle(false)
	if isVerbose {
		slog.Warnf("-> open verbose, and now command version is: %s", cliVersion)
	}
	appName := pkgJson.GetPackageJsonName()
	cmdGlobalEntry, err = withGlobalFlag(c, cliVersion, appName)
	if err != nil {
		return err
	}

	return nil
}

// GlobalAction
// do command Action flag.
func GlobalAction(c *cli.Context) error {
	if cmdGlobalEntry == nil {
		panic(fmt.Errorf("not init GlobalBeforeAction success to new cmdGlobalEntry"))
	}

	isVerbose := c.Bool("verbose")
	if isVerbose {
		slog.Infof("-> start run command: %s, version %s", cmdGlobalEntry.Name, cmdGlobalEntry.Version)
	}

	err := cmdGlobalEntry.globalExec()
	if err != nil {
		return err
	}
	return nil
}

// GlobalAfterAction
//
//	do command Action after flag global.
//
//nolint:golint,unused
func GlobalAfterAction(c *cli.Context) error {
	isVerbose := c.Bool("verbose")
	if isVerbose {
		slog.Infof("-> finish run command: %s, version %s", cmdGlobalEntry.Name, cmdGlobalEntry.Version)
	}
	return nil
}
