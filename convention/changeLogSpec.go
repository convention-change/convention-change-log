package convention

import (
	"encoding/json"
	"fmt"
)

// ConventionalChangeLogSpec
// struct
type ConventionalChangeLogSpec struct {

	// Types
	//
	Types []Types `json:"types,omitempty"`
}

var (
	defaultConventionalChangeLogSpec *ConventionalChangeLogSpec
)

// DefaultConventionalChangeLogSpec
//
// See: https://www.conventionalcommits.org
func DefaultConventionalChangeLogSpec() ConventionalChangeLogSpec {
	if defaultConventionalChangeLogSpec == nil {
		defaultConventionalChangeLogSpec = &ConventionalChangeLogSpec{
			Types: defaultType,
		}
	}

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

	return &spec, nil
}
