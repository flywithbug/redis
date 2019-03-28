# redis


redigo 封装



``` 
op := Options{
      		Host:        "127.0.0.1:6379",
      		Password:    "",
      		DB:          0,
      		MaxIdle:     100,
      		MaxActive:   500,
      		IdleTimeout: 400,
      	}
Init(op) //初始化Redis

func TestHSet(t *testing.T) {
	fmt.Println(HSet(1,3,9))
}

func TestHMSet(t *testing.T) {
	fmt.Println(HMSet("Hash","filed3","value3","filed2","value2"))

	fmt.Println(HMSet("Hash","filed3","value3","filed2","value2"))
}

func TestHSetNx(t *testing.T) {
	fmt.Println(HSetNx("Hash","field","value"))
}

func TestHDel(t *testing.T) {
	fmt.Println(HDel("Hash","filed"))
}

func TestHExists(t *testing.T) {
	fmt.Println(HExists("Hash","filed"))
	fmt.Println(HExists("Hash","filed123"))
}

func TestHGet(t *testing.T) {
	fmt.Println(HGet("Hash","filed"))
}

func TestHGetAll(t *testing.T) {
	fmt.Println(HGetAll("Hash"))
}

func TestHIncrBy(t *testing.T) {
	//fmt.Println(HSetNx("Hash","incre",1))
	fmt.Println(HIncrBy("Hash","incre",1))
}

func TestHIncrByFloat(t *testing.T) {
	//fmt.Println(HSetNx("Hash","InCre",0))
	fmt.Println(HIncrByFloat("Hash","InCre",2.4))
}

func TestHKeys(t *testing.T) {
	fmt.Println(redis.Strings(HKeys("Hash")))
}

func TestHLen(t *testing.T) {
	fmt.Println(HLen("Hash"))
}

func TestHMGet(t *testing.T) {
	fmt.Println(redis.Strings(HMGet("Hash","field","InCre")))
}

func TestHValues(t *testing.T) {
	fmt.Println(redis.Strings(HValues("Hash")))
}


```
