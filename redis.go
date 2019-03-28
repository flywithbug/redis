package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	redisPool *redis.Pool
)

type Options struct {
	Host        string
	Password    string
	DB          int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	Wait        bool
}


/*
 	db = 0
    maxIdle = 100
    maxActive = 500
    idleTimeout = 400
*/
func DefaultOptions() Options {
	return Options{
		Host:        "127.0.0.1:6379",
		Password:    "",
		DB:          0,
		MaxIdle:     100,
		MaxActive:   500,
		IdleTimeout: 400,
	}
}

func Init(options Options) {
	newRedisPool(options)
}

func newRedisPool(options Options) {
	redisPool = &redis.Pool{
		MaxIdle:     options.MaxIdle,
		MaxActive:   options.MaxActive,
		IdleTimeout: options.IdleTimeout,
		Wait:        options.Wait,
		Dial: func() (conn redis.Conn, e error) {
			timeOut := time.Duration(2) * time.Second
			conn, e = redis.Dial("tcp", options.Host, redis.DialPassword(options.Password),
				redis.DialDatabase(options.DB),
				redis.DialConnectTimeout(timeOut),
				redis.DialReadTimeout(timeOut),
				redis.DialWriteTimeout(timeOut))
			return conn, e
		},
	}
}

func Flush()  {
	rc := redisPool.Get()
	rc.Flush()
}
