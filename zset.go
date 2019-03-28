package redis

import (
	`github.com/gomodule/redigo/redis`
)

/*
Redis 有序集合和集合一样也是string类型元素的集合,且不允许重复的成员。
不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。
有序集合的成员是唯一的,但分数(score)却可以重复。
集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是O(1)。
集合中最大的成员数为 232 - 1 (4294967295, 每个集合可存储40多亿个成员)。
*/

const (
	RedisKey_ZADD = "ZADD"
	RedisKey_ZCARD = "Zcard"
	RedisKey_ZCOUNT = "ZCOUNT"
	RedisKey_ZINCRBY = "ZINCRBY"

	RedisKey_ZINTERSTORE = "ZINTERSTORE"

	RedisKey_ZLEXCOUNT = "ZLEXCOUNT" //计算字典区间成员数(分数都相同，按照字典排序)


	RedisKey_ZRANGE = "ZRANGE"

)

/*
Redis ZAdd 命令用于将一个或多个成员元素及其分数值加入到有序集当中。
如果某个成员已经是有序集的成员，那么更新这个成员的分数值，并通过重新插入这个成员元素，来保证该成员在正确的位置上。
分数值可以是整数值或双精度浮点数。
如果有序集合 key 不存在，则创建一个空的有序集并执行 ZAdd 操作。
当 key 存在但不是有序集类型时，返回一个错误。
	score is float or int
	values is interface{}
*/
func ZAdd(key interface{},scoreValues ...interface{})(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_ZADD, argsForm(scoreValues,key)...))
}


/*
	ZCard 命令用于计算集合中元素的数量。
*/
func ZCard(key interface{})(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_ZCARD, key))
}


/*
ZCount 命令用于计算有序集合中指定分数区间的成员数量。
*/
func ZCount(key interface{},min,max float64)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_ZCOUNT, key,min,max))
}


/*
ZIncrBy 命令对有序集合中指定成员的分数加上增量 increment
可以通过传递一个负数值 increment ，让分数减去相应的值，比如 ZINCRBY key -5 member ，就是让 member 的 score 值减去 5 。
当 key 不存在，或分数不是 key 的成员时， ZINCRBY key increment member 等同于 ZADD key increment member 。
当 key 不是有序集类型时，返回一个错误。
分数值可以是整数值或双精度浮点数。
*/

func ZIncrBy(key,value interface{},incr float64)(reply float64,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Float64(rc.Do(RedisKey_ZINCRBY, key,incr,value))
}


/*
Redis Zinterstore 命令计算给定的一个或多个有序集的交集，其中给定 key 的数量必须以 numKeys 参数指定，并将该交集(结果集)储存到 destination 。
默认情况下，结果集中某个成员的分数值是所有给定集下该成员分数值之和。
语法
redis ZInterStore 命令基本语法如下：
redis 127.0.0.1:6379> ZInterStore destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]
*/
func ZInterStore(dst interface{},numKey int ,src... interface{})(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_ZINTERSTORE, argsForm(src,dst,numKey)...))
}


/*
在有序集合中计算指定字典区间内成员数量
*/
func ZLexCount(key ,min,max interface{})(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_ZLEXCOUNT, key,min,max))
}


/*
Redis ZRange 返回有序集中，指定区间内的成员。
其中成员的位置按分数值递增(从小到大)来排序。
具有相同分数值的成员按字典序(lexicographical order )来排列。
如果你需要成员按
值递减(从大到小)来排列，请使用 ZREVRANGE 命令。
下标参数 start 和 stop 都以 0 为底，也就是说，以 0 表示有序集第一个成员，以 1 表示有序集第二个成员，以此类推。
你也可以使用负数下标，以 -1 表示最后一个成员， -2 表示倒数第二个成员，以此类推。
*/
func ZRange(key interface{},min,max int,needScore bool)(reply []interface{},err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	if needScore {
		return redis.Values(rc.Do(RedisKey_ZRANGE,key,min,max,"WithSCORES"))
	}
	return redis.Values(rc.Do(RedisKey_ZRANGE,key,min,max))
}

