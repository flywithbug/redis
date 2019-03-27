package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	RedisKey_SET    = "SET"
	RedisKey_GETSET = "GETSET"

	RedisKey_GET = "GET"
	RedisKey_DEL = "DEL"

	RedisKey_MGET = "MGET"

	RedisKey_SETNX = "SETNX" //（SET if Not eXists） 命令在指定的 key 不存在时，为 key 设置指定的值

	RedisKey_SETEX = "SETEX" //Setex 命令为指定的 key 设置值及其过期时间。如果 key 已经存在， SETEX 命令将会替换旧的值。

)

/*
给数据库中名称为key的string赋予值value
*/
func Set(key string, value string) (err error) {
	rc := redisPool.Get()
	defer rc.Close()
	_, err = rc.Do(RedisKey_SET, key, value)
	return
}

/*
给数据库中名称为key的string赋予值value 带有效期
*/
func SetExp(key string, value string, exp time.Duration) (err error) {
	rc := redisPool.Get()
	defer rc.Close()
	if exp > 0 {
		_,err = SetEx(key,value,int(exp))
		return
	}
	return Set(key,value)
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

/*
返回库中多个string的value
*/
func MGet(keys ...interface{})(reply []string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Strings(rc.Do(RedisKey_MGET, keys...))
}


/*
Setnx（SET if Not eXists） 命令在指定的 key 不存在时，为 key 设置指定的值。
*/
func SetNx(key, value string) (reply int, err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_SETNX, key, value))
}

/*
Setex 命令为指定的 key 设置值及其过期时间。如果 key 已经存在， SETEX 命令将会替换旧的值。
*/
func SetEx(key ,value string,exp int,)(reply string, err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_SETEX, key, exp,value))
}
