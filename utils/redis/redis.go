package redis

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	pool *redis.Pool
	MaxOpen int
	MaxIdle int
	TimeOut int64
)


func Init() {
	address := beego.AppConfig.String("redis_address")
	password := beego.AppConfig.String("redis_password")
	if "" == address {
		address = "127.0.0.1:6379"
	}
	if MaxIdle <= 0 {
		MaxIdle = 128
	}
	if MaxOpen <= 0 {
		MaxOpen = 128
	}
	if TimeOut <= 0 {
		TimeOut = 120
	}

	pool = &redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxOpen,
		IdleTimeout: time.Duration(TimeOut),
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial(
				"tcp",
				address,
				redis.DialPassword(password),
				redis.DialConnectTimeout(1*time.Second),
				redis.DialReadTimeout(1*time.Second),
				redis.DialWriteTimeout(1*time.Second),
			)
		},
	}
}


func WriteString(key, value string) (err error) {
	_, err = do("SET", key, value)
	return
}

func WriteStringExpire(key, value string, expireTime int64) (err error) {
	_, err = do("SETEX", key, expireTime, value)
	return
}

func ReadString(key string) (value string, err error) {
	value, err = redis.String(do("GET", key))
	return
}

func WriteStruct(key string, obj interface{}) (err error) {
	// 结构体先序列化
	bytes, err := json.Marshal(obj)
	if err != nil {
		return
	}
	err = WriteString(key, string(bytes))
	return
}

// 调用方式：ReadStruct("name", &Dog{}) ，第二个参数是引用
func ReadStruct(key string, obj interface{}) (err error) {
	res, err := ReadString(key)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(res), obj)
	return
}

func do(cmd string, key interface{}, args ...interface{}) (res interface{}, err error) {
	var params []interface{}
	con := pool.Get()
	if err = con.Err(); err != nil {
		return
	}
	params = append(params, key)

	params = append(params, args...)

	return con.Do(cmd, params...)

}