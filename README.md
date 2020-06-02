# How-to-become-a-go-programmer

# 1 go get golang.org/x/text 问题

##### 首先在你的项目路径src里新建golang.org/x文件目录，如果有就不用了

##### golang.org/x目录下

##### `git clone https://github.com/golang/text.git`

##### 回到在src目录下

##### `go install -x golang.org/x/text`

##### 会在pkg目录下生成一个text.a的包文件，就成功

