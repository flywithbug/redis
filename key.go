package redis

import (
	`github.com/gomodule/redigo/redis`
	`time`
)




const (
	RedisKey_DEL = "DEL" //该命令用于在 key 存在时删除 key。
	RedisKey_DUMP = "DUMP"  //序列化给定 key ，并返回被序列化的值。

	RedisKey_EXISTS = "EXISTS"  //检查给定 key 是否存在。

	RedisKey_Expire = "Expire" //为给定 key 设置过期时间，以秒计。
	RedisKey_Expireat = "Expireat" //EXPIREAT 的作用和 EXPIRE 类似，都用于为 key 设置过期时间。UNIX 时间戳(unix timestamp)。

	RedisKey_PERSIST = "PERSIST" //移除 key 的过期时间，key 将持久保持。

	RedisKey_TTL = "TTL" //以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)


	RedisKey_KEYS = "KEYS"  //查找所有符合给定模式( pattern)的 key 。

	RedisKey_Move = "MOVE" //将当前数据库的 key 移动到给定的数据库 db 当中。


	RedisKey_RANDOMKEY = "RANDOMKEY" //从当前数据库中随机返回一个 key 。

	RedisKey_Rename = "Rename" //修改 key 的名称
	RedisKey_RenameNx = "RenameNx" //Renamenx 命令用于在新的 key 不存在时修改 key 的名称 。


	RedisKey_Type = "TYPE"

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


/*
Keys 命令用于查找所有符合给定模式 pattern 的 key 。。 正则匹配
*/
func Keys(key string)(reply []string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Strings(rc.Do(RedisKey_KEYS, key))
}

/*
 MOVE 命令用于将当前数据库的 key 移动到给定的数据库 db 当中。
*/
func Move(key string,db int)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_Move, key,db))
}

/*
PERSIST 命令用于移除给定 key 的过期时间，使得 key 永不过期。
*/
func Persist(key string)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_PERSIST, key))
}

/*
以秒为单位，返回给定 key 的剩余生存时间(TTL, time to live)。
*/
func TTL(key string)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_TTL, key))
}

/*
从当前数据库中随机返回一个 key 。
*/
func RandomKey()(reply string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_RANDOMKEY))
}

/*
修改 key 的名称
*/
func Rename(key ,newKey string)(reply string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_Rename, key,newKey))
}

/*
RenameNx 命令用于在新的 key 不存在时修改 key 的名称 。新key存在时 不修改 reply = 1时表示修改成功，=0时表示修改失败
*/
func RenameNx(key ,newKey string)(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_RenameNx, key,newKey))
}


/*
Redis Type 命令用于返回 key 所储存的值的类型。

none (key不存在)
string (字符串)
list (列表)
set (集合)
zset (有序集)
hash (哈希表)
*/
func Type(key string)(reply string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_Type, key))
}
