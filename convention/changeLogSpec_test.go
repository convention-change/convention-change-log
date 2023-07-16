package convention

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConventionalChangeLogSpecByData(t *testing.T) {
	// mock LoadConventionalChangeLogSpecByData
	tests := []struct {
		name    string
		c       []byte
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestLoadConventionalChangeLogSpecByData/sample.golden
			c: []byte(`
{
  "types": [
    {"type": "feat", "section": "✨ Features", "hidden": false},
    {"type": "fix", "section": "🐛 Bug Fixes", "hidden": false}
  ],
  "tag-prefix": "",
  "skip": {
    "bump": false,
    "changelog": false,
    "commit": false,
    "tag": false
  }
}`),
		},
		{
			name: "full",
			c: []byte(`
{
  "types": [
    {"type": "feat", "section": "✨ Features", "hidden": false},
    {"type": "fix", "section": "🐛 Bug Fixes", "hidden": false},
    {"type": "docs", "section":"📝 Documentation", "hidden": true},
    {"type": "style", "section":"💄 Styles", "hidden": true},
    {"type": "refactor", "section":"♻ Refactor", "hidden": false},
    {"type": "perf", "section":"⚡ Performance Improvements", "hidden": false},
    {"type": "test", "section":"✅ Tests", "hidden": true},
    {"type": "build", "section":"👷‍ Build System", "hidden": false},
    {"type": "ci", "section":"🔧 Continuous Integration", "hidden": true},
    {"type": "chore", "section":"📦 Chores", "hidden": true},
    {"type": "revert", "section":"⏪ Reverts", "hidden": false}
  ],
  "tag-prefix": "",
  "skip": {
    "bump": false,
    "changelog": false,
    "commit": false,
    "tag": false
  }
}`),
		},
		{
			name: "Error", // testdata/TestLoadConventionalChangeLogSpecByData/sample.golden
			c: []byte(`
{
  "types": [
    {"type": "feat", "section": "✨ Features", "hidden": false},
    {"type": "fix", "section": "🐛 Bug Fixes", "hidden": false},
  ],
  "tag-prefix": "",
  "skip": {
    "bump": false,
    "changelog": false,
    "commit": false,
    "tag": false
  }
}`),
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.Simple),
			)

			// do LoadConventionalChangeLogSpecByData
			gotResult, gotErr := LoadConventionalChangeLogSpecByData(tc.c)
			if gotErr != nil {
				if !tc.wantErr {
					t.Fatalf("LoadConventionalChangeLogSpecByData got unexpected error: %v", gotErr)
				}
			}
			if tc.wantErr {
				if gotErr == nil {
					t.Fatal("LoadConventionalChangeLogSpecByData should return error")
				}
				return
			}
			// verify LoadConventionalChangeLogSpecByData
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}

func TestParseSectionFromType(t *testing.T) {
	// mock ParseSectionFromType
	tests := []struct {
		name       string
		wantRes    string
		logSpec    ConventionalChangeLogSpec
		commitType string
	}{
		{
			name:       "Features",
			wantRes:    "Features",
			logSpec:    DefaultConventionalChangeLogSpec(),
			commitType: FeatType,
		},
		{
			name:       "Documentation",
			wantRes:    "Documentation",
			logSpec:    DefaultConventionalChangeLogSpec(),
			commitType: DocsType,
		},
		{
			name:       "Custom",
			wantRes:    "Custom",
			logSpec:    DefaultConventionalChangeLogSpec(),
			commitType: "Custom",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do ParseSectionFromType
			gotResult := ParseSectionFromType(tc.logSpec, tc.commitType)

			// verify ParseSectionFromType

			assert.Equal(t, tc.wantRes, gotResult)
		})
	}
}

func TestSimplifyConventionalChangeLogSpec(t *testing.T) {
	// mock SimplifyConventionalChangeLogSpec
	tests := []struct {
		name    string
		wantErr error
	}{
		{
			name: "simplify", // testdata/TestSimplifyConventionalChangeLogSpec/simplify.golden
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ColoredDiff),
			)

			// do SimplifyConventionalChangeLogSpec
			gotResult := SimplifyConventionalChangeLogSpec()

			// verify SimplifyConventionalChangeLogSpec
			g.AssertJson(t, t.Name(), &gotResult)
		})
	}
}
