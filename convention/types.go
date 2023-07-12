package convention

const (
	FeatType     = "feat"
	FixType      = "fix"
	DocsType     = "docs"
	StyleType    = "style"
	RefactorType = "refactor"
	PerfType     = "perf"
	TestType     = "test"
	BuildType    = "build"
	CIType       = "ci"
	ChoreType    = "chore"
	RevertType   = "revert"
	MiscType     = "misc"
)

var defaultConventionalChangeLogSpec *ConventionalChangeLogSpec

// DefaultConventionalChangeLogSpec
//
// See: https://www.conventionalcommits.org
func DefaultConventionalChangeLogSpec() ConventionalChangeLogSpec {
	if defaultConventionalChangeLogSpec == nil {
		defaultConventionalChangeLogSpec = &ConventionalChangeLogSpec{
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
				{
					Type:    BuildType,
					Section: "Build System",
					Hidden:  false,
					Sort:    8,
				},
				{
					Type:    CIType,
					Section: "Continuous Integration",
					Hidden:  true,
					Sort:    9,
				},
				{
					Type:    ChoreType,
					Section: "Chores",
					Hidden:  true,
					Sort:    10,
				},
				{
					Type:    RevertType,
					Section: "Reverts",
					Hidden:  false,
					Sort:    11,
				},
				{
					Type:    MiscType,
					Section: "Miscellaneous",
					Hidden:  false,
					Sort:    12,
				},
			},
		}
	}

	return *defaultConventionalChangeLogSpec
}

// Types
//
//	struct of type
//
// See: https://github.com/conventional-changelog/conventional-changelog-config-spec/blob/master/versions/2.2.0/schema.json
//
//	add Sort field
type Types struct {

	// Type
	//
	Type string `json:"type,omitempty"`

	// Section
	//
	Section string `json:"section,omitempty"`

	// Hidden
	//
	Hidden bool `json:"hidden,omitempty"`

	// Sort
	Sort int `json:"sort,omitempty"`
}

// ConventionalChangeLogSpec
// struct
type ConventionalChangeLogSpec struct {

	// Types
	//
	Types []Types `json:"types,omitempty"`
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
