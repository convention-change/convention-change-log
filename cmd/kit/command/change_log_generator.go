package command

import (
	"github.com/convention-change/convention-change-log/changelog"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/sinlov-go/go-git-tools/git"
	"github.com/sinlov-go/go-git-tools/git_info"
	"github.com/sinlov-go/sample-markdown/sample_mk"
)

type ChangeLogGenerator struct {
	ChangeLogGeneratorFunc `json:"-"`

	rootPath      string
	repoGitRemote string
	repository    git.Repository

	headBranchName  string
	gitRemoteInfo   *git_info.GitRemoteInfo
	changeLogReader changelog.Reader

	genCfg GenerateConfig
	spec   *convention.ConventionalChangeLogSpec

	historyFirstTagName        string
	gitHttpInfoDefault         convention.GitRepositoryHttpInfo
	latestCommits              []git.Commit
	changeLogHistoryNodes      []sample_mk.Node
	generateMarkdownNodes      []sample_mk.Node
	featNodes                  []sample_mk.Node
	changeLogNowNodes          []sample_mk.Node
	changelogNowWithTitleNodes []sample_mk.Node
	changelogDesc              changelog.ConventionalChangeLogDesc
}

func NewChangeLogGenerator(rootPath string) *ChangeLogGenerator {
	return &ChangeLogGenerator{
		rootPath: rootPath,
	}
}

type ChangeLogGeneratorFunc interface {
	LoadRepository(gitCloneUrl, remote string) error

	CheckRepository() error

	CheckWorktreeDirty() error

	GetHeadBranchName() string

	GetGitRemoteInfo() git_info.GitRemoteInfo

	ChangeLogInit(cfg GenerateConfig, spec *convention.ConventionalChangeLogSpec) error

	GetHistoryFirstTagName() string

	GetLatestCommits() []git.Commit

	GenerateCommitAsMdNodes() error

	CheckLocalFileChangeByArgs() error

	DryRun()

	DoChangeRepoFileByCommitLog() error

	DoGitOperator() error

	DryRunChangeVersion()

	ChangeVersion() error
}

type GenerateConfig struct {
	GitCloneUrl string

	GitInfoScheme string

	ReleaseAs  string
	TagPrefix  string
	ReleaseTag string

	Infile  string
	Outfile string

	FromCommit string

	AutoPush bool

	SkipWorktreeDirtyCheck bool
	IsOnlyChangeVersion    bool

	AppendMonoRepoPath []string
	AppendMonoRepoAll  bool
}
