package constant

import "github.com/gookit/color"

const (
	CopyrightStartYear = "2023"

	// EnvKeyCliVerbose
	//	Provides the debug flag. This value is true when the command is open debug mode
	EnvKeyCliVerbose = "CLI_VERBOSE"

	EnvKeyDryRunDisable = "CLI_DRY_RUN_DISABLE"

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
	CmdOkEmoji          = color.Green.Render("✔")
	CmdInfoEmoji        = color.Blue.Render("ℹ")
	CmdWarnEmoji        = color.Warn.Render("ℹ")
	CmdHelpFinishDryRun = CmdOkEmoji + " finish dry run"

	CmdHelpOutputting = CmdOkEmoji + " outputting changes to %s"
	CmdHelpCommitting = CmdOkEmoji + " committing %s"
	CmdHelpTagRelease = CmdOkEmoji + " tagging release %s"

	CmdHelpGitPushRun    = CmdInfoEmoji + " Run `git push --follow-tags origin %s` to publish"
	CmdHelpFinishGitPush = CmdOkEmoji + " Finish `git push --follow-tags origin %s`"
	CmdHelpHasTagRelease = CmdOkEmoji + " has release %s"

	CmdHelpGitCommitFail        = CmdWarnEmoji + " try git command push fail, please check!"
	CmdHelpGitCommitFixHead     = CmdInfoEmoji + " you can use this command to check now"
	CmdHelpGitCommitCheckBranch = CmdInfoEmoji + " git branch -vv"
	CmdHelpGitPushTryAgain      = CmdInfoEmoji + " try `git push --follow-tags origin %s` to publish"
	CmdHelpGitPushFailHint      = CmdWarnEmoji + " if push error show [ error: failed to push ] , please check git remote protection settings"
	CmdHelpGitCommitErrorHint   = CmdInfoEmoji + " if commit error can use this command fix"
	CmdHelpGitCommitFixTag      = CmdWarnEmoji + " git tag --delete %s"
	CmdHelpGitCommitResetSoft   = CmdWarnEmoji + " git reset --soft HEAD^"
)
