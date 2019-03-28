package redis

import (
	"github.com/gomodule/redigo/redis"
)

/*
Redis hash 是一个string类型的field和value的映射表，hash特别适合用于存储对象。
Redis 中每个 hash 可以存储 232 - 1 键值对（40多亿）。
*/
const (
	RedisKey_HSET = "HSET"
	RedisKey_HDEL = "HDEL"

	RedisKey_HMSET = "HMSET"

	RedisKey_HSETNX = "HSETNX"

	RedisKey_HExists = "HExists" //命令用于查看哈希表的指定字段是否存在。
)

/*
Redis HSet 命令用于为哈希表中的字段赋值 。
如果哈希表不存在，一个新的哈希表被创建并进行 HSET 操作。
如果字段已经存在于哈希表中，旧值将被覆盖。
*/

func HSet(key string, filed string, value string) (reply int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_HSET, key, filed, value))
}

/*
Redis HSetNx 命令用于为哈希表中不存在的的字段赋值 。
如果哈希表不存在，一个新的哈希表被创建并进行 HSET 操作。
如果字段已经存在于哈希表中，操作无效。
如果 key 不存在，一个新哈希表被创建并执行 HSETNX 命令。
*/
func HSetNx(key, filed, value string) (reply int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_HSETNX, key, filed, value))
}

/*
Redis HSet 命令用于为哈希表中的字段赋值 。
如果哈希表不存在，一个新的哈希表被创建并进行 HSet 操作。
如果字段已经存在于哈希表中，旧值将被覆盖。
*/
func HMSet(key string, fieldsValues ...interface{}) (reply interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return rc.Do(RedisKey_HMSET, argsForm(fieldsValues, key)...)
}

/*
HDel 命令用于删除哈希表 key 中的一个或多个指定字段，不存在的字段将被忽略。
*/
func HDel(key string, fields ...interface{}) (reply int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_HDEL, argsForm(fields, key)...))
}

/*
Hexists 命令用于查看哈希表的指定字段是否存在。
*/
func HExists(key, field string) (reply int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_HExists, key, field))
}
