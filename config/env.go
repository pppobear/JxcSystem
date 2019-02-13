package config

import "github.com/gin-gonic/gin"

// 环境配置文件
// 可配置多个环境配置，进行切换

type Env struct {
	Debug            bool
	DatabaseIP       string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	ServerPort       string
	RedisIP          string
	RedisPort        string
	RedisPassword    string
	RedisDb          int
	RedisJwtDb       int
	RedisCacheDb     int
	AppSecret        string

	// 日志配置
	LogLevel       string
	SQLLog         bool
	LogWriters     string
	LoggerFile     string
	LogFormatText  bool
	RollingPolicy  string
	LogRotateDate  int
	LogRotateSize  int
	LogBackupCount int
}

var devEnv = Env{
	Debug: true,

	ServerPort:       "4000",
	DatabaseIP:       "127.0.0.1",
	DatabasePort:     "1433",
	DatabaseUsername: "SA",
	DatabasePassword: "Password*",
	DatabaseName:     "Jxc",

	RedisIP:       "127.0.0.1",
	RedisPort:     "6379",
	RedisPassword: "",
	RedisDb:       0,

	RedisJwtDb:   1,
	RedisCacheDb: 2,

	LogLevel:       "Debug",
	LogWriters:     "file,stdout",
	LoggerFile:     "storage/logs/api.log",
	LogFormatText:  false,
	RollingPolicy:  "size",
	LogRotateDate:  1,
	LogRotateSize:  1024,
	LogBackupCount: 7,

	AppSecret: "something-very-secret",
}

func GetEnv() *Env {
	return &devEnv
}

func Setup() {
	if devEnv.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
