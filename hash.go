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
	RedisKey_HGET = "HGET"
	RedisKey_HGETALL = "HGETALL"

	RedisKey_HINCRBY  = "HINCRBY"

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

func HSet(key , field , value interface{}) (reply int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_HSET, key, field, value))
}


/*
 HGet 命令用于返回哈希表中指定字段的值。
*/
func HGet(key,filed interface{} )(reply string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_HGET, key,filed))
}
/*
 HGetAll 命令用于返回哈希表中，所有的字段和值。
 在返回值里，紧跟每个字段名(field name)之后是字段的值(value)，所以返回值的长度是哈希表大小的两倍。
*/
func HGetAll(key interface{})(reply []string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Strings(rc.Do(RedisKey_HGETALL, key))
}

/*
HIncrBy 命令用于为哈希表中的字段值加上指定增量值。
增量也可以为负数，相当于对指定字段进行减法操作。
如果哈希表的 key 不存在，一个新的哈希表被创建并执行 HIncrBy 命令。
如果指定的字段不存在，那么在执行命令前，字段的值被初始化为 0 。
对一个储存字符串值的字段执行 HIncrBy 命令将造成一个错误。
本操作的值被限制在 64 位(bit)有符号数字表示之内。
*/
func HIncrBy(key ,field interface{},incr int) (reply int, err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_HINCRBY, key, field, incr))
}

/*
Redis HSetNx 命令用于为哈希表中不存在的的字段赋值 。
如果哈希表不存在，一个新的哈希表被创建并进行 HSET 操作。
如果字段已经存在于哈希表中，操作无效。
如果 key 不存在，一个新哈希表被创建并执行 HSETNX 命令。
*/
func HSetNx(key, field, value interface{}) (reply int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_HSETNX, key, field, value))
}

/*
Redis HSet 命令用于为哈希表中的字段赋值 。
如果哈希表不存在，一个新的哈希表被创建并进行 HSet 操作。
如果字段已经存在于哈希表中，旧值将被覆盖。
*/
func HMSet(key interface{}, fieldsValues ...interface{}) (reply interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return rc.Do(RedisKey_HMSET, argsForm(fieldsValues, key)...)
}

/*
HDel 命令用于删除哈希表 key 中的一个或多个指定字段，不存在的字段将被忽略。
*/
func HDel(key interface{}, fields ...interface{}) (reply int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_HDEL, argsForm(fields, key)...))
}

/*
Hexists 命令用于查看哈希表的指定字段是否存在。
*/
func HExists(key, field interface{}) (reply int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_HExists, key, field))
}
