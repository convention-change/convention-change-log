package constant

import "github.com/gookit/color"

const (
	CopyrightStartYear = "2023"

	// EnvKeyCliVerbose
	//	Provides the debug flag. This value is true when the command is open debug mode
	EnvKeyCliVerbose = "CLI_VERBOSE"

	// EnvKeyCliTimeoutSecond
	//	Provides the timeout second flag
	EnvKeyCliTimeoutSecond = "CLI_CONFIG_TIMEOUT_SECOND"

	// EnvLogLevel
	//	env ENV_WEB_LOG_LEVEL default ""
	EnvLogLevel string = "CLI_LOG_LEVEL"

	DefaultChangelogMarkdownFile    = "CHANGELOG.md"
	DefaultChangelogLastContentFile = "CHANGELOG.txt"

	VersionRcFileName = ".versionrc"

	KitName = "convention-change-log"
	KitUrl  = "https://github.com/convention-change/convention-change-log"

	LogLineSpe = "---"
)

var (
	CmdOkEmoji           = color.Green.Render("✔")
	CmdInfoEmoji         = color.Blue.Render("ℹ")
	CmdWarnEmoji         = color.Warn.Render("ℹ")
	CmdHelpOutputting    = CmdOkEmoji + " outputting changes to %s"
	CmdHelpCommitting    = CmdOkEmoji + " committing %s"
	CmdHelpTagRelease    = CmdOkEmoji + " tagging release %s"
	CmdHelpGitPush       = CmdInfoEmoji + " Run `git push --follow-tags origin %s` to publish"
	CmdHelpFinishGitPush = CmdOkEmoji + " Finish `git push --follow-tags origin %s`"

	CmdHelpGitCommitFail        = CmdWarnEmoji + " must check git commit !"
	CmdHelpGitCommitFixHead     = CmdWarnEmoji + " can use this command fix"
	CmdHelpGitCommitCheckStatus = CmdWarnEmoji + " git status"
	CmdHelpGitCommitFixTag      = CmdWarnEmoji + " git tag --delete %s"
	CmdHelpGitCommitResetSoft   = CmdWarnEmoji + " git reset --soft HEAD^"
)
