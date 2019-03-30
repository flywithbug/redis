package redis

import (
	`github.com/gomodule/redigo/redis`
	`time`
)

/*
Redis 发布订阅(pub/sub)是一种消息通信模式：发送者(pub)发送消息，订阅者(sub)接收消息。
Redis 客户端可以订阅任意数量的频道。
*/


const (
	RedisKey_PUBLISH = "PUBLISH"
	RedisKey_SUBSCRIBE = "SUBSCRIBE"

	RedisKey_PUBSUB = "PUBSUB"
)

type CallBack func(message *redis.Message,err error)

func SubScribe(call CallBack,channel ...interface{})  {
	nSleep := 1
	for {
		connected := ListenSubscribe(call,channel...)
		if !connected {
			nSleep *=2
			if nSleep > 60 {
				nSleep = 60
			}
		}else {
			nSleep = 1
		}
		time.Sleep(time.Duration(nSleep)*time.Second)
	}
}

func ListenSubscribe(call CallBack,channel ...interface{})bool  {
	rc := redisPool.Get()
	psc := redis.PubSubConn{Conn:rc}
	psc.Subscribe(channel...)
	defer func() {
		psc.Close()
		rc.Close()
	}()
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			call(&v,nil)
		case redis.Subscription:
			println(v.Channel, v.Kind, v.Count)
		case error:
			call(nil,v)
			return true
		}
	}
	return false
}


func ListenPSubscribe(call CallBack,channel ...interface{})bool  {
	rc := redisPool.Get()
	psc := redis.PubSubConn{Conn:rc}
	psc.PSubscribe(channel...)
	defer func() {
		psc.Close()
		rc.Close()
	}()
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			call(&v,nil)
		case redis.Subscription:
			println(v.Channel,v.Kind,v.Count)
		case error:
			println(v.Error())
			call(nil,v)
			return true
		}
	}
	return false
}


/*
Redis PSubScribe 命令订阅一个或多个符合给定模式的频道。
每个模式以 * 作为匹配符，比如 it* 匹配所有以 it 开头的频道( it.news 、 it.blog 、 it.tweets 等等)。
news.* 匹配所有以 news. 开头的频道( news.it 、 news.global.today 等等)，诸如此类。
*/
func PSubScribe(call CallBack,channel ...interface{})   {
	nSleep := 1
	for {
		connected := ListenPSubscribe(call,channel...)
		if !connected {
			nSleep *=2
			if nSleep > 60 {
				nSleep = 60
			}
		}else {
			nSleep = 1
		}
		time.Sleep(time.Duration(nSleep)*time.Second)
	}
}

func PUnsubscribe(channel ...interface{})error  {
	rc := redisPool.Get()
	defer rc.Close()
	psc := redis.PubSubConn{Conn:rc}
	return psc.PUnsubscribe(channel...)
}

func PubSub()(reply []string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Strings(rc.Do(RedisKey_PUBSUB,"CHANNELS"))
}


func Publish(channel,value interface{})(reply int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Int(rc.Do(RedisKey_PUBLISH,channel,value))
}
