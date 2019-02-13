package logger

import (
	"github.com/lexkong/log"

	"pppobear.cn/jxc-backend/config"
)

func init() {
	env := config.GetEnv()
	passLagerCfg := log.PassLagerCfg{
		Writers:        env.LogWriters,
		LoggerLevel:    env.LogLevel,
		LoggerFile:     env.LoggerFile,
		LogFormatText:  env.LogFormatText,
		RollingPolicy:  env.RollingPolicy,
		LogRotateDate:  env.LogRotateDate,
		LogRotateSize:  env.LogRotateSize,
		LogBackupCount: env.LogBackupCount,
	}

	_ = log.InitWithConfig(&passLagerCfg)
}
