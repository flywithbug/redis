package redis

import (
	`fmt`
	`testing`
)

func TestBgReWriteAof(t *testing.T) {
	fmt.Println(BgReWriteAof())
}

func TestBgSave(t *testing.T) {
	fmt.Println(BgSave())
}

func TestClientKill(t *testing.T) {
	fmt.Println(ClientKill("aaa"))
}

func TestClientList(t *testing.T) {
	fmt.Println(ClientList())
}

//func TestClientSetName(t *testing.T) {
//	fmt.Println(ClientSetName("hello-world-connection"))
//	time.Sleep(time.Second*2)
//	fmt.Println(ClientGetName())
//}
