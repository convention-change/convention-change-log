package log

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/sinlov/gh-conventional-kit/constant"
	"github.com/sinlov/gh-conventional-kit/utils/env_kit"
)

// InitLog
//
//	Initialization github.com/bar-counter/slog
func InitLog(isShowDebug, isHideLineno bool) error {

	envLogLevel := env_kit.FetchOsEnvStr(constant.EnvLogLevel, "")
	if envLogLevel == "" {
		if isShowDebug {
			envLogLevel = slog.DEBUG
		} else {
			envLogLevel = slog.INFO
		}
	} else {
		fmt.Printf("-> app clog level InitLog by env: %s=%s\n", constant.EnvLogLevel, envLogLevel)
	}

	passLagerCfg := slog.PassLagerCfg{
		LoggerLevel:    envLogLevel,
		Writers:        "stdout",
		LoggerFile:     "",
		LogFormatText:  true,
		LogHideLineno:  isHideLineno,
		RollingPolicy:  "",
		LogRotateDate:  1,
		LogRotateSize:  8,
		LogBackupCount: 7,
	}
	err := slog.InitWithConfig(&passLagerCfg)
	return err
}
