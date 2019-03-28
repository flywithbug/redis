package redis

import (
	`errors`
	"github.com/gomodule/redigo/redis"
	"time"
)

const (
	RedisKey_SET    = "SET"
	RedisKey_SETEX = "SETEX" //Setex 命令为指定的 key 设置值及其过期时间。如果 key 已经存在， SETEX 命令将会替换旧的值。

	RedisKey_GETSET = "GETSET"

	RedisKey_GET = "GET"

	RedisKey_MSET = "MSET" //批量设置多个string的值
	RedisKey_MGET = "MGET"

	RedisKey_MSETNX = "MSETNX"  //如果所有名称为key i的string都不存在


	RedisKey_SETNX = "SETNX" //（SET if Not eXists） 命令在指定的 key 不存在时，为 key 设置指定的值


	RedisKey_INCR = "INCR"
	RedisKey_INCRBY = "INCRBY"
	RedisKey_DECR = "DECR"
	RedisKey_DECRBY = "DECRBY"

	RedisKey_APPEND = "append"
	RedisKey_SUBSTR = "substr"



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
	if exp > 0 {
		_,err = SetEx(key,value,int(exp))
		return
	}
	return Set(key,value)
}

/*
Setex 命令为指定的 key 设置值及其过期时间。如果 key 已经存在， SETEX 命令将会替换旧的值。
*/
func SetEx(key ,value string,exp int,)(reply string, err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_SETEX, key, exp,value))
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
给名称为key的string赋予上一次的value
*/
func GetSet(key, value string) (reply string, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_GETSET, key, value))
}



/*
Mset 命令用于同时设置一个或多个 key-value 对。
*/
func MSet(keyValues ...interface{})(reply string,err error)  {
	if len(keyValues) %2 != 0 {
		return "",errors.New("key-value  not right")
	}
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_MSET, keyValues...))
}

/*
 Msetnx 命令用于所有给定 key 都不存在时，同时设置一个或多个 key-value 对。
*/
func MSetNx(keys ...interface{})(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_MSETNX, keys...))
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
命令将 key 中储存的数字值增一。
如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作。
如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
本操作的值限制在 64 位(bit)有符号数字表示之内。
*/
func Incr(key string)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_INCR, key))
}

/*
Incrby 命令将 key 中储存的数字加上指定的增量值。
如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCRBY 命令。
如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
本操作的值限制在 64 位(bit)有符号数字表示之内。
*/
func IncrBy(key string,incr int)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_INCRBY, key,incr))
}
/*
Decr 命令将 key 中储存的数字值减一。
如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 DECR 操作。
如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
本操作的值限制在 64 位(bit)有符号数字表示之内。
*/

func Decr(key string)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_DECR, key))
}

/*
DecrBy 命令将 key 所储存的值减去指定的减量值。
如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 DECRBY 操作。
如果值包含错误的类型，或字符串类型的值不能表示为数字，那么返回一个错误。
本操作的值限制在 64 位(bit)有符号数字表示之内。
*/
func DecrBy(key string,decr int)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_DECRBY, key,decr))
}

/*
Append 命令用于为指定的 key 追加值。
如果 key 已经存在并且是一个字符串， APPEND 命令将 value 追加到 key 原来的值的末尾。
如果 key 不存在， APPEND 就简单地将给定 key 设为 value ，就像执行 SET key value 一样。
*/
func Append(key ,value string)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_APPEND, key,value))
}

/*
返回名称为key的string的value的子串
*/
func SubStr(key string,begin,end int)(reply string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_SUBSTR, key,begin,end))
}



