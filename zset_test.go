package redis

import (
	`fmt`
	`testing`
)

func TestZAdd(t *testing.T) {
	fmt.Println(ZAdd("zAdd",1.1,"value1",2.2,"value2",3.3,"value3"))
}

func TestZCard(t *testing.T) {
	fmt.Println(ZCard("zAdd"))
}

func TestZCount(t *testing.T) {
	fmt.Println(ZCount("zAdd",1,2))
}
