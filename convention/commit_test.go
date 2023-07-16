package convention

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/sebdah/goldie/v2"
	"github.com/sinlov-go/go-git-tools/git"
	"github.com/stretchr/testify/assert"
)

func TestNewCommitWithOutType(t *testing.T) {
	// mock NewCommitWithLogSpec

	tests := []struct {
		name       string
		c          git.Commit
		gitRepoUrl string
		wantErr    error
	}{
		{
			name: "sample",
			c: git.Commit{
				Message: "feat: add commit convention",
				Author: git.Author{
					When: time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local),
				},
			},
			gitRepoUrl: "https://github.com/convention-change/convention-change-log",
		},
		{
			name: "Commit message with scope",
			c: git.Commit{
				Message: "feat(lang): add polish language",
			},
			gitRepoUrl: "https://github.com/convention-change/convention-change-log",
		},
		{
			name: "Commit message with hash",
			c: git.Commit{
				Message: "feat: add polish hash",
				Hash:    plumbing.NewHash("a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2"),
			},
			gitRepoUrl: "https://github.com/convention-change/convention-change-log",
		},
		{
			name: "Commit message with hash and breaking change",
			c: git.Commit{
				Message: "feat: new api\n\nBREAKING CHANGE: this is describe of new api breaking changes\n\nfix #1",
				Hash:    plumbing.NewHash("a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2"),
			},
			gitRepoUrl: "https://github.com/convention-change/convention-change-log",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ColoredDiff),
			)
			versionRcPath := filepath.Join(filepath.Dir(g.GoldenFileName(t, tc.name)), ".versionrc")
			data, err := os.ReadFile(versionRcPath)
			if err != nil {
				t.Fatal(err)
			}
			logSpecByData, err := LoadConventionalChangeLogSpecByData(data)
			if err != nil {
				t.Fatal(err)
			}

			// do NewCommitWithLogSpec
			gotResult, gotErr := NewCommitWithLogSpec(tc.c, *logSpecByData, tc.gitRepoUrl)
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			// verify NewCommitWithLogSpec
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}

func TestNewCommit(t *testing.T) {
	tests := []struct {
		name    string
		c       git.Commit
		wantErr error
	}{
		{
			name: "Commit message with not character to draw attention to breaking change",
			c: git.Commit{
				Message: "refactor!: drop support for Node 6",
			},
		},
		{
			name: "Commit message with no body",
			c: git.Commit{
				Message: "docs: correct spelling of CHANGELOG",
			},
		},
		{
			name: "Commit message with scope",
			c: git.Commit{
				Message: "feat(lang): add polish language",
			},
		},
		{
			name: "Uppercase",
			c: git.Commit{
				Message: "REFACTOR!: drop support for Node 6",
			},
		},
		{
			name: "Mixedcase",
			c: git.Commit{
				Message: "Docs: correct spelling of CHANGELOG",
			},
		},
		{
			name: "Misc",
			c: git.Commit{
				Message: "random git message",
			},
		},
		{
			name: "Misc with author date",
			c: git.Commit{
				Message: "random git message",
				Author: git.Author{
					When: time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local),
				},
			},
		},
		{
			name:    "Empty commit",
			wantErr: ErrEmptyCommit,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t)

			gotResult, gotErr := NewCommit(tc.c)
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}

func TestAppendMarkdownLink(t *testing.T) {
	// mock AppendMarkdownLink
	type gitInfo struct {
		shortHash string
		hash      string
		host      string
		owner     string
		repo      string
	}
	tests := []struct {
		name    string
		c       git.Commit
		gitInfo gitInfo
		wantErr error
	}{
		{
			name: "sample", // testdata/TestAppendMarkdownLink/sample.golden
			c: git.Commit{
				Message: "feat: add commit convention",
				Author: git.Author{
					When: time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local),
				},
			},
			gitInfo: gitInfo{
				shortHash: "a1b2c3d",
				hash:      "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0",
				host:      "github.com",
				owner:     "convention-change",
				repo:      "convention-change-log",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			gotResult, gotCommitErr := NewCommit(tc.c)
			if gotCommitErr != nil {
				t.Fatal(gotCommitErr)
			}
			// do AppendMarkdownLink
			gotResult.AppendMarkdownLink(tc.gitInfo.shortHash, tc.gitInfo.hash, tc.gitInfo.host, tc.gitInfo.owner, tc.gitInfo.repo)
			//assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			// verify AppendMarkdownLink
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
