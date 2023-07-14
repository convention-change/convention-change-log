package changelog

import (
	"testing"
	"time"

	"github.com/sebdah/goldie/v2"
	"github.com/sinlov-go/convention-change-log/convention"
	"github.com/sinlov-go/sample-markdown/sample_mk"
	"github.com/stretchr/testify/assert"
)

func TestGenerateMarkdownNodes(t *testing.T) {
	// mock GenerateMarkdownNodes
	tests := []struct {
		name          string
		commits       []convention.Commit
		logSpec       convention.ConventionalChangeLogSpec
		changelogDesc ConventionalChangeLogDesc
		wantErr       error
	}{
		{
			name: "empty old data",
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
			changelogDesc: ConventionalChangeLogDesc{
				Version:      "v1.0.0",
				When:         time.Date(2020, 1, 18, 0, 1, 23, 45, time.UTC),
				Location:     time.UTC,
				ToolsKitName: "convention-change-log",
				ToolsKitURL:  "https://github.com/sinlov-go/convention-change-log",
			},
		},
		{
			name: "version notes url",
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
			changelogDesc: ConventionalChangeLogDesc{
				Version:         "v1.0.0",
				VersionNotesUrl: "https://github.com/sinlov-go/convention-change-log/compare/v1.0.1...v1.1.0",
				When:            time.Date(2020, 1, 18, 0, 1, 23, 45, time.UTC),
				Location:        time.UTC,
				ToolsKitName:    "convention-change-log",
				ToolsKitURL:     "https://github.com/sinlov-go/convention-change-log",
			},
		},
		{
			name: "many commits",
			commits: []convention.Commit{
				{
					RawHeader: "feat: new feature",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "feat: support new client",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "fix: new fix",
					Type:      convention.FixType,
				},
				{
					RawHeader: "fix: wrong color",
					Type:      convention.FixType,
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
					RawHeader: "chore(gitlab): release on gitlab",
					Type:      convention.ChoreType,
				},
				{
					RawHeader: "unleash the dragon",
					Type:      convention.MiscType,
				},
			},
			logSpec: convention.DefaultConventionalChangeLogSpec(),
			changelogDesc: ConventionalChangeLogDesc{
				Version:      "v1.0.0",
				When:         time.Date(2020, 1, 18, 0, 0, 0, 0, time.UTC),
				Location:     time.UTC,
				ToolsKitName: "convention-change-log",
				ToolsKitURL:  "https://github.com/sinlov-go/convention-change-log",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do GenerateMarkdownNodes
			gotResult, gotErr := GenerateMarkdownNodes(tc.commits,
				tc.logSpec,
				tc.changelogDesc,
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
