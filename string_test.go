package redis

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	fmt.Println(Set("b", "b"))
}

func TestSetExp(t *testing.T) {
	fmt.Println(SetExp("c", "b", 0))
}

func TestGet(t *testing.T) {
	fmt.Println(Remove("c"))
}

func TestGetSet(t *testing.T) {
	fmt.Println(GetSet("a", "2323"))
}

func TestMGet(t *testing.T) {
	fmt.Println(MGet("a", "b"))
}

func TestSetNx(t *testing.T) {
	fmt.Println(SetNx("a", "b"))
	fmt.Println(SetNx("c", "b"))
}

func TestSetEx(t *testing.T) {
	fmt.Println(SetEx("d", "aaa23",50))

}

func TestMSet(t *testing.T) {
	fmt.Println(MSet("a","b","c","d","e","f"))
}

func TestMSetNx(t *testing.T) {
	fmt.Println(MSet("a","aa","b","bb","c","cc"))

	fmt.Println(MGet("a","b","c"))
}

func TestIncr(t *testing.T) {
	fmt.Println(Incr("incrKey"))
}

func TestIncrBy(t *testing.T) {
	fmt.Println(IncrBy("incrKey",5))
}

func TestDecr(t *testing.T) {
	fmt.Println(Decr("incrKey"))
	fmt.Println(DecrBy("incrKey",2))
}

func TestAppend(t *testing.T) {
	fmt.Println(Append("a","2322"))
}

func TestSubStr(t *testing.T) {
	fmt.Println(SubStr("a",0,5))

}
