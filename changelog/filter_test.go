package changelog

import (
	"github.com/convention-change/convention-change-log/convention"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name    string
		commits []convention.Commit
		logSpec convention.ConventionalChangeLogSpec
		want    map[string][]convention.Commit
	}{
		{
			name: "with DefaultConventionalChangeLogSpec",
			commits: []convention.Commit{
				{
					RawHeader: "feature A",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "fix B",
					Type:      convention.FixType,
				},
				{
					RawHeader: "build",
					Type:      convention.BuildType,
				},
			},
			logSpec: convention.DefaultConventionalChangeLogSpec(),
			want: map[string][]convention.Commit{
				convention.FeatType: {
					{
						RawHeader: "feature A",
						Type:      convention.FeatType,
					},
				},
				convention.FixType: {
					{
						RawHeader: "fix B",
						Type:      convention.FixType,
					},
				},
				convention.BuildType: {
					{
						RawHeader: "build",
						Type:      convention.BuildType,
					},
				},
			},
		},
		{
			name: "with DefaultConventionalChangeLogSpec more feature",
			commits: []convention.Commit{
				{
					RawHeader: "feature A",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "feature B",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "fix B",
					Type:      convention.FixType,
				},
				{
					RawHeader: "ci biz",
					Type:      convention.CIType,
				},
			},
			logSpec: convention.DefaultConventionalChangeLogSpec(),
			want: map[string][]convention.Commit{
				convention.FeatType: {
					{
						RawHeader: "feature A",
						Type:      convention.FeatType,
					},
					{
						RawHeader: "feature B",
						Type:      convention.FeatType,
					},
				},
				convention.FixType: {
					{
						RawHeader: "fix B",
						Type:      convention.FixType,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := filter(tc.commits, tc.logSpec)
			assert.Equal(t, tc.want, got)
		})
	}
}
