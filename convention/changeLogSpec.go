package convention

import (
	"encoding/json"
	"fmt"
)

const (
	DefaultHeader                     = "Changelog"
	DefaultCommitUrlFormat            = "{{scheme}}://{{host}}/{{owner}}/{{repository}}/commit/{{hash}}"
	DefaultCompareUrlFormat           = "{{scheme}}://{{host}}/{{owner}}/{{repository}}/compare/{{previousTag}}...{{currentTag}}"
	DefaultIssueUrlFormat             = "{{scheme}}://{{host}}/{{owner}}/{{repository}}/issues/{{id}}"
	DefaultUserUrlFormat              = "{{scheme}://{{host}}/{{user}}"
	DefaultReleaseCommitMessageFormat = "chore(release): {{currentTag}}"
)

// ConventionalChangeLogSpec
// struct
// scheme See: https://github.com/conventional-changelog/conventional-changelog-config-spec/blob/master/versions/2.2.0/schema.json
type ConventionalChangeLogSpec struct {

	// Types
	//
	Types []Types `json:"types,omitempty"`

	// TagPrefix
	//	default is v
	TagPrefix string `json:"tag-prefix,omitempty"`

	// HashLength
	//	default is 8
	HashLength uint `json:"hash-length,omitempty"`

	// IssuePrefixes
	// default is ["#"]
	IssuePrefixes []string `json:"issuePrefixes,omitempty"`

	// Header
	//	A string to be used as the main header of the CHANGELOG
	// default DefaultHeader
	Header string `json:"header,omitempty"`

	// CommitUrlFormat
	//	A URL representing a specific commit at a Hash
	// default DefaultCommitUrlFormat
	CommitUrlFormat string `json:"commitUrlFormat,omitempty"`

	// CompareUrlFormat
	//	A URL representing the comparison between two git shas
	// default DefaultCompareUrlFormat
	CompareUrlFormat string `json:"compareUrlFormat,omitempty"`

	// IssueUrlFormat
	//	A URL representing the issue format (allowing a different URL format to be swapped in for Gitlab, Bitbucket, etc)
	// default DefaultIssueUrlFormat
	IssueUrlFormat string `json:"issueUrlFormat,omitempty"`

	// UserUrlFormat
	//	A URL representing a user's profile URL on GitHub, Gitlab, etc. This URL is used for substituting @bcoe with https://github.com/bcoe in commit messages.
	// default DefaultUserUrlFormat
	UserUrlFormat string `json:"userUrlFormat,omitempty"`

	// ReleaseCommitMessageFormat
	//	A string to be used to format the auto-generated release commit message
	// default DefaultReleaseCommitMessageFormat
	ReleaseCommitMessageFormat string `json:"releaseCommitMessageFormat,omitempty"`
}

var (
	defaultConventionalChangeLogSpec *ConventionalChangeLogSpec
)

// SimplifyConventionalChangeLogSpec
// return simplify ConventionalChangeLogSpec
func SimplifyConventionalChangeLogSpec() *ConventionalChangeLogSpec {
	spec := &ConventionalChangeLogSpec{
		Types: []Types{
			{
				Type:    FeatType,
				Section: "Features",
				Hidden:  false,
				Sort:    1,
			},
			{
				Type:    FixType,
				Section: "Bug Fixes",
				Hidden:  false,
				Sort:    2,
			},
			{
				Type:    DocsType,
				Section: "Documentation",
				Hidden:  true,
				Sort:    3,
			},
			{
				Type:    StyleType,
				Section: "Styles",
				Hidden:  true,
				Sort:    4,
			},
			{
				Type:    RefactorType,
				Section: "Refactor",
				Hidden:  false,
				Sort:    5,
			},
			{
				Type:    PerfType,
				Section: "Performance Improvements",
				Hidden:  false,
				Sort:    6,
			},
			{
				Type:    TestType,
				Section: "Tests",
				Hidden:  true,
				Sort:    7,
			},
		},
	}
	return spec
}

// DefaultConventionalChangeLogSpec
//
// See: https://www.conventionalcommits.org
func DefaultConventionalChangeLogSpec() ConventionalChangeLogSpec {
	if defaultConventionalChangeLogSpec == nil {
		defaultConventionalChangeLogSpec = &ConventionalChangeLogSpec{
			Types: defaultType,
		}
	}
	if defaultConventionalChangeLogSpec.TagPrefix == "" {
		defaultConventionalChangeLogSpec.TagPrefix = "v"
	}
	defaultConventionalChangeLogSpec.HashLength = 8
	defaultConventionalChangeLogSpec.IssuePrefixes = []string{"#"}
	defaultConventionalChangeLogSpec.Header = DefaultHeader
	defaultConventionalChangeLogSpec.CommitUrlFormat = DefaultCommitUrlFormat
	defaultConventionalChangeLogSpec.CompareUrlFormat = DefaultCompareUrlFormat
	defaultConventionalChangeLogSpec.IssueUrlFormat = DefaultIssueUrlFormat
	defaultConventionalChangeLogSpec.UserUrlFormat = DefaultUserUrlFormat
	defaultConventionalChangeLogSpec.ReleaseCommitMessageFormat = DefaultReleaseCommitMessageFormat

	return *defaultConventionalChangeLogSpec
}

// ParseSectionFromType
//
//	parse section from type
//	if not found, return type itself
func ParseSectionFromType(logSpec ConventionalChangeLogSpec, commitType string) string {
	for _, t := range logSpec.Types {
		if t.Type == commitType {
			return t.Section
		}
	}

	return commitType
}

// LoadConventionalChangeLogSpecByData
//
//	this function will load ConventionalChangeLogSpec from json data
//	if type sort is 0, will set default sort by convention.defaultType
//
// scheme See: https://github.com/conventional-changelog/conventional-changelog-config-spec/blob/master/versions/2.2.0/schema.json
func LoadConventionalChangeLogSpecByData(logSpec []byte) (*ConventionalChangeLogSpec, error) {
	var spec ConventionalChangeLogSpec
	err := json.Unmarshal(logSpec, &spec)
	if err != nil {
		return nil, fmt.Errorf("load ConventionalChangeLogSpec by data error: %s", err.Error())
	}

	var newType []Types

	// check sort
	for _, t := range spec.Types {
		if t.Sort == 0 {
			for _, dt := range defaultType {
				if dt.Type == t.Type {
					t.Sort = dt.Sort
				}
			}
		}
		newType = append(newType, t)
	}
	spec.Types = newType

	if spec.TagPrefix == "" {
		spec.TagPrefix = "v"
	}

	if spec.HashLength == 0 {
		spec.HashLength = 8
	}

	if spec.IssuePrefixes == nil || len(spec.IssuePrefixes) == 0 {
		spec.IssuePrefixes = []string{"#"}
	}

	if spec.Header == "" {
		spec.Header = DefaultHeader
	}

	if spec.CommitUrlFormat == "" {
		spec.CommitUrlFormat = DefaultCommitUrlFormat
	}

	if spec.CompareUrlFormat == "" {
		spec.CompareUrlFormat = DefaultCompareUrlFormat
	}

	if spec.IssueUrlFormat == "" {
		spec.IssueUrlFormat = DefaultIssueUrlFormat
	}

	if spec.UserUrlFormat == "" {
		spec.UserUrlFormat = DefaultUserUrlFormat
	}

	if spec.ReleaseCommitMessageFormat == "" {
		spec.ReleaseCommitMessageFormat = DefaultReleaseCommitMessageFormat
	}

	return &spec, nil
}
