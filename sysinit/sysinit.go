package sysinit

import (
	"byn/utils/redis"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

func sysinit() {
	// 初始化日志
	InitLogHandle()
	// 初始化 redis 服务
	InitRedisCache()

}

func InitLogHandle() {
	filename := "byn.log"
	level := "debug"
	levelNumber := 7
	if name := beego.AppConfig.String("log::filename"); name != "" {
		filename = name
	}
	if l := beego.AppConfig.String("log::level"); l != "" {
		level = strings.ToLower(l)
	}
	switch level {
	case "info":
		levelNumber = 6
	case "notice":
		levelNumber = 5
	case "warning":
		levelNumber = 4
	case "error":
		levelNumber = 3
	}
	params := map[string]interface{}{
		"filename": filename,
		"level":    levelNumber,
	}
	configString, _ := json.Marshal(params)
	_ = logs.SetLogger(logs.AdapterFile, string(configString))
	//日志默认不输出调用的文件名的和行号
	logs.EnableFuncCallDepth(true)

}

func InitRedisCache() {
	redis.TimeOut = 120
	redis.MaxOpen = 128
	redis.MaxIdle = 128
	redis.Init()
}
