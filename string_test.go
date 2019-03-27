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
