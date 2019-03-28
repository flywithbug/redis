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
