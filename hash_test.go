package redis

import (
	`fmt`
	`testing`
)

func TestHSet(t *testing.T) {
	fmt.Println(HSet("Hash","filed","value"))
}

func TestHMSet(t *testing.T) {
	fmt.Println(HMSet("Hash","filed3","value3","filed2","value2"))
}

func TestHSetNx(t *testing.T) {
	fmt.Println(HSetNx("Hash","field","value"))
}

func TestHDel(t *testing.T) {
	fmt.Println(HDel("Hash","filed"))
}

func TestHExists(t *testing.T) {
	fmt.Println(HExists("Hash","filed"))
	fmt.Println(HExists("Hash","filed123"))
}

func TestHGet(t *testing.T) {
	fmt.Println(HGet("Hash","filed"))
}

func TestHGetAll(t *testing.T) {
	fmt.Println(HGetAll("Hash"))
}

func TestHIncrBy(t *testing.T) {
	//fmt.Println(HSetNx("Hash","incre",1))
	fmt.Println(HIncrBy("Hash","incre",1))
}
