package changelog

import (
	"testing"
	"time"

	"github.com/convention-change/convention-change-log/convention"
	"github.com/sebdah/goldie/v2"
	"github.com/sinlov-go/sample-markdown/sample_mk"
	"github.com/stretchr/testify/assert"
)

func TestGenerateMarkdownNodes(t *testing.T) {

	gitHttpInfoDefault := convention.GitRepositoryHttpInfo{
		Scheme:     "https",
		Host:       "github.com",
		Owner:      "convention-change",
		Repository: "convention-change-log",
	}

	// mock GenerateMarkdownNodes
	tests := []struct {
		name              string
		gitRepositoryInfo convention.GitRepositoryHttpInfo
		changelogDesc     ConventionalChangeLogDesc
		commits           []convention.Commit
		logSpec           convention.ConventionalChangeLogSpec
		wantErr           error
	}{
		{
			name:              "empty old data",
			gitRepositoryInfo: gitHttpInfoDefault,
			changelogDesc: ConventionalChangeLogDesc{
				Version:      "v1.0.0",
				When:         time.Date(2020, 1, 18, 0, 1, 23, 45, time.UTC),
				Location:     time.UTC,
				ToolsKitName: "convention-change-log",
				ToolsKitURL:  "https://github.com/convention-change/convention-change-log",
			},
			commits: []convention.Commit{
				{
					RawHeader: "feat: new feature",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "fix: new fix",
					Type:      convention.FixType,
				},
				{
					RawHeader: "commit sample 1",
				},
				{
					RawHeader: "commit sample 2\nmore info",
				},
			},
			logSpec: convention.DefaultConventionalChangeLogSpec(),
		},
		{
			name:              "version notes url",
			gitRepositoryInfo: gitHttpInfoDefault,
			changelogDesc: ConventionalChangeLogDesc{
				Version:      "v1.1.0",
				PreviousTag:  "v1.0.0",
				When:         time.Date(2020, 1, 18, 0, 1, 23, 45, time.UTC),
				Location:     time.UTC,
				ToolsKitName: "convention-change-log",
				ToolsKitURL:  "https://github.com/convention-change/convention-change-log",
			},
			commits: []convention.Commit{
				{
					RawHeader: "feat: new feature",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "fix: new fix",
					Type:      convention.FixType,
				},
				{
					RawHeader: "chore: new build",
					Type:      convention.ChoreType,
				},
			},
			logSpec: convention.DefaultConventionalChangeLogSpec(),
		},
		{
			name:              "many commits",
			gitRepositoryInfo: gitHttpInfoDefault,
			changelogDesc: ConventionalChangeLogDesc{
				Version:      "v1.0.0",
				When:         time.Date(2020, 1, 18, 0, 0, 0, 0, time.UTC),
				Location:     time.UTC,
				ToolsKitName: "convention-change-log",
				ToolsKitURL:  "https://github.com/convention-change/convention-change-log",
			},
			commits: []convention.Commit{
				{
					RawHeader: "feat: new feature",
					Type:      convention.FeatType,
					IssueInfo: convention.IssueInfo{
						IssueReferencesId: 1,
						IssuePrefix:       "#",
						IssueReference:    "fix",
					},
				},
				{
					RawHeader: "feat: support new client",
					Type:      convention.FeatType,
					BreakingChanges: convention.BreakingChanges{
						Describe: "desc breaking changes support new client",
					},
				},
				{
					RawHeader: "fix: new fix",
					Type:      convention.FixType,
					BreakingChanges: convention.BreakingChanges{
						Describe: "desc breaking changes new fix",
					},
				},
				{
					RawHeader: "fix: wrong color",
					Type:      convention.FixType,
					IssueInfo: convention.IssueInfo{
						IssueReferencesId: 2,
						IssuePrefix:       "#",
						IssueReference:    "fix",
					},
				},
				{
					RawHeader: "chore: new build",
					Type:      convention.ChoreType,
				},
				{
					RawHeader: "chore(github): release on github",
					Type:      convention.ChoreType,
				},
				{
					RawHeader: "unleash the dragon",
					Type:      convention.MiscType,
				},
				{
					RawHeader: "chore(gitlab): release on gitlab",
					Type:      convention.ChoreType,
				},
			},
			logSpec: convention.DefaultConventionalChangeLogSpec(),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do GenerateMarkdownNodes
			gotResult, gotErr := GenerateMarkdownNodes(
				tc.gitRepositoryInfo,
				tc.changelogDesc,
				tc.commits,
				tc.logSpec,
			)
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			// verify GenerateMarkdownNodes
			generateText := sample_mk.GenerateText(gotResult)
			g.Assert(t, t.Name(), []byte(generateText))
		})
	}
}
