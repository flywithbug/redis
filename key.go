package redis

import (
	`github.com/gomodule/redigo/redis`
	`time`
)

const (
	RedisKey_DEL = "DEL"
	RedisKey_DUMP = "DUMP"

	RedisKey_EXISTS = "EXISTS"

	RedisKey_Expire = "Expire"
	RedisKey_Expireat = "Expireat"
)



/*
删除数据，
*/
func Del(key string) (reply int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_DEL, key))
}

/*
 DUMP 命令用于序列化给定 key ，并返回被序列化的值。
*/
func Dump(key string)(reply interface{},err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return rc.Do(RedisKey_DUMP, key)
}


/*
 EXISTS 命令用于检查给定 key 是否存在。
*/
func Exists(key string)(reply bool,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Bool(rc.Do(RedisKey_EXISTS, key))
}

/*
Expire 命令用于设置 key 的过期时间，key 过期后将不再可用。单位以秒计。
*/
func Expire(key string,exp int)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_Expire, key,exp))
}

/*
Expireat 命令用于以 UNIX 时间戳(unix timestamp)格式设置 key 的过期时间。key 过期后将不再可用。 timestamp 是到期时间
*/
// timestamp 是到期的时间 time.Now().Add(time.Second*10))
func ExpireAt(key string,timestamp time.Time)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_Expireat, key,timestamp.Unix()))
}


