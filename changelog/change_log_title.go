package changelog

import "time"

type ConventionalChangeLogDesc struct {
	Version  string
	When     time.Time
	Location *time.Location

	ToolsKitName string
	ToolsKitURL  string
}
