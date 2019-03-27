package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

const (
	RedisKey_LPUSH = "LPUSH" //表头
	RedisKey_RPUSH = "RPUSH" //表尾

	RedisKey_LLEN = "LLEN" //表长度

	RedisKey_LRANGE = "LRANGE" //获取range 内的元素
	RedisKey_LTRIM  = "LTRIM"  //列表进行修剪(trim)

	RedisKey_LINDEX = "LINDEX" //返回名称为key的list中index位置的元素
	RedisKey_LSET   = "LSET"   //给名称为key的list中index位置的元素赋值
	RedisKey_LREM   = "LREM"   //删除count个key的list中值为value的元素

	RedisKey_LPOP  = "LPOP"  //返回表头 并删除
	RedisKey_RPOP  = "RPOP"  //返回表尾 并删除
	RedisKey_BRPOP = "BRPOP" //命令移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	RedisKey_BLPOP = "BLPOP" //命令移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。

	RedisKey_RPOPLPUSH  = "RPOPLPUSH"  //返回并删除名称为srckey的list的尾元素，并将该元素添加到名称为dstkey的list的头部
	RedisKey_BRPOPLPUSH = "BRPOPLPUSH" // Brpoplpush 命令从列表中取出最后一个元素，并插入到另外一个列表的头部； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
)

/*
-----------------------------------------------------------------------------------------
------------------------------- 数组操作 首字母 L 表示List
-----------------------------------------------------------------------------------------
*/

/*数组中添加元素
head  true 表示添加在表头部分 false 表示在数组末尾添加元素
*/
func LPush(key string, value interface{}, head bool) (err error) {
	rc := redisPool.Get()
	defer rc.Close()
	if head {
		_, err = rc.Do(RedisKey_LPUSH, key, value)
	} else {
		_, err = rc.Do(RedisKey_RPUSH, key, value)
	}
	return
}

//返回并删除key list的首元素
func LPop(key string, head bool) (replay interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	if head {
		return rc.Do(RedisKey_LPOP, key)
	} else {
		return rc.Do(RedisKey_RPOP, key)
	}
}

//返回并删除key list的首元素 timeout second
func LBPop(key string, head bool, timeout int) (replay interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	commName := RedisKey_BRPOP
	if head {
		commName = RedisKey_BLPOP
	}
	list, err := redis.Values(rc.Do(commName, key, timeout))
	if err != nil {
		return
	}
	if len(list) == 2 {
		return list[1], err
	}
	return nil, fmt.Errorf("result not right")
}

//返回名称为key的list的长度
func LLen(key string) int {
	rc := redisPool.Get()
	defer rc.Close()
	num, _ := redis.Int64(rc.Do(RedisKey_LLEN, key))
	return int(num)
}

//返回名称为key的list中start至end之间的元素
func LRange(key string, start, end int) (values []interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	values, err = redis.Values(rc.Do(RedisKey_LRANGE, key, start, end))
	return
}

/*
对列表进行裁剪，只保留区域内的元素，不再区域内的元素将被删除
下标 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。
你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
*/
func LTrim(key string, start, end int) (err error) {
	rc := redisPool.Get()
	defer rc.Close()
	_, err = rc.Do(RedisKey_LTRIM, key, start, end)
	return
}

/*
返回名称为key的list中index位置的元素
*/
func LIndex(key string, index int) (replay interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return rc.Do(RedisKey_LINDEX, key, index)
}

//给名称为key的list中index位置的元素赋值
func LSet(key string, index int, value interface{}) (replay interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return rc.Do(RedisKey_LSET, key, index, value)
}

//删除count个key的list中值为value的元素
func LRem(key string, count int, value interface{}) (replay interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return rc.Do(RedisKey_LREM, key, count, value)
}

/*
 R pop L push 命令用于移除列表的最后一个元素，并将该元素添加到另一个列表并返回。
返回并删除名称为srckey的list的尾元素，并将该元素添加到名称为dstkey的list的头部
*/
func LRPopLPush(srcKey, dstKey string) (replay interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return rc.Do(RedisKey_RPOPLPUSH, srcKey, dstKey)
}

/*
 Brpoplpush 命令从列表中取出最后一个元素，并插入到另外一个列表的头部； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
*/
func LBRPopLPush(srcKey, dstKey string) (replay interface{}, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return rc.Do(RedisKey_BRPOPLPUSH, srcKey, dstKey)
}
