## very beginning

```go
装好go安装包，开始要大展拳脚的时候，被TM的go环境恶心的想放弃。
go module 是在go1.13版本之后的
```

## 1 环境变量GO111MODULE

```go
// GO111MODULE 一共三种模式
1 off 那么go命令行将不会使用新的module功能，相反的，它将会在vendor目录下和GOPATH目录中查找依赖包。也把这种模式叫GOPATH模式。
2 on 那么go命令行就会使用modules功能，而不会访问GOPATH。也把这种模式称作module-aware模式，这种模式下，GOPATH不再在build时扮演导入的角色，但是尽管如此，它还是承担着存储下载依赖包的角色。它会将依赖包放在GOPATH/pkg/mod目录下。
3 auto 这种模式是默认的模式，也就是说在你不设置的情况下，就是auto。这种情况下，go命令行会根据当前目录来决定是否启用module功能。只有当当前目录在GOPATH/src目录之外而且当前目录包含go.mod文件或者其子目录包含go.mod文件才会启用。
```

## 2 设置代理【加速库的下载】

```go
 ~$ go env -w GOPROXY=https://goproxy.cn,direct // 临时生效，重启失效
```

## 3 初始化module

```go
go mod init xxx // xxx为项目名
// 进入到项目目录进行init 初始化设置
```

## 4 检查项目依赖

```go
go mod tidy
// tidy会检测该文件夹目录下所有引入的依赖,写入 go.mod 文件
```

## 5 下载依赖

```go
go mod download
```

## 6 导入依赖

```go
go mod vendor
执行此命令,会将刚才下载至 GOPATH 下的依赖转移至该项目根目录下的 vendor(自动新建) 文件夹下
```

## 7 goland 中开启 go module

```go
settings -> go -> go module(vgo)
```

## 8 在写作中使用go module

```go
要注意的是, 在项目管理中,如使用git,请将 vendor 文件夹放入白名单,不然项目中带上包体积会很大
git设置白名单方式为在git托管的项目根目录新建 .gitignore 文件
设置忽略即可.

但是 go.mod 和 go.sum 不要忽略
另一人clone项目后在本地进行依赖更新(同上方依赖更新)即可
```

## 9 go module的常用命令

```go
go mod init  # 初始化go.mod
go mod tidy  # 更新依赖文件
go mod download  # 下载依赖文件
go mod vendor  # 将依赖转移至本地的vendor文件
go mod edit  # 手动修改依赖文件
go mod graph  # 打印依赖图
go mod verify  # 校验依赖
```



