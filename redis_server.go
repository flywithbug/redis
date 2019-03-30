package redis

import (
	`github.com/gomodule/redigo/redis`
)

//TODO

const  (
	RedisKey_CLIENT = "CLIENT"

	RedisKey_BGREWRITEAOF = "BGREWRITEAOF"
	RedisKey_BGSAVE = "BGSAVE"

	RedisKey_CLIENT_LIST = "LIST"


	RedisKey_CLIENT_GETNAME = "GETNAME"
	RedisKey_CLIENT_SETNAME = "SETNAME"

)

/*
Redis BgReWriteAof 命令用于异步执行一个 AOF（AppendOnly File） 文件重写操作。重写会创建一个当前 AOF 文件的体积优化版本。
即使 BgReWriteAof 执行失败，也不会有任何数据丢失，因为旧的 AOF 文件在 BgReWriteAof 成功之前不会被修改。
注意：从 Redis 2.4 开始， AOF 重写由 Redis 自行触发， BgReWriteAof 仅仅用于手动触发重写操作。
*/
func BgReWriteAof()(replay string, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_BGREWRITEAOF))
}

/*
Redis BgSave 命令用于在后台异步保存当前数据库的数据到磁盘。
BgSave 命令执行之后立即返回 OK ，然后 Redis fork 出一个新子进程，
原来的 Redis 进程(父进程)继续处理客户端请求，而子进程则负责将数据保存到磁盘，然后退出。
*/
func BgSave()(replay string, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_BGSAVE))
}

/*
Redis Client Kill 命令用于关闭客户端连接。
ip:port
*/
func ClientKill(ipPort string)(replay string, err error) {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_CLIENT,RedisKey_SCRIPT_KILL,ipPort))
}

func ClientList()(replay string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_CLIENT,RedisKey_CLIENT_LIST))
}

//func ClientSetName(name string)(replay string,err error)  {
//	rc := redisPool.Get()
//	defer rc.Close()
//	return redis.String(rc.Do(RedisKey_CLIENT,RedisKey_CLIENT_SETNAME,name))
//}
//
//func ClientGetName()(replay string,err error)  {
//	rc := redisPool.Get()
//	defer rc.Close()
//	return redis.String(rc.Do(RedisKey_CLIENT,RedisKey_CLIENT_GETNAME))
//}



