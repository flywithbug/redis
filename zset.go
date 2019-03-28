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

