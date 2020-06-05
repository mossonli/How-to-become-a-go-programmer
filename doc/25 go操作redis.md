## Go操作redis

### 1 当下nosql的经典应用

```go
1 目前是sql和nosql混用
2 阿里巴巴商品信息的存储方式
	商品的基本信息用 mysql
	商品的详情信息 描述性信息 mongodb
	商品图片	TFS	
	关键字		 ISearch 搜索引擎
	商品支付	支付接口
	商品热点信息	redis
```

### go操作redis

`1 键值操作`

```go
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}

func main()  {
	// 建立连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleError(err, "redis 连接")
	defer conn.Close()
	// 键值操作
	// 1 设置键值
	res, err := conn.Do("set", "abc", 100) // 设置key value
	HandleError(err, "set key value")
	fmt.Printf("set result :%s\n", res)
	// 2 获取键值
	ret, err := conn.Do("get", "abc")
	HandleError(err, "get value")
	fmt.Printf("get result :%s\n", ret)
	// res 是字节类型
	// 将结果转换成int
	r, err := redis.Int(ret, err)
	HandleError(err, "redis 字节转int")
	fmt.Println(r)
}
```

`2 批量操作`

```go
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}

func main()  {
	// 建立连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleError(err, "redis 连接")
	fmt.Println("连接成功")
	defer conn.Close()
	// 批量操作
	// 存数据
	res, err := conn.Do("mset", "x", 100, "y", 200)
	fmt.Printf("res:%v, err:%v\n", res, err)
	HandleError(err, "mset")
	// 取数据
	resg, err := conn.Do("mget", "x", "y")
	HandleError(err, "mget bytes")
	fmt.Printf("res:%v, err:%v\n", resg, err)
	r, err := redis.Ints(resg, err)
	HandleError(err, "mget")
	fmt.Println(r, err)
	for k, v := range r{
		fmt.Println(k, v)
	}
}
```

`3 过期时间的设置`

```go
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}

func main()  {
	// 建立连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandleError(err, "redis 连接")
	fmt.Println("连接成功")
	defer conn.Close()
	// 设置过期时间
	res, err := conn.Do("expire", "abc", 10)
	HandleError(err, "设置超时时间")
	fmt.Println(res)

}
```

`4 redis的链接池`

```go
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}

func main()  {
	// 建立连接池
	pool := &redis.Pool{
		// 设置最大的空闲数
		MaxIdle: 16,
		// 设置最大的活跃数，0代表无限
		MaxActive: 0,
		// 闲置空闲时间，单位秒
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	defer pool.Close()
	// 从池子里面取连接
	conn := pool.Get()
	defer conn.Close()
	res, err := conn.Do("set", "q", 100)
	res, err = conn.Do("get", "q")
	r, err := redis.Int(res, err)
	fmt.Println(r)
}
```



