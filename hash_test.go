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
	fmt.Println(HSetNx("Hash","filed","value"))
}

func TestHDel(t *testing.T) {
	fmt.Println(HDel("Hash","filed"))
}

func TestHExists(t *testing.T) {
	fmt.Println(HExists("Hash","filed"))
	fmt.Println(HExists("Hash","filed123"))

}
