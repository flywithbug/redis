package redis

import (
	`errors`
	`github.com/gomodule/redigo/redis`
)

const (
	RedisKey_PFADD = "PFADD"
	RedisKey_PFCOUNT = "PFCOUNT"
	RedisKey_PFMERGE = "PFMERGE"
)

var (
	elementErr = errors.New("element is nil")
)

/*
Redis PFAdd 命令将所有元素参数添加到 HyperLogLog 数据结构中。
*/
func PFAdd(key interface{},element ...interface{})(reply int,err error)  {
	if len(element) == 0 {
		return 0,elementErr
	}
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_PFADD, argsForm(element,key)...))
}

/*
Redis PFCount 命令返回给定 HyperLogLog 的基数估算值。
*/
func PFCount(key interface{})(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_PFCOUNT, key))
}

/*
Redis PFMerge 命令将多个 HyperLogLog 合并为一个 HyperLogLog ，
合并后的 HyperLogLog 的基数估算值是通过对所有 给定 HyperLogLog 进行并集计算得出的。
*/
func PFMerge(keys ...interface{})(reply string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_PFMERGE, keys...))
}
