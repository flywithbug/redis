package redis

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	fmt.Println(Set("a", "b"))
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
