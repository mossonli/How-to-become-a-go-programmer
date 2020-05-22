## Websocket[长链接]

```html
1 WebSocket是一种在单个TCP连接上进行全双工通信的协议
2	WebSocket使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据
3	在WebSocket API中，浏览器和服务器只需要完成一次握手，两者之间就直接可以创建持久性的连接，并进行双向数据传输
4	需要安装第三方包：
		cmd中：go get -u -v github.com/gorilla/websocket
```

`local.html`

```html
<!DOCTYPE html>
<html>
<head>
    <title></title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8">
    <style>
        p {
            text-align: left;
            padding-left: 20px;
        }
    </style>
</head>
<body>
<div style="width: 800px;height: 600px;margin: 30px auto;text-align: center">
    <h1>聊天室</h1>
    <div style="width: 800px;border: 1px solid gray;height: 300px;">
        <div style="width: 200px;height: 300px;float: left;text-align: left;">
            <p><span>当前在线:</span><span id="user_num">0</span></p>
            <div id="user_list" style="overflow: auto;">
            </div>
        </div>
        <div id="msg_list" style="width: 598px;border:  1px solid gray; height: 300px;overflow: scroll;float: left;">
        </div>
    </div>
    <br>
    <textarea id="msg_box" rows="6" cols="50" onkeydown="confirm(event)"></textarea><br>
    <input type="button" value="发送" onclick="send()">
</div>
</body>
</html>
<script type="text/javascript">
    var uname = prompt('请输入用户名', 'user' + uuid(8, 16));
    var ws = new WebSocket("ws://127.0.0.1:8080/ws");
    ws.onopen = function () {
        var data = "系统消息：建立连接成功";
        listMsg(data);
    };
    ws.onmessage = function (e) {
        var msg = JSON.parse(e.data);
        var sender, user_name, name_list, change_type;
        switch (msg.type) {
            case 'system':
                sender = '系统消息: ';
                break;
            case 'user':
                sender = msg.from + ': ';
                break;
            case 'handshake':
                var user_info = {'type': 'login', 'content': uname};
                sendMsg(user_info);
                return;
            case 'login':
            case 'logout':
                user_name = msg.content;
                name_list = msg.user_list;
                change_type = msg.type;
                dealUser(user_name, change_type, name_list);
                return;
        }
        var data = sender + msg.content;
        listMsg(data);
    };
    ws.onerror = function () {
        var data = "系统消息 : 出错了,请退出重试.";
        listMsg(data);
    };
    function confirm(event) {
        var key_num = event.keyCode;
        if (13 == key_num) {
            send();
        } else {
            return false;
        }
    }
    function send() {
        var msg_box = document.getElementById("msg_box");
        var content = msg_box.value;
        var reg = new RegExp("\r\n", "g");
        content = content.replace(reg, "");
        var msg = {'content': content.trim(), 'type': 'user'};
        sendMsg(msg);
        msg_box.value = '';
    }
    function listMsg(data) {
        var msg_list = document.getElementById("msg_list");
        var msg = document.createElement("p");
        msg.innerHTML = data;
        msg_list.appendChild(msg);
        msg_list.scrollTop = msg_list.scrollHeight;
    }
    function dealUser(user_name, type, name_list) {
        var user_list = document.getElementById("user_list");
        var user_num = document.getElementById("user_num");
        while(user_list.hasChildNodes()) {
            user_list.removeChild(user_list.firstChild);
        }
        for (var index in name_list) {
            var user = document.createElement("p");
            user.innerHTML = name_list[index];
            user_list.appendChild(user);
        }
        user_num.innerHTML = name_list.length;
        user_list.scrollTop = user_list.scrollHeight;
        var change = type == 'login' ? '上线' : '下线';
        var data = '系统消息: ' + user_name + ' 已' + change;
        listMsg(data);
    }
    function sendMsg(msg) {
        var data = JSON.stringify(msg);
        ws.send(data);
    }
    function uuid(len, radix) {
        var chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');
        var uuid = [], i;
        radix = radix || chars.length;
        if (len) {
            for (i = 0; i < len; i++) uuid[i] = chars[0 | Math.random() * radix];
        } else {
            var r;
            uuid[8] = uuid[13] = uuid[18] = uuid[23] = '-';
            uuid[14] = '4';
            for (i = 0; i < 36; i++) {
                if (!uuid[i]) {
                    r = 0 | Math.random() * 16;
                    uuid[i] = chars[(i == 19) ? (r & 0x3) | 0x8 : r];
                }
            }
        }
        return uuid.join('');
    }
</script>
```

`server.go 入口文件`

```go
package main

import "github.com/gorilla/mux"
import "fmt"
import "net/http"

func main()  {
	// 创建路由 ~/Documents/Learn/goLearn/goProject/src/awesomeProject »
	router := mux.NewRouter()
	// ws控制器不断去处理管道数据，进行同步
	go h.run()
	// 指定ws 的回调函数
	router.HandleFunc("/ws", wsHandler)
	// 开启服务监听
	if err := http.ListenAndServe("127.0.0.1:8080", router); err!= nil{
		fmt.Println("err", err)
	}
}
```

`hub.go 连接器`

```go
package main

import (
	"encoding/json"
)

// 将连接器对象初始化
var h = hub{
	// connections 注册了连接器
	connections: make(map[*connection]bool),
	// 从连接器发送的信息
	broadcast: make(chan []byte),
	// 从连接器注册请求
	register: make(chan *connection),
	// 从连接器销毁请求
	unregister:make(chan *connection),
}

// 处理ws的逻辑实现
func (h *hub) run()  {
	// 监听数据管道，正在后端处理管道数据
	for {
		//跟进不同的数据管道，处理不同的逻辑
		select {
		// 注册
		case c:= <- h.register:
			// 标注注册
			h.connections[c] = true
			// 组装data数据
			c.data.Ip = c.ws.RemoteAddr().String()
			// 更改类型
			c.data.Type = "handshake"
			// 更新用户列表
			c.data.UserList = user_list
			data_b,_ := json.Marshal(c.data)
			// 将数据放入管道
			c.send <- data_b
		case c := <- h.unregister:
			// 注销
			if _,ok := h.connections[c]; ok{
				delete(h.connections, c)
				close(c.send)
			}
		case data := <- h.broadcast:
			// 处理数据流转，将数据同步到所有的用户,遍历所有的连接
			// c是单个连接
			for c:= range h.connections{
				// 将数据同步
				select {
				case c.send <- data:
				default:
					// 防止死循环
					delete(h.connections, c)
					close(c.send)
				}
			}
		default:

		}
	}
}
```

`data.go 传输的数据对象`

```go
package main

// 将链接中传输的数据抽象出对象
type Data struct {
	Ip string `json:"ip"`
	// 标识信息类型【login 登录、handshake 握手信息、system系统信息、logout 推出信息、user 普通信息】
	Type string `json:"type"`
	From string `json:"from"` // 代表哪一个用户说的
	Content string `json:"content"` // 传输内容
	User string `json:"user"` //
	UserList []string `json:"user_list"` // 用户列表
}
```

`connection.go 连接器`

```go
package main
// 抽象数据结构
// ws 的链接器 数据 管道
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)
type connection struct {
	// ws的连接器
	ws *websocket.Conn
	// 管道 byte的切片，byte可以
	send chan []byte
	// 数据
	data *Data
}

// 抽象一个ws的连接器、处理ws中的各种逻辑
type hub struct {
	// connections 注册了连接器
	connections map[*connection]bool
	// 从连接器发送的信息
	broadcast chan []byte
	// 从连接器注册请求
	register chan *connection
	// 从连接器销毁请求
	unregister chan *connection
}

// ws的读和写
// 1 往ws中写数据
func (c *connection)writeToWs()  {
	// 从管道遍历数据
	for message := range c.send{
		// 数据写出
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

var user_list = []string{}

// ws链接中读取数据
func (c *connection)reader()  {
	// 不断读取socket数据
	for  {
		message, p, err := c.ws.ReadMessage()
		fmt.Println(message)
		if err != nil{
			fmt.Println(err)
			// 读取不到数据
			h.unregister <- c
			break
		}
		// 读取数据 反序列化到对象
		json.Unmarshal(p, &c.data)
		// 跟进data中的type判断应该进行什么操作
		switch c.data.Type {
		case "login":
			// 弹出窗口，输入用户名
			c.data.User = c.data.Content
			c.data.From = c.data.User
			// 登录后将用户加入用户列表
			user_list = append(user_list, c.data.User)
			// 每一个登录的用户都要看到所有已经登录的用户
			c.data.UserList = user_list
			// 数据序列化
			data_b, err := json.Marshal(c.data)
			if err != nil{
				fmt.Println(err)
			}
			h.broadcast <- data_b
		// 普通用户
		case "user":
			c.data.Type = "user"
			data_b, err := json.Marshal(c.data)
			if err != nil{
				fmt.Println(err)
			}
			h.broadcast <- data_b
		case "logot":
			c.data.Type = "user"
			// 用户列表删除
			user_list = removeUser(user_list, c.data.User)
			c.data.UserList = user_list
			c.data.Content = c.data.User
			// 数据序列化，让所有人知道xx下线了
			data_b, err := json.Marshal(c.data)
			if err != nil{
				fmt.Println(err)
			}
			h.broadcast <- data_b
			h.unregister <- c
		default:
			fmt.Println("其他")
		}
	}
}
// 删除用户 删除用户切片中的数据
func removeUser(slice []string, user string) []string {
	// 严谨判断
	count := len(slice)
	if count == 0{
		return slice
	}
	if count == 1 && slice[0] == user{
		return []string{}
	}
	// 定义新的返回切片
	var my_slice = []string{}
	// 删除传入切片中的指定用户，其他的用户放到新的切片中
	for i := range slice{
		// 利用索引删除用户
		if slice[i] == user && i == count{
			return slice[:count]
		}else if slice[i] == user{
			my_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return my_slice
}

// 定义一个升级器 将http 升级为ws请求
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ws的回调函数
// http.ResponseWriter 响应
// *http.Request 请求
func wsHandler(w http.ResponseWriter, r *http.Request)  {
	ws,err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		fmt.Println(err)
	}
	// 创建链接对象 去做事情
	// 初始化连接对象
	c := &connection{send: make(chan []byte, 128), ws:ws, data: &Data{}}
	// 在ws中注册一下
	h.register <- c
	// ws 将数据读写跑起来
	go c.writeToWs()
	c.reader()

	//当主函数执行完毕，执行
	defer func() {
		c.data.Type = "logout"
		// 用户列表删除
		user_list = removeUser(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		// 数据序列化，让所有人知道xx下线了
		data_b, err := json.Marshal(c.data)
		if err != nil{
			fmt.Println(err)
		}
		h.broadcast <- data_b
		h.unregister <- c
		//
	}()

}
```

