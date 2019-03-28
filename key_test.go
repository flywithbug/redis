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

func TestKeys(t *testing.T) {
	fmt.Println(Keys("S*"))
}

func TestMove(t *testing.T) {
	fmt.Println(Move("Set",1))
}

func TestPersist(t *testing.T) {
	fmt.Println(Set("c","c"))
	fmt.Println(ExpireAt("c",time.Now().Add(time.Second*6)))
	time.Sleep(time.Second*2)
	fmt.Println(TTL("c"))
	fmt.Println(Persist("c"))
}

func TestRandomKey(t *testing.T) {
	fmt.Println(RandomKey())
}

func TestRename(t *testing.T) {
	//fmt.Println(Set("d","d"))
	//fmt.Println(Rename("c","e"))

	fmt.Println(RenameNx("e","a"))
}

func TestType(t *testing.T) {
	fmt.Println(Type("a"))
	fmt.Println(Type("aSet1"))
	fmt.Println(Type("Set0"))

}
