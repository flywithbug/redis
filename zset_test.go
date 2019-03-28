package redis

import (
	`fmt`
	`testing`
)

func TestZAdd(t *testing.T) {
	fmt.Println(ZAdd("zAdd",1,"value1",2,"value2"))
}

func TestZCard(t *testing.T) {
	fmt.Println(ZCard("zAdd"))
}
