package redis

import (
	`fmt`
	`github.com/gomodule/redigo/redis`
	`testing`
)

func TestZAdd(t *testing.T) {
	fmt.Println(ZAdd("zAdd",1,"value1",2,"value2",3,"value3"))
	fmt.Println(ZAdd("zAdd1",1,"value1",2,"value2",3,"value3"))

}

func TestZCard(t *testing.T) {
	fmt.Println(ZCard("zAdd"))
}

func TestZCount(t *testing.T) {
	fmt.Println(ZCount("zAdd",1,2))
}

func TestZIncrBy(t *testing.T) {
	fmt.Println(ZIncrBy("zAdd","value1",1))
}

func TestZInterStore(t *testing.T) {
	fmt.Println(ZInterStore("zAdd2",2,"zAdd","zAdd1"))
}

func TestZLexCount(t *testing.T) {
	fmt.Println(ZLexCount("zAdd","[value1","[value3"))
}

func TestZRange(t *testing.T) {
	fmt.Println(redis.Strings(ZRange("zAdd",0,4,true)))
}
