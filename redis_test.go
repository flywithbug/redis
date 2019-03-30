package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)

type ListBody struct {
	Name string
}

func init() {
	op := DefaultOptions()
	Init(op)
}

func TestNewRedisPool(t *testing.T) {

	//rc := redisPool.Get()
	//defer rc.Close()
	//_, err := rc.Do("AUTH", "abcd")
	//fmt.Println(err)
	//fmt.Println(Set("a", "b", 5))
	//Remove("a")
	//fmt.Println(string(Get("a")))
	//list := make([]ListBody,0)
	//for i := 10; i < 20; i++ {
	//	l := ListBody{}
	//	l.Name = fmt.Sprintf("a:%d", i)
	//	fmt.Println(LPush("aSet1", l, false))
	//}
	//for i := 10; i < 20; i++ {
	//	l := ListBody{}
	//	l.Name = fmt.Sprintf("a:%d", i)
	//	fmt.Println(LPush("aSet2", l, false))
	//}
	//fmt.Println(redis.String(LBPop("aSet", false, 6)))
	//r, _ := LBPops(false, 6, "aSet")
	fmt.Println(redis.String(LRPopLPush("aSet", "Set")))
}

func TestLBRPopLPush(t *testing.T) {
	fmt.Println(redis.String(LRPopLPush("aSet", "Set")))
}

func TestPing(t *testing.T) {
	fmt.Println(Ping())
}

func TestSelect(t *testing.T) {
	fmt.Println(Select(1))
}
