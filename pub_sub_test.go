package redis

import (
	`fmt`
	`github.com/gomodule/redigo/redis`
	`testing`
	`time`
)

func init()  {
	op := DefaultOptions()
	Init(op)
}

func TestSubScribe(t *testing.T) {
	 SubScribe(func(message *redis.Message,err error) {
		if err != nil {
			fmt.Println(err)
		}else {
			fmt.Println(message.Channel)
		}
	},"channel1")
}

func TestPUnsubscribe(t *testing.T) {
	PUnsubscribe("channel1")
}

func TestPSubScribe(t *testing.T) {
	PSubScribe(func(message *redis.Message,err error) {
		fmt.Println(message.Channel)
	},"channel1")
}

func TestPubSub(t *testing.T) {
	fmt.Println(PubSub())
}

func TestPublish(t *testing.T) {
	fmt.Println("enter")
	go func() {
		SubScribe(func(message *redis.Message,err error) {
			if err == nil {
				fmt.Println("receive:",message.Channel)
				fmt.Println(string(message.Data))
			}
		},"channel1")
	}()
	for true {
		time.Sleep(time.Second*1)
		fmt.Println(Publish("channel1","hello"))
	}
}
