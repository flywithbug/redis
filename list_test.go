package redis

import (
	"fmt"
	"testing"
)

func init() {
	op := DefaultOptions()
	Init(op)
}

func TestLBPops(t *testing.T) {

	fmt.Println(LBPop("aSet", true, 10))

}
