package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func init() {
	op := DefaultOptions()
	Init(op)
}

func TestSAdd(t *testing.T) {
	fmt.Println(SAdd("Set2", "aa1", "bb1"))
	fmt.Println(SAdd("Set1", "abd31", "d231"))
}

func TestSExistMember(t *testing.T) {
	fmt.Println(SExistMember("Set", "a"))
}

func TestSMove(t *testing.T) {
	fmt.Println(SMove("Set", "Set1", "c"))
}

func TestSCard(t *testing.T) {
	fmt.Println(SCard("Set"))
}

func TestSInter(t *testing.T) {
	//replay, err := SInterStore("Set2", "Set", "Set1")
	fmt.Println(SInterStore("Set2", "Set", "Set1"))
}

func TestSRem(t *testing.T) {
	fmt.Println(SRem("Set1","a","bb","b","c"))
}

func TestSUnion(t *testing.T) {
	fmt.Println(redis.Strings(SUnion("Set", "Set1")))
}

func TestSUnionStore(t *testing.T) {
	fmt.Println(SUnionStore("Set0", "Set", "Set1"))
}

func TestSDiff(t *testing.T) {
	fmt.Println(redis.Strings(SDiff("Set", "Set1")))

}

func TestSDiffStore(t *testing.T) {
	fmt.Println(SDiffStore("Set3", "Set", "Set1"))
}

func TestSMembers(t *testing.T) {
	fmt.Println(redis.Strings(SMembers("Set3")))
}

func TestSRandMember(t *testing.T) {
	fmt.Println(redis.Strings(SRandMember("Set0",6)))
}

func TestSScan(t *testing.T) {
	fmt.Println(SScan("Set0",100,"a*"))
}
