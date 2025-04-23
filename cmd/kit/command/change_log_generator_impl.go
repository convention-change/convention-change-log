package command

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/bar-counter/slog"
	"github.com/convention-change/convention-change-log/changelog"
	"github.com/convention-change/convention-change-log/cmd/kit/command/exit_cli"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/convention-change/convention-change-log/internal/pkg_kit"
	goGit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/gookit/color"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	"github.com/sinlov-go/go-git-tools/git"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/sinlov-go/sample-markdown/sample_mk"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// LoadRepository
// load git repository info
func (c *ChangeLogGenerator) LoadRepository(gitCloneUrl, remote string) error {
	c.repoGitRemote = remote
	if gitCloneUrl == "" {

		repositoryOpen, errOpen := git.NewRepositoryRemoteByPath(remote, c.rootPath)
		if errOpen != nil {
			return exit_cli.Format("load local git repository error: %s", errOpen)
		}
		c.repository = repositoryOpen

	} else {
		repositoryClone, errClone := git.NewRepositoryRemoteClone(remote, memory.NewStorage(), nil, &goGit.CloneOptions{
			URL:        gitCloneUrl,
			RemoteName: remote,
		})
		if errClone != nil {
			return exit_cli.Format("clone git repository %s \nerror: %s", gitCloneUrl, errClone)
		}
		c.repository = repositoryClone
	}
	if c.repository == nil {
		return fmt.Errorf("can not load git repository")
	}

	return nil
}

func (c *ChangeLogGenerator) CheckRepository() error {
	headBranchName, err := c.repository.HeadBranchName()
	if err != nil {
		return fmt.Errorf("check repo err, can not get branch name by HEAD, please check now branch is in HEAD")
	}
	c.headBranchName = headBranchName

	gitRemoteInfo, err := c.repository.RemoteInfo(c.repoGitRemote, 0)
	if err != nil {
		return fmt.Errorf("check repo err, can not get remote info, error: %s", err)
	}

	c.gitRemoteInfo = gitRemoteInfo

	return nil
}

func (c *ChangeLogGenerator) CheckWorktreeDirty() error {
	submodules, errCheckHasSubmodules := c.repository.CheckHasSubmodules()
	if errCheckHasSubmodules == nil && submodules {
		dirty, errCheckSubmodulesIsDirty := c.repository.CheckSubmodulesIsDirty()
		if errCheckSubmodulesIsDirty != nil {
			color.Print(cmdErrorHelperCheckRepositoryInGitSubModule)
			color.Println("")
			color.Println("")
			return fmt.Errorf("check repo err, can not check submodules is dirty, error: %s", errCheckSubmodulesIsDirty)
		} else if dirty {
			color.Print(cmdErrorHelperCheckRepositoryInGitSubModule)
			color.Println("")
			color.Println("")
			return fmt.Errorf("check repo err, submodules is dirty, please commit submodules")
		}
	}
	var isWorktreeDirty bool
	var errWorktreeDirty error
	if c.repository.IsCitCmdAvailable() {
		isWorktreeDirty, errWorktreeDirty = c.repository.CheckWorkTreeIsDirtyWithGitCmd()
	} else {
		isWorktreeDirty, errWorktreeDirty = c.repository.CheckLocalBranchIsDirty()
	}

	if errWorktreeDirty != nil {
		return fmt.Errorf("check repo err, can not check worktree is dirty, error: %s", errWorktreeDirty)
	} else {
		if isWorktreeDirty {
			color.Print(cmdErrorHelperCheckRepositoryInNowIsDirty)
			color.Println("")
			color.Println("")
			return fmt.Errorf("check repo err, worktree is dirty, please check")
		}
	}
	return nil
}

func (c *ChangeLogGenerator) GetHeadBranchName() string {
	return c.headBranchName
}

func (c *ChangeLogGenerator) GetGitRemoteInfo() git_info.GitRemoteInfo {
	return *c.gitRemoteInfo
}

func (c *ChangeLogGenerator) ChangeLogInit(cfg GenerateConfig, spec *convention.ConventionalChangeLogSpec) error {

	c.genCfg = cfg
	c.spec = spec

	changeLogNodes := make([]sample_mk.Node, 0)
	reader, errHistory := changelog.NewReader(c.genCfg.Infile, *c.spec)
	if errHistory != nil && c.genCfg.ReleaseAs == "" {
		c.genCfg.ReleaseAs = convention.DefaultSemverVersion
		c.genCfg.ReleaseTag = fmt.Sprintf("%s%s", c.genCfg.TagPrefix, c.genCfg.ReleaseAs)
	} else {
		if len(reader.HistoryNodes()) > 0 {
			changeLogNodes = append(changeLogNodes, reader.HistoryNodes()...)
		}
	}
	c.changeLogReader = reader
	c.changeLogHistoryNodes = changeLogNodes

	historyFirstTagName := ""
	if c.genCfg.FromCommit == "" {
		if errHistory != nil {
			// this not find any tag and history
			c.genCfg.FromCommit = ""
		} else {
			historyFirstTag := reader.HistoryFirstTag()
			tagSearchByName, errTagSearchByName := c.repository.CommitTagSearchByName(historyFirstTag)
			if errTagSearchByName != nil {
				c.genCfg.FromCommit = ""
			} else {
				c.genCfg.FromCommit = tagSearchByName.Hash.String()
				historyFirstTagName = historyFirstTag
			}
		}
	}
	c.historyFirstTagName = historyFirstTagName

	latestCommits, errLatestCommits := c.repository.Log("", c.genCfg.FromCommit)
	if errLatestCommits != nil {
		return errLatestCommits
	}
	c.latestCommits = latestCommits

	gitInfoScheme := c.genCfg.GitInfoScheme
	if c.spec.CoverGitInfoScheme != "" {
		gitInfoScheme = c.spec.CoverGitInfoScheme
	}

	gitHttpHost := c.gitRemoteInfo.Host
	if c.spec.CoverHttpHost != "" {
		gitHttpHost = c.spec.CoverHttpHost
	}

	gitHttpInfoDefault := convention.GitRepositoryHttpInfo{
		Scheme:     gitInfoScheme,
		Host:       gitHttpHost,
		Owner:      c.gitRemoteInfo.User,
		Repository: c.gitRemoteInfo.Repo,
	}

	c.gitHttpInfoDefault = gitHttpInfoDefault

	return nil
}

func (c *ChangeLogGenerator) GetHistoryFirstTagName() string {
	return c.headBranchName
}

func (c *ChangeLogGenerator) GetLatestCommits() []git.Commit {
	return c.latestCommits
}

func (c *ChangeLogGenerator) GenerateCommitAsMdNodes() error {
	conventionCommits, errConvert := convention.ConvertGitCommits2ConventionCommits(c.latestCommits, *c.spec, c.gitHttpInfoDefault)
	if errConvert != nil {
		return fmt.Errorf("ConvertGitCommits2ConventionCommits err: %v", errConvert)
	}
	generateMarkdownNodes, featNodes, errConvert := changelog.GenerateMarkdownNodes(
		c.gitHttpInfoDefault,
		conventionCommits,
		*c.spec,
	)
	if errConvert != nil {
		return fmt.Errorf("GenerateMarkdownNodes err: %v", errConvert)
	}

	c.generateMarkdownNodes = generateMarkdownNodes
	c.featNodes = featNodes

	if c.genCfg.ReleaseAs != "" {
		// check cli settings ReleaseAs tag it exists
		oldReleaseTag, errOldReleaseTag := c.repository.CommitTagSearchByName(c.genCfg.ReleaseTag)
		if errOldReleaseTag != nil {
			slog.Debugf("not find tag: %s err: %v", c.genCfg.ReleaseTag, errOldReleaseTag)
		}
		if oldReleaseTag != nil {
			errReleaseTagExist := fmt.Errorf("want release tag is exist: %s, tag message is %s", c.genCfg.ReleaseTag, oldReleaseTag.Message)
			return errReleaseTagExist
		}
		// check version as semver
		_, errSemverCheck := semver.NewVersion(c.genCfg.ReleaseAs)
		if errSemverCheck != nil {
			return fmt.Errorf("NewVersion release-as is not semver err: %v", errSemverCheck)
		}
	} else {
		// find new version as semver by historyFirstTagName
		historyVersion, errHistorySemver := semver.NewVersion(c.historyFirstTagName)
		if errHistorySemver != nil {
			return fmt.Errorf("find new version as semver by historyFirstTagName err: %v", errHistorySemver)
		}
		var version semver.Version
		if len(featNodes) > 0 {
			version = historyVersion.IncMinor()
		} else {
			version = historyVersion.IncPatch()
		}
		c.genCfg.ReleaseAs = version.String()
		c.genCfg.ReleaseTag = fmt.Sprintf("%s%s", c.genCfg.TagPrefix, c.genCfg.ReleaseAs)
	}

	changelogDesc := changelog.ConventionalChangeLogDesc{
		Version:      c.genCfg.ReleaseTag,
		When:         time.Now(),
		Location:     time.Local,
		ToolsKitName: KitName,
		ToolsKitURL:  KitUrl,
	}
	if c.changeLogReader.HistoryFirstTagShort() != "" {
		changelogDesc.PreviousTag = c.changeLogReader.HistoryFirstTagShort()
	}

	c.changelogDesc = changelogDesc

	nodesGenerateWithTitle, errAddTitle := changelog.AddMarkdownChangelogNodesTitle(
		generateMarkdownNodes,
		c.gitHttpInfoDefault,
		changelogDesc,
		*c.spec,
	)
	c.changeLogNowNodes = nodesGenerateWithTitle

	if errAddTitle != nil {
		return fmt.Errorf("AddMarkdownChangelogNodesTitle err: %v", errAddTitle)
	}

	changelogNodesWithHead, err := changelog.AddMarkdownChangelogNodesHead(nodesGenerateWithTitle, changelogDesc, *c.spec)
	if err != nil {
		return fmt.Errorf("AddMarkdownChangelogNodesHead err: %v", err)
	}
	c.changelogNowWithTitleNodes = changelogNodesWithHead

	return nil
}

func (c *ChangeLogGenerator) DryRun() {

	latestMarkdownContent := sample_mk.GenerateText(c.changelogNowWithTitleNodes)
	color.Printf(cmdHelpDryRunOutputting, c.genCfg.Outfile)
	color.Println("")

	color.Println(LogLineSpe)
	color.Grayf("%s\n", latestMarkdownContent)
	color.Println(LogLineSpe)

	color.Printf(cmdHelpDryRunCommitting, c.genCfg.Infile)
	color.Println("")
	color.Printf(cmdHelpDryRunTagRelease, c.genCfg.ReleaseTag)
	color.Println("")
	color.Println("")
	color.Print(cmdHelpFinishDryRun)
	color.Println("")
	color.Println("")

	color.Print(cmdHelperRepositorySafeTitle)
	color.Println()
	if c.headBranchName == "main" {
		color.Print(cmdHelperRepositorySafeFormBranchMain)
		color.Println()
	} else {
		color.Print(cmdHelperRepositorySafeFormBranchNow)
		color.Println()
	}
	color.Println()

	c.DryRunChangeVersion()
}

func (c *ChangeLogGenerator) DoChangeRepoFileByCommitLog() error {
	// add history
	c.changeLogHistoryNodes = append(c.changelogNowWithTitleNodes, c.changeLogHistoryNodes...)

	fullMarkdownContent := sample_mk.GenerateText(c.changeLogHistoryNodes)

	errChangeLocalFile := c.changeRepoLocalFiles(fullMarkdownContent)
	if errChangeLocalFile != nil {
		return fmt.Errorf("changeRepoLocalFiles err: %v", errChangeLocalFile)
	}

	errAppendChangeLocalFile := c.appendMonoRepoFiles(c.changeLogNowNodes, c.changeLogHistoryNodes)
	if errAppendChangeLocalFile != nil {
		return fmt.Errorf("appendMonoRepoFiles err: %v", errAppendChangeLocalFile)
	}

	return nil
}

func (c *ChangeLogGenerator) CheckLocalFileChangeByArgs() error {
	// if open --append-monorepo-all will replace all mono-repo-pkg-path
	if c.genCfg.AppendMonoRepoAll {
		c.genCfg.AppendMonoRepoPath = c.spec.MonoRepoPkgPathList
		slog.Info("now use all path by setting `.monorepo-pkg-path at config file.")
		return nil
	}
	if len(c.genCfg.AppendMonoRepoPath) > 0 {
		for _, appendPath := range c.genCfg.AppendMonoRepoPath {
			if !string_tools.StringInArr(appendPath, c.spec.MonoRepoPkgPathList) {
				return fmt.Errorf("args [ --append-monorepo %s ] not in config { .monorepo-pkg-path } list", appendPath)
			}
		}
	}

	return nil
}

func (c *ChangeLogGenerator) appendMonoRepoFiles(newNodes []sample_mk.Node, oldNodes []sample_mk.Node) error {

	if len(c.genCfg.AppendMonoRepoPath) == 0 {
		return nil
	}

	headMarkdownContent := sample_mk.GenerateText(newNodes)

	color.Magentaf("will append change log version to ( %s ) to file:\n", c.genCfg.ReleaseAs)

	for _, appendPath := range c.genCfg.AppendMonoRepoPath {
		if !string_tools.StringInArr(appendPath, c.spec.MonoRepoPkgPathList) {
			color.Warnf("append mono-repo-path %s not in spec [ monorepo-pkg-path ] list, ignore", appendPath)
			continue
		}
		monoRepoChangeLogPath := filepath.Join(c.rootPath, appendPath, c.genCfg.Outfile)
		appendContent := []byte(headMarkdownContent + "\n")
		if filepath_plus.PathExistsFast(monoRepoChangeLogPath) {
			errAppendChangeHead := filepath_plus.AppendFileHead(monoRepoChangeLogPath, appendContent)
			if errAppendChangeHead != nil {
				return fmt.Errorf("append change log head to file failed, %v", errAppendChangeHead)
			}
		} else {
			errWriteFile := filepath_plus.WriteFileByByte(monoRepoChangeLogPath, appendContent, os.FileMode(0o666), true)
			if errWriteFile != nil {
				return fmt.Errorf("WriteFileByByte err: %v", errWriteFile)
			}
		}
		color.Greenf("append change log head at: %s\n", monoRepoChangeLogPath)
	}
	return nil
}

func (c *ChangeLogGenerator) changeRepoLocalFiles(fullChangeLogContent string) error {
	errWriteFile := filepath_plus.WriteFileByByte(c.genCfg.Outfile, []byte(fullChangeLogContent), os.FileMode(0o666), true)
	if errWriteFile != nil {
		return fmt.Errorf("WriteFileByByte err: %v", errWriteFile)
	}

	errChangeVersion := c.ChangeVersion()
	if errChangeVersion != nil {
		return errChangeVersion
	}

	return nil
}

func (c *ChangeLogGenerator) DoGitOperator() error {

	errDoGit := c.doGit(c.headBranchName)
	if errDoGit != nil {
		color.Print(cmdHelpGitCommitFail)
		color.Println("")
		color.Print(cmdHelpGitCommitFixHead)
		color.Println("")
		color.Println("")
		color.Print(cmdHelpGitCommitCheckBranch)
		color.Println("")
		color.Printf(cmdHelpGitPushTryAgain, c.headBranchName)
		color.Println("")
		color.Print(cmdHelpGitPushFailHint)
		color.Println("")
		color.Println("")
		color.Print(cmdHelpGitCommitErrorHint)
		color.Println("")
		color.Printf(cmdHelpGitCommitFixTag, c.genCfg.ReleaseTag)
		color.Println("")
		color.Print(cmdHelpGitCommitResetSoft)
		color.Println("")
		color.Println("")
		return errDoGit
	}

	return nil
}

func (c *ChangeLogGenerator) doGit(branchName string) error {
	// disable git-go repository issues until https://github.com/go-git/go-git/issues/180 is fixed

	cmdOutput, err := exec.Command("git", "add", "--all").CombinedOutput()
	if err != nil {
		return err
	}
	slog.Debugf("git add output:\n%s", cmdOutput)

	color.Printf(cmdHelpOutputting, c.genCfg.Outfile)
	color.Println("")
	color.Printf(cmdHelpCommitting, c.genCfg.Infile)
	color.Println("")

	releaseCommit := new(convention.ReleaseCommitMessageRenderTemplate)
	releaseCommit.CurrentTag = c.genCfg.ReleaseAs
	releaseCommitMsg, errRender := convention.RaymondRender(c.spec.ReleaseCommitMessageFormat, releaseCommit)
	if errRender != nil {
		return errRender
	}

	cmdOutput, err = exec.Command("git", "commit", "-m", releaseCommitMsg).CombinedOutput()
	if err != nil {
		slog.Errorf(err, "git commit output:\n%s", cmdOutput)
		return err
	}
	slog.Debugf("git commit output:\n%s", cmdOutput)

	color.Printf(cmdHelpTagRelease, c.genCfg.ReleaseTag)
	color.Println("")

	cmdOutput, err = exec.Command("git", "tag", c.genCfg.ReleaseTag, "-m", releaseCommitMsg).CombinedOutput()
	if err != nil {
		slog.Errorf(err, "git tag output:\n%s", cmdOutput)
		return err
	}
	slog.Debugf("git tag output:\n%s", cmdOutput)

	if c.genCfg.AutoPush {
		cmdOutput, err = exec.Command("git", "push", "--follow-tags", "origin", branchName).CombinedOutput()
		if err != nil {
			slog.Error("git push error", err)
			return err
		}

		slog.Debugf("git push output:\n%s", cmdOutput)
		color.Printf(cmdHelpFinishGitPush, branchName)
		color.Println("")
		color.Printf(cmdHelpHasTagRelease, c.genCfg.ReleaseTag)
		color.Println("")
		return nil
	}

	color.Printf(cmdHelpGitPushRun, branchName)
	color.Println("")
	return nil
}

func (c *ChangeLogGenerator) DryRunChangeVersion() {
	if len(c.spec.MonoRepoPkgPathList) > 0 {
		// try update monorepo pkg list
		color.Magentaf("will update [ monorepo-pkg-path ] package.json version to ( %s )in file list:\n", c.genCfg.ReleaseAs)
		for _, pkgPath := range c.spec.MonoRepoPkgPathList {
			// replace file line by regexp
			subModulePkgJsonPath := filepath.Join(c.rootPath, pkgPath, "package.json")
			color.Greenf("%s\n", subModulePkgJsonPath)
			if !filepath_plus.PathExistsFast(subModulePkgJsonPath) {
				slog.Warnf("not find update monorepo-pkg-path package.json path: %s", subModulePkgJsonPath)
				continue
			}
		}
	}

	if len(c.genCfg.AppendMonoRepoPath) > 0 {
		color.Magentaf("will append change log version to ( %s ) to file:\n", c.genCfg.ReleaseAs)
		for _, appendPath := range c.genCfg.AppendMonoRepoPath {
			if !string_tools.StringInArr(appendPath, c.spec.MonoRepoPkgPathList) {
				color.Warnf("append mono-repo-path %s not in spec [ monorepo-pkg-path ] list, ignore", appendPath)
				continue
			}
			subModuleChangeLogPath := filepath.Join(c.rootPath, appendPath, c.genCfg.Outfile)
			color.Greenf("%s\n", subModuleChangeLogPath)
		}
	}

	color.Println()
}

func (c *ChangeLogGenerator) ChangeVersion() error {

	if c.genCfg.ReleaseAs != "" {
		// try update node
		pkgJsonPath := filepath.Join(c.rootPath, "package.json")
		if filepath_plus.PathExistsFast(pkgJsonPath) {
			// replace file line by regexp
			slog.Debugf("try update node version in file: %s", pkgJsonPath)
			err := pkg_kit.ReplaceJsonVersionByLine(pkgJsonPath, c.genCfg.ReleaseAs)
			if err != nil {
				slog.Error("ReplaceJsonVersionByLine", err)
			}
			pkgJsonLockPath := filepath.Join(c.rootPath, "package-lock.json")
			if filepath_plus.PathExistsFast(pkgJsonLockPath) {
				output, errNpmInstall := exec.Command("npm", "install").CombinedOutput()
				if errNpmInstall != nil {
					slog.Error("do Command npm install error", errNpmInstall)
				}
				slog.Debugf("npm install output:\n%s", output)
			}
			slog.Infof("update root package.json version ( %s )  in file: %s", c.genCfg.ReleaseAs, pkgJsonPath)
		}

		// try update monorepo
		if len(c.spec.MonoRepoPkgPathList) > 0 {
			// try update monorepo pkg list
			for _, pkgPath := range c.spec.MonoRepoPkgPathList {
				// replace file line by regexp
				subModulePkgJsonPath := filepath.Join(c.rootPath, pkgPath, "package.json")
				slog.Debugf("try update submodule package.json version in file: %s", subModulePkgJsonPath)
				if !filepath_plus.PathExistsFast(subModulePkgJsonPath) {
					slog.Warnf("not find update submodule package.json path: %s", subModulePkgJsonPath)
					continue
				}
				slog.Infof("update mono-repo package.json version ( %s )  in file: %s", c.genCfg.ReleaseAs, subModulePkgJsonPath)
				err := pkg_kit.ReplaceJsonVersionByLine(subModulePkgJsonPath, c.genCfg.ReleaseAs)
				if err != nil {
					return fmt.Errorf("submodule package.json change ReplaceJsonVersionByLine %v", err)
				}
			}
		}
	}

	return nil
}
