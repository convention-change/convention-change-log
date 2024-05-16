package command

import "github.com/gookit/color"

const (
	KitName    = "convention-change-log"
	KitUrl     = "https://github.com/convention-change/convention-change-log"
	LogLineSpe = "---"
)

var (
	cmdOkEmoji       = color.Green.Render("✔")
	cmdOkDryRunEmoji = color.Gray.Render("✔")
	cmdInfoEmoji     = color.Blue.Render("ℹ")
	cmdWarnEmoji     = color.Warn.Render("ℹ")
	cmdErrorEmoji    = color.Red.Render("ℹ")

	cmdHelpOutputting       = cmdOkEmoji + " outputting changes to %s"
	cmdHelpDryRunOutputting = cmdOkDryRunEmoji + " outputting changes to %s"
	cmdHelpCommitting       = cmdOkEmoji + " committing %s"
	cmdHelpDryRunCommitting = cmdOkDryRunEmoji + " committing %s"
	cmdHelpTagRelease       = cmdOkEmoji + " tagging release %s"
	cmdHelpDryRunTagRelease = cmdOkDryRunEmoji + " tagging release %s"

	cmdHelpFinishDryRun = cmdOkDryRunEmoji + " finish dry run, preview operation will not take effect, only the new collected change log will be printed."

	cmdHelpGitPushRun    = cmdInfoEmoji + " Run `git push --follow-tags origin %s` to publish"
	cmdHelpFinishGitPush = cmdOkEmoji + " Finish `git push --follow-tags origin %s`"
	cmdHelpHasTagRelease = cmdOkEmoji + " has release %s"

	cmdHelpGitCommitFail        = cmdWarnEmoji + " try git command push fail, please check!"
	cmdHelpGitCommitFixHead     = cmdInfoEmoji + " you can use this command to check now"
	cmdHelpGitCommitCheckBranch = cmdInfoEmoji + " git branch -vv"
	cmdHelpGitPushTryAgain      = cmdInfoEmoji + " try `git push --follow-tags origin %s` to publish"
	cmdHelpGitPushFailHint      = cmdWarnEmoji + " if push error show [ error: failed to push ] , please check git remote protection settings"
	cmdHelpGitCommitErrorHint   = cmdInfoEmoji + " if commit error can use this command fix"
	cmdHelpGitCommitFixTag      = cmdWarnEmoji + " git tag --delete %s"
	cmdHelpGitCommitResetSoft   = cmdWarnEmoji + " git reset --soft HEAD^"

	cmdHelperRepositorySafeTitle          = cmdWarnEmoji + " ensure that the tag has been synchronized when it is created, please confirm it first."
	cmdHelperRepositorySafeFormBranchMain = cmdInfoEmoji + " run: git status --porcelain ; git checkout $(git config --get init.defaultBranch) ; git pull --verbose; git fetch --tags"
	cmdHelperRepositorySafeFormBranchNow  = cmdInfoEmoji + " run: git status --porcelain ; git pull --verbose; git fetch --tags"

	cmdErrorHelperCheckRepositoryInGitSubModule = cmdErrorEmoji + " try fix as: git submodule update --init --recursive"
	cmdErrorHelperCheckRepositoryInNowIsDirty   = cmdErrorEmoji + " please commit then git pull, then run check as: git status --porcelain"
)
