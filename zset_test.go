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
	fmt.Println(redis.Strings(ZRevRange("zAdd",0,4,true)))
}

func TestZRangeByLex(t *testing.T) {
	fmt.Println(redis.Strings(ZRangeByLex("zAdd","-","[value3")))
}

func TestZRangeByScore(t *testing.T) {
	fmt.Println(ZAdd("salary",2500,"Jack"))
	fmt.Println(ZAdd("salary",5000,"Tom"))
	fmt.Println(ZAdd("salary",12000,"Peter"))

	fmt.Println(redis.Strings(ZRangeByScore("salary","-inf","+inf",false)))
	fmt.Println(redis.Strings(ZRangeByScore("salary","-inf","+inf",true)))
	fmt.Println(redis.Strings(ZRangeByScore("salary","-inf","5000",true)))
	fmt.Println(redis.Strings(ZRangeByScore("salary","(5000",40000,true)))
	fmt.Println(redis.Strings(ZRangeByScore("salary",5000,40000,true)))

}

func TestZRank(t *testing.T) {
	fmt.Println(ZRank("salary","Tom"))
	fmt.Println(ZRank("salary","Jack"))

	fmt.Println(ZRank("salary","Peter"))

}

func TestZRem(t *testing.T) {
	fmt.Println(ZAdd("salary",2500,"Jack"))
	fmt.Println(ZAdd("salary",5000,"Tom"))
	fmt.Println(ZAdd("salary",12000,"Peter"))
	fmt.Println(ZRem("salary","Tom","Peter"))
}

func TestZRemRangeByLex(t *testing.T) {
	fmt.Println(ZAdd("salary",2500,"Jack"))
	fmt.Println(ZAdd("salary",5000,"Tom"))
	fmt.Println(ZAdd("salary",12000,"Peter"))
	fmt.Println(ZRemRangeByLex("salary","[Jack","[x"))
}

func TestZRemRangeByRank(t *testing.T) {
	fmt.Println(ZRemRangeByRank("salary",0,3))
}
