package convention

import (
	"fmt"
	"github.com/sinlov-go/go-git-tools/git"
	"strings"
)

// https://www.conventionalcommits.org/en/v1.0.0/
// <type>[optional scope]: <description>
// [optional body]
// [optional footer(s)]

// Commit conventional commit
type Commit struct {
	// Commit as is
	RawHeader string

	Type  string
	Scope string
}

// NewCommitWithLogSpec return conventional commit from git commit
func NewCommitWithLogSpec(c git.Commit, spec ConventionalChangeLogSpec) (Commit, error) {
	result, err := NewCommitWithOptions(
		GetRawHeader(c),
		GetTypeAndScope(c),
		AddAuthorDate(c),
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
	return result, err
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

// AppendMarkdownLink
// will append [shortHash](gitHost/gitOwner/gitRepo/commit/hash)
func (c *Commit) AppendMarkdownLink(shortHash, hash string, gitHost, gitOwner, gitRepo string) {
	c.RawHeader = fmt.Sprintf("%s [%s](%s/%s/%s/commit/%s)", c.RawHeader, shortHash, gitHost, gitOwner, gitRepo, hash)
}

func (c *Commit) String() string {
	return c.RawHeader
}
