package redis

/*
Redis 发布订阅(pub/sub)是一种消息通信模式：发送者(pub)发送消息，订阅者(sub)接收消息。
Redis 客户端可以订阅任意数量的频道。
*/

/*
Subscribe 命令用于订阅给定的一个或多个频道的信息。
*/


const (
	RedisKey_PUBLISH = "PUBLISH"
	RedisKey_SUBSCRIBE = "SUBSCRIBE"
)

