## Socket 编程

### 1 网络的基本概念

```html
	网络编程的目的：直接或间接地通过网络协议与其他计算机进行通讯
	网络编程中两个主要问题：
		如何准确定位网络上一台或多台主机（通过IP地址）
		找到主机后如何进行数据传输（有OSI模型和TCP/IP模型）
	OSI模型将网络分为7层，过于理想化，未能广泛推广
	TCP/IP是事实上的国际标准
    OSI				TCP/IP		TCP/IP对应的协议
    应用层			应用层			HTTP、FTP、DNS
    表示层		
    会话层		
    传输层			传输层			TCP    UDP
    网络层			网络层			IP ARP ICMP
    数据链路层	 物理数据层	Link
    物理层		
	IP分类：最开始是32位整数，后来用圆点隔开，分成4段，8位一段
    	A类：保留给政府结构，1.0.0.1 ~ 126.255.255.254
    	B类：分配给中型企业，1.0.0.1 ~ 126.255.255.254
    	C类：分配给任何需要的个人，192.0.0.1 ~ 223.255.255.254
    	D类：用于组播，224.0.0.1 ~ 239.255.255.254
    	E类：用于实验，240.0.0.1 ~ 255.255.255.254
	回收地址：127.0.0.1，指本地机，一般用于测试使用
```

### 2 http和https

```html
HTTP发送明文不安全
考虑加密，对称加密(怎么加，就怎么解)
```

```go
// http 对称加密
1 服务端创建一对公钥和私钥
2 客户端向服务端发送请求，获取公钥
3 服务端把公钥返回给客户端
4 客户端接收公钥
5 客户端创建随机字符来做对称密钥
6 客户端使用公钥对对称密码加密，发送给服务端
7 服务端接收对称密钥
// https CA机构
服务端申请证书
```

### 3 socket

```html
1 Socket又称“套接字”，应用程序通常通过“套接字”向网络发出请求或者应答网络请求
2常用的Socket类型有两种：流式Socket和数据报式Socket，流式是一种面向连接的Socket，针对于面向连接的TCP服务应用，数据报式Socket是一种无连接的Socket，针对于无连接的UDP服务应用
TCP：比较靠谱，面向连接，比较慢
UDP：不是太靠谱，比较快
```

### 4 tcp 

```html
实现服务端与客户端通信，保持连接，服务器接收到exit时再断开连接
```

> 服务端

```go
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 1.创建TCP服务端监听
	listenner, err := net.Listen("tcp", "0.0.0.0:8888") // tcp 表示链接类型是tcp
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listenner.Close()
	// 2.服务端不断等待请求处理
	for {
		// 阻塞等待客户端连接
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ClientConn(conn)
	}
}

// 处理服务端逻辑
func ClientConn(conn net.Conn) {
	defer conn.Close()
	// 获取客户端地址
	ipAddr := conn.RemoteAddr().String()
	fmt.Println(ipAddr, "连接成功")
	// 缓冲区 接收数据
	buf := make([]byte, 1024)
	for {
		// n是读取的长度
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 切出有效数据
		result := buf[:n]
		fmt.Printf("接收到数据，来自[%s]    [%d]:%s\n", ipAddr, n, string(result))
		// 接收到exit，退出连接
		if string(result) == "exit" {
			fmt.Println(ipAddr, "退出连接")
			return
		}
		// 回复客户端
		conn.Write([]byte(strings.ToUpper(string(result))))
	}
}
```

> 客户端

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	// 1.连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// 缓冲区
	buf := make([]byte, 1024)
	for {
		fmt.Printf("请输入发送的内容：")
		fmt.Scan(&buf)
		fmt.Printf("发送的内容是：%s\n", string(buf))
		// 发送数据
		conn.Write(buf)

		// 接收服务端返回信息
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := buf[:n]
		fmt.Printf("接收到数据:%s\n", string(result))
	}
}
```

### 5 udp

`服务端`

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	// UDP的服务端监听
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for {
		// 缓冲区
		var data [1024]byte
		// 接收UDP传输
		count, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("data:%s addr:%v \n", string(data[0:count]), addr)
		// 返回信息
		_, err = listen.WriteToUDP([]byte("666"), addr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
```

`客户端`

```go
package main

import (
   "fmt"
   "net"
)

// UDP客户端
func main() {
   // 1.连接服务端
   conn, err := net.DialUDP("udp4", nil, &net.UDPAddr{
      IP:   net.IPv4(127, 0, 0, 1),
      Port: 8080,
   })
   if err != nil {
      fmt.Println(err)
      return
   }
   defer conn.Close()
   // 写数据到服务端
   _, err = conn.Write([]byte("老铁"))
   if err != nil {
      fmt.Println(err)
      return
   }
   data := make([]byte, 16)
   count, addr, err := conn.ReadFromUDP(data)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println("count:", count)
   fmt.Println("addr:", addr)
   fmt.Println("data:", string(data))
}
```



