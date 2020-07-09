# Golang 第二天（golang 包详解）

## 一、Golang 中的包的介绍和定义

包（package）是多个 Go 源码的集合，是一种高级的代码复用方案，Go 语言为我们提供了很多内置包，如 fmt、strconv、string、sort、errors、time 等。

Golang 中的包可以分为三种：

- 系统内置包
- 自定义包
- 第三方包

系统内置包：Golang 语言给我们提供内置包，映入后可以直接使用，如 fmt、strconv、strings、sort 等

自定义包：开发者自己写的包

第三方包：属于自定义包的一种，需要下载安装到本地后才可以使用



## Golang 包管理工具 go mod

在 Golang1.11 版本之前如果我们要自定义包的话必须把项目放在 GOPATH 目录。Go1.11 版本之后无需手动配置环境变量，使用 go mod 管理项目，也不需要非得把项目放到 GOPATH 指定目录下，你可以在你磁盘的任何位置新建一个项目， Go1.13 以后可以彻底不要 GOPATH。

### go mod 使用

```
Usage:

        go mod <command> [arguments]

The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph
        init        initialize new module in current directory
        tidy        add missing and remove unused modules
        vendor      make vendored copy of dependencies
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed
```



## Golang 包的引用

### 1. 系统内置包

Go 语言为我们提供了很多内置包，如 fmt、strconv、string、sort、errors、time 等，可以直接引用。

```
// 引用单个包
import "fmt"

// 引用多个包
import (
	"fmt"
	"sort"
	"time"
)
```



### 2. 自定义包

- 在当前 day02 目录中新建 funcs 目录，下面新建 funcs.go 文件

  ```
  package funcs	// package 一定要放在最上面
  
  // 包内小写字母开头的为私有变量，包外无法使用
  var Gender = "public gender"
  ```

- 在 main.go 中引用 funcs 包

  ```
  package main
  
  import (
  	"fmt"
  	"day02/funcs"
  )
  
  func main() {
  	fmt.Println("hello world")
  
  	fmt.Println(funcs.Gender)
  }
  ```



### 3. 引用第三方包

第三方包以 decimal 为例

- 全局引用 go mod download packageName

  > go mod download github.com/shopspring/decimal



- 全局引用 go get packageName

  > go get github.com/shopspring/decimal



-  当前项目引用 go vendor

  1. 在当前项目中导入 decimal 包

     ```
     import (
        "fmt"
        "github.com/shopspring/decimal"
     )
     ```

     使用 go vendor 命令，当前项目下会生成一个 vendor 文件夹，vendor文件夹下 module.txt 存放项目第三方包信息，当我们使用 go get xxx 或者 go mod download xxx 第三方包时可能出现 `dial tcp 172.217.160.113:443: i/o timeout`, 解决方式如下所示。

     

### 解决 go 下载第三方包 timeout 问题

修改本地 go  proxy：https://shockerli.net/post/go-get-golang-org-x-solution/

从 Go1.11 版本开始，官方支持 `go module` 包依赖管理工具。其实还新增了 `GOPROXY` 环境变量。如果设置该环境变量，下载源码时将会通过这个环境变量设置的代理地址，而不再是以前的直接从代码库下载。

[goproxy.io](https://github.com/goproxyio/goproxy) 这个开源项目帮我们实现下载第三方依赖包加速。该项目允许开发者一键构建自己 `GOPROXY` 代理服务，同时也提供公用的代理服务 [https://goproxy.io，](https://goproxy.io，) 我们只需要设置该环境变量即可正常下载被墙的源码包：

```
export GOPATH=https://goproxy.io
```

不过，需要依赖于 `go module` 功能。可通过 `export GO111MUDULE=on` 开启 `MODULE`。

如果项目不在 `GOPATH` 中，则无法使用 `go get ...`，但可以使用 `go mod ...` 命令。也可以通过置空这个环境变量关闭， `export GOPATH=`

对于 Windows 用户，可以在 `PowerShell` 中设置：

```
$env:GOPATH = "<https://goproxy.io>"
```

最后我们当然推荐使用 `GOPATH` 这个环境变量的解决方式，前提是 `Go version ≥ 1.11`。