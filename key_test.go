package redis

import (
	`fmt`
	`testing`
	`time`
)

func TestDump(t *testing.T) {
	reply ,err := Dump("Set0")
	fmt.Println(err)
	fmt.Println(reply)
}

func TestExists(t *testing.T) {
	fmt.Println(Exists("Set0"))
	fmt.Println(Exists("2322"))
}

func TestExpire(t *testing.T) {
	fmt.Println(Expire("a",10))
}

func TestExpireat(t *testing.T) {
	fmt.Println(Set("c","c"))
	fmt.Println(ExpireAt("c",time.Now().Add(time.Second*10)))
}
