package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	RedisKey_SET    = "SET"
	RedisKey_GETSET = "GETSET"

	RedisKey_GET = "GET"
	RedisKey_DEL = "DEL"
)

/*
给数据库中名称为key的string赋予值value
*/
func Set(key string, value string) (err error) {
	return SetExp(key, value, 0)
}

/*
给数据库中名称为key的string赋予值value 带有效期
*/
func SetExp(key string, value string, exp time.Duration) (err error) {
	rc := redisPool.Get()
	defer rc.Close()
	if exp > 0 {
		_, err = rc.Do(RedisKey_SET, key, value, "EX", fmt.Sprintf("%d", exp))
	} else {
		_, err = rc.Do(RedisKey_SET, key, value)
	}
	return
}

/*
返回数据库中名称为key的string的value
*/
func Get(key string) (reply string, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_GET, key))
}

/*
删除数据，
*/

func Remove(key string) (reply int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_DEL, key))
}

/*
给名称为key的string赋予上一次的value
*/
func GetSet(key, value string) (reply string, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_GETSET, key, value))
}
