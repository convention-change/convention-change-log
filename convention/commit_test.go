package convention

import (
	"encoding/json"
	"github.com/sebdah/goldie/v2"
	"github.com/sinlov-go/go-git-tools/git"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewCommitWithOutType(t *testing.T) {
	// mock NewCommitWithLogSpec

	tests := []struct {
		name    string
		c       git.Commit
		wantErr error
	}{
		{
			name: "sample",
			c: git.Commit{
				Message: "feat: add commit convention",
				Author: git.Author{
					When: time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local),
				},
			},
		},
		{
			name: "Commit message with scope",
			c: git.Commit{
				Message: "feat(lang): add polish language",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)
			versionRcPath := filepath.Join(filepath.Dir(g.GoldenFileName(t, tc.name)), ".versionrc")
			data, err := os.ReadFile(versionRcPath)
			if err != nil {
				t.Fatal(err)
			}
			var logSpec ConventionalChangeLogSpec
			err = json.Unmarshal(data, &logSpec)
			if err != nil {
				t.Fatal(err)
			}

			// do NewCommitWithLogSpec
			gotResult, gotErr := NewCommitWithLogSpec(tc.c, logSpec)
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
			name: "sample", // TODO: testData/TestAppendMarkdownLink/sample.golden
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
				owner:     "sinlov-go",
				repo:      "go-git-tools",
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
