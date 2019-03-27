package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

const (
	RedisKey_SADD = "SADD"
	RedisKey_SREM = "SREM"
	RedisKey_SPOP = "SPOP"

	RedisKey_SMOVE = "SMOVE"

	RedisKey_SCARD = "SCARD"

	RedisKey_SEXIST = "SISMEMBER"

	RedisKey_SINTER      = "SINTER"      //交集
	RedisKey_SINTERSTORE = "SINTERSTORE" //求交集并将交集保存到dstKey集合中

	RedisKey_SUNION      = "SUNION"      //求并集
	RedisKey_SUNIONSTORE = "SUNIONSTORE" //求并集并将并集保存到dstkey的集合

	RedisKey_SDIFF      = "SDIFF"      //差集
	RedisKey_SDIFFSTORE = "SDIFFSTORE" //求交集并将交集保存到dstkey的集合

	RedisKey_SMEMBERS = "SMEMBERS" //返回名称为key的set的所有元素

	RedisKey_SRANDMEMBER = "SRANDMEMBER" //随机返回名称为key的set的一个元素

)

/*
向名称为key的set中添加元素member
*/
func SAdd(key interface{}, values ...interface{}) (err error) {
	rc := redisPool.Get()
	defer rc.Close()
	_, err = rc.Do(RedisKey_SADD, argsForm(values, key)...)
	return
}

/*
删除名称为key的set中的元素member
*/
func SRem(key string, value interface{}) (err error) {
	rc := redisPool.Get()
	defer rc.Close()
	_, err = rc.Do(RedisKey_SREM, key, value)
	return
}

/*
随机返回并删除名称为key的set中一个元素
*/
func SPop(key string) (replay interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return rc.Do(RedisKey_SPOP, key)
}

/*
移到集合元素
*/
func SMove(srcKey, dstKey string, value string) (err error) {
	rc := redisPool.Get()
	defer rc.Close()
	_, err = rc.Do(RedisKey_SMOVE, srcKey, dstKey, value)
	return
}

func SCard(key string) (int, error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_SCARD, key))
}

func SExistMember(key string, value string) (bool, error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Bool(rc.Do(RedisKey_SEXIST, key, value))
}

/*
求交集
*/
func SInter(keys ...interface{}) (replay []interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Values(rc.Do(RedisKey_SINTER, keys...))
}

/*
求交集并将交集保存到dstKey集合中
*/
func SInterStore(keys ...interface{}) (count int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	if len(keys) < 3 {
		return 0, fmt.Errorf("args not right")
	}
	return redis.Int(rc.Do(RedisKey_SINTERSTORE, keys...))
}

func SUnion(keys ...interface{}) (replay []interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Values(rc.Do(RedisKey_SUNION, keys...))
}

func SUnionStore(keys ...interface{}) (count int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	if len(keys) < 3 {
		return 0, fmt.Errorf("args not right")
	}
	return redis.Int(rc.Do(RedisKey_SUNIONSTORE, keys...))
}

func SDiff(keys ...interface{}) (replay []interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Values(rc.Do(RedisKey_SDIFF, keys...))
}

func SDiffStore(keys ...interface{}) (count int, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	if len(keys) < 3 {
		return 0, fmt.Errorf("args not right")
	}
	return redis.Int(rc.Do(RedisKey_SDIFFSTORE, keys...))
}

func SMembers(key string) (replay []interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Values(rc.Do(RedisKey_SMEMBERS, key))
}

func SRandMember(key string) (replay interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return rc.Do(RedisKey_SRANDMEMBER, key)
}
