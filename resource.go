package convention_change_log

import (
	_ "embed"
)

//go:embed package.json
var PackageJson string

//go:embed resource/versionrc-beauty.json
var ResVersionRcBeautyJson string
