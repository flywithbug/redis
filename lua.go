package redis

import (
	`github.com/gomodule/redigo/redis`
)

const (
	RedisKey_EVAL = "EVAL"
	RedisKey_EVALSHA = "EVALSHA"
	RedisKey_SCRIPT = "SCRIPT"

	RedisKey_SCRIPT_EXISTS = "EXISTS"

	RedisKey_SCRIPT_FLUSH = "FLUSH"
	RedisKey_SCRIPT_KILL = "KILL"
	RedisKey_SCRIPT_LOAD = "LOAD"

)

/*
Redis 脚本使用 Lua 解释器来执行脚本。 Redis 2.6 版本通过内嵌支持 Lua 环境。执行脚本的常用命令为 EVAL。
*/
func Eval(script string,keys ...interface{})(reply []interface{},err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Values(rc.Do(RedisKey_EVAL,argsForm(keys,script,len(keys)/2)...))
}

/*
Redis EvalSha 命令根据给定的 sha1 校验码，执行缓存在服务器中的脚本。
将脚本缓存到服务器的操作可以通过 SCRIPT LOAD 命令进行。
这个命令的其他地方，比如参数的传入方式，都和 EVAL 命令一样。
*/
func EvalSha(sha string,numKeys int)(reply string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_EVALSHA,sha,numKeys))
}

func ScriptLoad(script string,keys... interface{})(reply string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_SCRIPT,argsForm(keys,"LOAD",script)...))
}

//func ScriptLoad(script string)(reply string,err error)  {
//	rc := redisPool.Get()
//	defer rc.Close()
//	return redis.String(rc.Do(RedisKey_SCRIPT,RedisKey_SCRIPT_LOAD,script))
//}


func ScriptExists(sha... interface{})(reply []int,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.Ints(rc.Do(RedisKey_SCRIPT,argsForm(sha,RedisKey_SCRIPT_EXISTS)...))
}

func ScriptFlush()(reply string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_SCRIPT_FLUSH))
}

func ScriptKill()(reply string,err error)  {
	rc := redisPool.Get()
	defer rc.Close()
	return redis.String(rc.Do(RedisKey_SCRIPT_KILL))
}
