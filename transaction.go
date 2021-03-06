package redis

/*
Redis 事务可以一次执行多个命令， 并且带有以下两个重要的保证：
 - 批量操作在发送 EXEC 命令前被放入队列缓存。
 - 收到 EXEC 命令后进入事务执行，事务中任意命令执行失败，其余的命令依然被执行。
 - 在事务执行过程，其他客户端提交的命令请求不会插入到事务执行命令序列中。
一个事务从开始到执行会经历以下三个阶段：
 - 开始事务。
 - 命令入队。
 - 执行事务。
*/


const (
	RedisKey_MULTI = "MULTI"
	RedisKey_DISCARD = "DISCARD"
)


/*
Redis Discard 命令用于取消事务，放弃执行事务块内的所有命令
*/
func Discard()  {
	rc := redisPool.Get()
	defer rc.Close()
}
