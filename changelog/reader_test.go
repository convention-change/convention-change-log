package changelog

import (
	"fmt"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestNewReader(t *testing.T) {
	// mock NewReader
	tests := []struct {
		name                     string
		changelogMdPath          string
		spec                     convention.ConventionalChangeLogSpec
		wantHistoryTagShort      string
		wantHistoryTag           string
		wantHistoryFirstTitle    string
		wantFirstChangeUrl       string
		wantHistoryFirstNodesLen int
		wantHistoryNodesLen      int
		wantError                bool
	}{
		{
			name:            "not found file",
			changelogMdPath: "not-found",
			spec:            convention.DefaultConventionalChangeLogSpec(),
			wantError:       true,
		},
		{
			name:            "empty",
			changelogMdPath: "empty",
			spec:            convention.DefaultConventionalChangeLogSpec(),
			wantError:       true,
		},
		{
			name:                     "first release", // testdata/TestNewReader/sample.golden
			changelogMdPath:          "first_release",
			spec:                     convention.DefaultConventionalChangeLogSpec(),
			wantHistoryTagShort:      "1.0.0",
			wantHistoryTag:           "v1.0.0",
			wantHistoryFirstTitle:    "1.0.0 (2023-07-11)",
			wantFirstChangeUrl:       "",
			wantHistoryFirstNodesLen: 3,
			wantHistoryNodesLen:      3,
		},
		{
			name:                     "sample", // testdata/TestNewReader/sample.golden
			changelogMdPath:          "sample",
			spec:                     convention.DefaultConventionalChangeLogSpec(),
			wantHistoryTagShort:      "1.1.0",
			wantHistoryTag:           "v1.1.0",
			wantHistoryFirstTitle:    "[1.1.0](https://github.com/sinlov-go/sample-markdown/compare/v1.0.0...v1.1.0) (2023-07-18)",
			wantFirstChangeUrl:       "https://github.com/sinlov-go/sample-markdown/compare/v1.0.0...v1.1.0",
			wantHistoryFirstNodesLen: 3,
			wantHistoryNodesLen:      6,
		},
		{
			name:                     "break change", // testdata/TestNewReader/sample.golden
			changelogMdPath:          "break_change",
			spec:                     convention.DefaultConventionalChangeLogSpec(),
			wantHistoryTagShort:      "1.2.0",
			wantHistoryTag:           "v1.2.0",
			wantHistoryFirstTitle:    "[1.2.0](https://github.com/sinlov-go/sample-markdown/compare/v1.0.0...v1.2.0) (2023-07-18)",
			wantFirstChangeUrl:       "https://github.com/sinlov-go/sample-markdown/compare/v1.0.0...v1.2.0",
			wantHistoryFirstNodesLen: 5,
			wantHistoryNodesLen:      8,
		},
		{
			name:                     "patch version", // testdata/TestNewReader/sample.golden
			changelogMdPath:          "patch_version",
			spec:                     convention.DefaultConventionalChangeLogSpec(),
			wantHistoryTagShort:      "1.2.1",
			wantHistoryTag:           "v1.2.1",
			wantHistoryFirstTitle:    "[1.2.1](https://github.com/sinlov-go/sample-markdown/compare/v1.2.0...v1.2.1) (2023-07-19)",
			wantFirstChangeUrl:       "https://github.com/sinlov-go/sample-markdown/compare/v1.2.0...v1.2.1",
			wantHistoryFirstNodesLen: 2,
			wantHistoryNodesLen:      10,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewReader
			changeLogPath := filepath.Join(filepath.Dir(g.GoldenFileName(t, tc.name)), fmt.Sprintf("%s-log.md.golden", t.Name()))
			reader, gotErr := NewReader(changeLogPath, tc.spec)
			assert.Equal(t, tc.wantError, gotErr != nil)
			if tc.wantError {
				return
			}
			// verify NewReader
			assert.Equal(t, tc.wantHistoryTagShort, reader.HistoryFirstTagShort())
			assert.Equal(t, tc.wantHistoryTag, reader.HistoryFirstTag())
			assert.Equal(t, tc.wantHistoryFirstTitle, reader.HistoryFirstTitle())
			g.Assert(t, t.Name(), []byte(reader.HistoryFirstContent()))
			assert.Equal(t, tc.wantFirstChangeUrl, reader.HistoryFirstChangeUrl())
			assert.Equal(t, tc.wantHistoryFirstNodesLen, len(reader.HistoryFirstNodes()))
			assert.Equal(t, tc.wantHistoryNodesLen, len(reader.HistoryNodes()))
		})
	}
}
