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
		name                  string
		changelogMdPath       string
		spec                  convention.ConventionalChangeLogSpec
		wantHistoryTagShort   string
		wantHistoryFirstTitle string
		wantError             bool
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
			name:                  "first release", // testdata/TestNewReader/sample.golden
			changelogMdPath:       "first_release",
			spec:                  convention.DefaultConventionalChangeLogSpec(),
			wantHistoryTagShort:   "1.0.0",
			wantHistoryFirstTitle: "1.0.0 (2023-07-11)",
		},
		{
			name:                  "sample", // testdata/TestNewReader/sample.golden
			changelogMdPath:       "sample",
			spec:                  convention.DefaultConventionalChangeLogSpec(),
			wantHistoryTagShort:   "1.1.0",
			wantHistoryFirstTitle: "[1.1.0](https://github.com/sinlov-go/sample-markdown/compare/v1.0.0...v1.1.0) (2023-07-18)",
		},
		{
			name:                  "break change", // testdata/TestNewReader/sample.golden
			changelogMdPath:       "break_change",
			spec:                  convention.DefaultConventionalChangeLogSpec(),
			wantHistoryTagShort:   "1.1.0",
			wantHistoryFirstTitle: "[1.1.0](https://github.com/sinlov-go/sample-markdown/compare/v1.0.0...v1.1.0) (2023-07-18)",
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
			assert.Equal(t, tc.wantHistoryFirstTitle, reader.HistoryFirstTitle())
			g.Assert(t, t.Name(), []byte(reader.HistoryFirstContent()))
		})
	}
}
