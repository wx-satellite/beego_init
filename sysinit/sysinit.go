package sysinit

import "byn/utils/redis"

func sysinit() {
	// 初始化 redis 服务
	InitRedisCache()

}


func InitRedisCache() {
	redis.TimeOut = 120
	redis.MaxOpen = 128
	redis.MaxIdle = 128
	redis.Init()
}
