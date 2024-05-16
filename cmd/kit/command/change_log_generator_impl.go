package command

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/bar-counter/slog"
	"github.com/convention-change/convention-change-log/changelog"
	"github.com/convention-change/convention-change-log/cmd/kit/command/exit_cli"
	"github.com/convention-change/convention-change-log/cmd/kit/constant"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/convention-change/convention-change-log/internal/pkgJson"
	goGit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/gookit/color"
	"github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
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
	c.changeLogNodes = changeLogNodes

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

	gitHttpHost := c.gitRemoteInfo.Host
	if c.spec.CoverHttpHost != "" {
		gitHttpHost = c.spec.CoverHttpHost
	}

	gitHttpInfoDefault := convention.GitRepositoryHttpInfo{
		Scheme:     c.genCfg.GitInfoScheme,
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
		ToolsKitName: constant.KitName,
		ToolsKitURL:  constant.KitUrl,
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

	if errAddTitle != nil {
		return fmt.Errorf("AddMarkdownChangelogNodesTitle err: %v", errAddTitle)
	}

	changelogNodesWithHead, err := changelog.AddMarkdownChangelogNodesHead(nodesGenerateWithTitle, changelogDesc, *c.spec)
	if err != nil {
		return fmt.Errorf("AddMarkdownChangelogNodesHead err: %v", err)
	}
	c.changelogNodesWithHead = changelogNodesWithHead

	return nil
}

func (c *ChangeLogGenerator) DryRun() {
	latestMarkdownContent := sample_mk.GenerateText(c.changelogNodesWithHead)
	color.Printf(constant.CmdHelpOutputting, c.genCfg.Outfile)
	color.Println("")

	color.Println(constant.LogLineSpe)
	color.Grayf("%s\n", latestMarkdownContent)
	color.Println(constant.LogLineSpe)

	color.Printf(constant.CmdHelpCommitting, c.genCfg.Infile)
	color.Println("")
	color.Printf(constant.CmdHelpTagRelease, c.genCfg.ReleaseTag)
	color.Println("")
	color.Printf(constant.CmdHelpFinishDryRun)
	color.Println("")
}

func (c *ChangeLogGenerator) DoChangeRepoFileByCommitLog() error {

	// add history
	c.changeLogNodes = append(c.changelogNodesWithHead, c.changeLogNodes...)

	fullMarkdownContent := sample_mk.GenerateText(c.changeLogNodes)

	errChangeLocalFile := c.changeRepoLocalFiles(fullMarkdownContent)
	if errChangeLocalFile != nil {
		return fmt.Errorf("changeRepoLocalFiles err: %v", errChangeLocalFile)
	}

	return nil
}

func (c *ChangeLogGenerator) changeRepoLocalFiles(fullChangeLogContent string) error {
	errWriteFile := filepath_plus.WriteFileByByte(c.genCfg.Outfile, []byte(fullChangeLogContent), os.FileMode(0766), true)
	if errWriteFile != nil {
		return fmt.Errorf("WriteFileByByte err: %v", errWriteFile)
	}

	if c.genCfg.ReleaseAs != "" {
		// try update node
		pkgJsonPath := filepath.Join(c.rootPath, "package.json")
		if filepath_plus.PathExistsFast(pkgJsonPath) {
			// replace file line by regexp
			slog.Debugf("try update node version in file: %s", pkgJsonPath)
			err := pkgJson.ReplaceJsonVersionByLine(pkgJsonPath, c.genCfg.ReleaseAs)
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
				err := pkgJson.ReplaceJsonVersionByLine(pkgJsonPath, c.genCfg.ReleaseAs)
				if err != nil {
					slog.Error("submodule package.json version ReplaceJsonVersionByLine", err)
				}
			}
		}
	}

	return nil
}

func (c *ChangeLogGenerator) DoGitOperator() error {

	errDoGit := c.doGit(c.headBranchName)
	if errDoGit != nil {
		color.Printf(constant.CmdHelpGitCommitFail)
		color.Println("")
		color.Printf(constant.CmdHelpGitCommitFixHead)
		color.Println("")
		color.Println("")
		color.Printf(constant.CmdHelpGitCommitCheckBranch)
		color.Println("")
		color.Printf(constant.CmdHelpGitPushTryAgain, c.headBranchName)
		color.Println("")
		color.Printf(constant.CmdHelpGitPushFailHint)
		color.Println("")
		color.Println("")
		color.Printf(constant.CmdHelpGitCommitErrorHint)
		color.Println("")
		color.Printf(constant.CmdHelpGitCommitFixTag, c.genCfg.ReleaseTag)
		color.Println("")
		color.Printf(constant.CmdHelpGitCommitResetSoft)
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

	color.Printf(constant.CmdHelpOutputting, c.genCfg.Outfile)
	color.Println("")
	color.Printf(constant.CmdHelpCommitting, c.genCfg.Infile)
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

	color.Printf(constant.CmdHelpTagRelease, c.genCfg.ReleaseTag)
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
		color.Printf(constant.CmdHelpFinishGitPush, branchName)
		color.Println("")
		color.Printf(constant.CmdHelpHasTagRelease, c.genCfg.ReleaseTag)
		color.Println("")
		return nil
	}

	color.Printf(constant.CmdHelpGitPushRun, branchName)
	color.Println("")
	return nil
}
