## GIN

### 1 安装

```go
go get -u github.com/gin-gonic/gin
```

### 2 简单的gin框架

```go
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}

func main()  {
	// 1 创建路由
	router := gin.Default()
	// 2 绑定路由函数
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	// 3 执行
	err := router.Run(":8000")
	HandleError(err, "gin start")
}
// 浏览器输入 127.0.0.1:8000
```

### 3 自定义server创建

```go
func main() {
	// 1 创建路由
	router := gin.Default()
	s := &http.Server{
		Addr:         ":8000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// 2 绑定路由函数
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	// 3 执行
	s.ListenAndServe()
}
```

## Gin的路由系统

`1 基本路由`

```go
// gin 框架中采用的路由库是基于httprouter做的
/* 路由设计
普通的
/xxx/getxx
/xxx/postxx
/xxx/updatexx
/xxx/delxx
restful
/xxx/user		GET	获取
/xxx/user		POST 增加
/xxx/user		UPDATE 更新
/xxx/user		DEL		删除
*/
1 api参数
func main() {
	router := gin.Default()
  // *表示 *后面可以是任意字符串
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		actioin := c.Param("action")
		c.String(http.StatusOK, name, actioin)
	})
	router.Run(":8000")
}
2 url参数
// 127.0.0.1:8000/user?name=alex
func main() {
	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "mosson") // 没有默认是mosson
		c.String(http.StatusOK, name)
	})
	router.Run(":8000")
}
```

`2 gin 接受表单参数`

`login.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<form action="http://127.0.0.1:8000/form" method="post" enctype="application/x-www-form-urlencoded">
    用户名：<input type="text" name="username"><br>
    密码：<input type="text" name="password"><br>
    兴趣：
    <input type="checkbox" name="hobby" value="run"> 跑步
    <input type="checkbox" name="hobby" value="game"> 游戏
    <input type="checkbox" name="hobby" value="run"> money
    <input type="submit" value="登陆">
</form>
</body>
</html>
```

`tt.go`

```go
func main() {
	router := gin.Default()
	router.POST("/form", func(c *gin.Context) {
		// 取默认值
		type1 := c.DefaultPostForm("xxx", "yyy")
		// 获取单个值
		username := c.PostForm("username")
		password := c.PostForm("password")
		// 取多选
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK, fmt.Sprintf("type1:%s username:%s password:%s" +
			"hobby:%v", type1, username, password, hobbys))
	})
	router.Run(":8000")
}
// 手动在文件中打开 html文件输入用户信息，然后提交
```

`3 gin上传单个文件`

`upload.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>file</title>
</head>
<body>
<form action="http://127.0.0.1:8000/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="file">
    <input type="submit" value="上传">
</form>
</body>
</html>
```

`upload.go`

```go
func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		//
		file, err := c.FormFile("file")
		HandleError(err, "gin 文件上传")
		// 存文件
		c.SaveUploadedFile(file, "/Users/mosson/Documents/学习/Go/goPro/"+file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("%s upload success", file.Filename))
	})
	router.Run(":8000")
}
```

`4 多个图片上传`

`multiupload.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>file</title>
</head>
<body>
<form action="http://127.0.0.1:8000/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="files" multiple>
    <input type="submit" value="上传">
</form>
</body>
</html>
```

`uploads.go`

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		//
		form, err := c.MultipartForm()
		HandleError(err, "gin 文件上传")
		// 获取所有的文件
		files := form.File["files"]
		for _, file := range files{
			if err := c.SaveUploadedFile(file, "/Users/mosson/Documents/学习/Go/goPro/"+file.Filename);err != nil{
				c.String(http.StatusBadRequest, fmt.Sprintf("%s upload faild", file.Filename))
			}
		}
		c.String(http.StatusOK, fmt.Sprintf(" upload success"))
	})
	router.Run(":8000")
}
```

`5 路由组`

`routerGoup.go`

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(err, why)
	}
}

func main() {
	router := gin.Default()
	// 定义路由组，处理Get请求
	v1 := router.Group("/v1")
	// 在路由组v1下，创建需要的方法
	{
		v1.GET("/login", login)
		v1.GET("/read", read)
	}
	v2 := router.Group("/v2")
	// 在路由组v1下，创建需要的方法
	{
		v2.POST("/login", login)
		v2.POST("/read", read)
	}
	router.Run(":8000")
}

func login(c *gin.Context)  {
	name := c.DefaultQuery("name", "zhangsan")
	c.String(http.StatusOK, fmt.Sprintf("你好%s\n", name))
}
func read(c *gin.Context)  {
	name := c.DefaultQuery("name", "lisi")
	c.String(http.StatusOK, fmt.Sprintf("你好%s\n", name))
}
// 测试访问
curl http://127.0.0.1:8000/v1/login
curl http://127.0.0.1:8000/v2/login -X POST
```

### gin的路由原理

```go
gin 是采用httprouter的路由规则【高效】
httprouter会将所有的路由规则，构造一颗前缀树
```





