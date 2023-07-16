package convention

import (
	"fmt"
	"strings"

	"github.com/sinlov-go/go-git-tools/git"
)

// https://www.conventionalcommits.org/en/v1.0.0/
// <type>[optional scope]: <description>
// [optional body]
// [optional footer(s)]
// BREAKING CHANGE: <breaking change description>
// [issueReference] [issuePrefix]<issueId>

type BreakingChanges struct {
	Describe          string
	IssueReference    string
	IssuePrefix       string
	IssueReferencesId uint64
}

// Commit conventional commit
type Commit struct {
	// Commit as is
	RawHeader string

	Type  string
	Scope string

	BreakingChanges BreakingChanges
}

// NewCommitWithLogSpec
//
//	c git.Commit
//	spec ConventionalChangeLogSpec
//	gitRepoUrl git repo url if not set will not contain Hash link
//	return conventional commit from git commit
func NewCommitWithLogSpec(c git.Commit, spec ConventionalChangeLogSpec, gitRepoUrl string) (Commit, error) {
	result, err := NewCommitWithOptions(
		GetRawHeader(c),
		GetTypeAndScope(c),
		AddAuthorDate(c),
		GetBreakChanges(c, spec),
	)
	if err != nil {
		return result, err
	}
	if len(spec.Types) > 0 {
		for _, typeSpec := range spec.Types {
			if strings.Index(result.RawHeader, typeSpec.Type) == 0 {
				spHead := strings.Split(result.RawHeader, ":")
				if len(spHead) > 1 {
					result.RawHeader = strings.TrimSpace(spHead[1])
				}
			}
		}

	}
	if !c.Hash.IsZero() && gitRepoUrl != "" {
		hashFull := c.Hash.String()
		hashShort := c.Hash.String()[:spec.HashLength]
		result.RawHeader = fmt.Sprintf("%s [%s](%s/commit/%s)", result.RawHeader, hashShort, gitRepoUrl, hashFull)
	}

	return result, nil
}

// NewCommit return conventional commit from git commit
func NewCommit(c git.Commit) (Commit, error) {
	return NewCommitWithOptions(
		GetRawHeader(c),
		GetTypeAndScope(c),
		AddAuthorDate(c),
	)
}

// NewCommitWithOptions return conventional commit with custom option
func NewCommitWithOptions(opts ...OptionFn) (result Commit, err error) {
	for _, opt := range opts {
		if err = opt(&result); err != nil {
			return
		}
	}

	return
}

// AppendMarkdownCommitLink
// will append [shortHash](RaymondRender(commitUrlFormat)) by {{Host}}/{{Owner}}/{{Repository}}/commit/{{Hash}}
func (c *Commit) AppendMarkdownCommitLink(commitUrlFormat string, shortHash, hash string, gitHost, gitOwner, gitRepo string) error {
	commitRt := new(CommitRenderTemplate)
	commitRt.Host = gitHost
	commitRt.Owner = gitOwner
	commitRt.Repository = gitRepo
	commitRt.Hash = hash
	commitUrl, err := RaymondRender(commitUrlFormat, commitRt)
	if err != nil {
		return err
	}
	c.RawHeader = fmt.Sprintf("%s [%s](%s)", c.RawHeader, shortHash, commitUrl)
	return nil
}

func (c *Commit) String() string {
	return c.RawHeader
}
