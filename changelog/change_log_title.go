package changelog

import "time"

type ConventionalChangeLogDesc struct {
	Version string

	PreviousTag string

	When     time.Time
	Location *time.Location

	ToolsKitName string
	ToolsKitURL  string
}
