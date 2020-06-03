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

```

