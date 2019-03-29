package redis

import (
	`fmt`
	`testing`
)

func TestPFAdd(t *testing.T) {
	fmt.Println(PFAdd("pfAddKey","Jack","Tom"))
	fmt.Println(PFAdd("pfAddKey1","Jack","Tom"))

}

func TestPFCount(t *testing.T) {
	fmt.Println(PFCount("pfAddKey"))
	fmt.Println(PFCount("pfAddKey1"))

}

func TestPFMerge(t *testing.T) {
	fmt.Println(PFMerge("pfAddKey0","pfAddKey1","pfAddKey"))
	fmt.Println(PFCount("pfAddKey0"))
}
