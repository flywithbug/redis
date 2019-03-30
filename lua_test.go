package redis

import (
	`fmt`
	`github.com/gomodule/redigo/redis`
	`testing`
)

func TestEval(t *testing.T) {
	fmt.Println(redis.Strings(Eval("return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}", "key1","key2","first","second")))
}

func TestScriptLoad(t *testing.T) {
	reply ,_ := ScriptLoad( "return 'hello Mo'",)
	fmt.Println(reply)
	fmt.Println(EvalSha(reply,0))
	fmt.Println(ScriptExists(reply))
}

func TestScriptFlush(t *testing.T) {
	fmt.Println(ScriptFlush())
}
