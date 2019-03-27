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
	fmt.Println(SAdd("Set1", "aa", "bb"))
	fmt.Println(SAdd("Set", "abd3", "d23"))
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
	fmt.Println(redis.String(SRandMember("Set3")))
}
