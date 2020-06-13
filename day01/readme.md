# Golang 第一天

## Go 发展历史

1. 诞生历史

诞生于 2006 年1月2号，2009年发布并正式开源，2012 年第一个正式版Go 1.9 发布

## 环境搭建

这里使用 mac 自带安装工具 brew 安装：

> brew install go

配置 golong 的相关环境

> vim ~/.zshrc (或者 vim ~/.bashrc)

将下面内容添加到上面的文件中：

```shell
#GOROOT
export GOROOT=/usr/local/opt/go\@1.14

#GOPATH
export GOPATH=$HOME/code/golang

#GOPATH root bin
export PATH=$PATH:$GOROOT/bin
```

> GOPATH 可以根据个人习惯设置为其他目录

让刚刚的配置生效

> source ~/.zshrc (或者 source ~/.bashrc)

测试 golang 是否安装成功

> go env

```shell
~ go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
...
```

如果出现以上内容则安装成功

## Hello world

1. 和 python 一样, 把相同功能代码放到一个目录，称为包。包可以被其他包引用。



2. main 包是用来生成可执行文件，每个程序只有一个 main 包，包的主要用途是提高代码的可复用性，下面 hello.go 文件

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

> 终端运行：go run hello.go

3. 基本命令

- go run 快速执行 go 文件，就像执行脚本一样
- go build 编译程序，生成二进制可执行文件
- go install 安装可执行文件到 bin  目录
- go test 执行单元测试或压力测试
- go env 显示 go 相关的环境变量
- go fmt 格式化源代码



4. Go 程序结构

- go 源码按 package 进行组织，并且 package 要放到非注释的第一行
- 有个程序只有一个 main 包和一个 main 函数
- main 函数是程序的执行入口



5. Go 代码注释的方式

- 单行注释：// 代码 x x x x
- 多行注视：/*  x x x x  */



## Golang 语言特性

1. 垃圾回收

- 内存自动回收，再也不需要开发人员管理内存
- 开发人员专注业务实现，降低啦心智负担，可以更专注于代码部分
- 只需要 new 分配内存，不需要释放



2. 天然并发

- 从语言层面支持并发，只需要 go 一下
- goroutine，轻量级线程，创建成千上万个 goroite 成为可能

Go 串行实例：

```go
package main

import (
	"fmt"
	"time"
)

func calc()  {
	for i:=0; i < 10; i++ {
		time.Sleep(1*time.Second)
		fmt.Println( "current time", i)
	}
}

func main() {
	calc()
	fmt.Println("exited")
}
```

运行结果：

```shell
day01 git:(master) ✗ go run hello_world.go
current time 0
current time 1
current time 2
current time 3
current time 4
current time 5
current time 6
current time 7
current time 8
current time 9
exited
```



在耗时任务中加入 `go 函数名` 实现并发的效果

```
package main

import (
	"fmt"
	"time"
)

func calc()  {
	for i:=0; i < 10; i++ {
		time.Sleep(1*time.Second)
		fmt.Println( "current time", i)
	}
}

func main() {
	// 这里使用 go 实现并发
	go calc()
	fmt.Println("exited")
}

```

运行结果：

```
day01 git:(master) ✗ go run hello_world.go
exited
```

这个结果是因为主线程任务在打印出 `exited` 就已经执行完毕，子线程 `go calc()` 还未执行，所以为输出，有兴趣的可以在 `fmt.Println("exited")` 这行中加入 time.Sleep(11*time.Second) 可以看到不一样的效果



3. Channel(管道)

- 管道，类似 unix/linux 中的 pipe
- 多个 goroute 之间通过 channel 进行通信
- 支持任何类型



4. 多返回值

- 一个函数可以返回多个值



5. 编译语言

- 性能只比 C 语言差 10%
- 开发效率和 python php 差不多



## Go 标识符和关键字

1. 标志符
   - 标志符是用来表示GO中的变量名或函数名，以字母或_开头。后面跟着字母或数字
2. 关键字
   - 关键字是GO语言预先定义好的，有特许含义的标志符。eg: func、default、package、main ...



## GO 变量

1. Go语言中的函数名、变量名、常量名、类型名、语句标号和包名等所有的命名，都遵循一个简单的命名规则：一个名字必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线。大写字母和小写字母是不同的：heapSort和Heapsort是两个不同的名字。

2. 从根本上说，变量相当于是对一块数据存储空间的命名，程序可以通过定义一个变量来申请一块数据存储空间，之后可以通过引用变量名来使用这块存储空间。
3. 语法：var identifier type

```go
var a int = 1
var b string = "hello"
var c bool = true
```

如果一次申明多个变量可以使用如下写法：

```go
var (
	a int = 1
	b string = "hello"
	c bool = true
)
```

> 注意如果只申明变量未给变量赋值 int 类型的变量会默认值为 0, string 类型的为空字符串，bool 类型默认值为 false



## GO 常量

1. 常量使用 const 修饰，代表永远是只读的，不能修改
2. 语法：const identifier [type] = value，其中 type 可以省略。

```go
const b string = "hello world" 等价于 const b = "hello world"
const num int = 1 等价于 const num = 1
```

如果一次申明多个常量：

```
const (
	a int = 100
	b string = "hello"
)
```

3. Iota 枚举

常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，**但是不用每行都写一遍初始化表达式。在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一**

```
const (
    x = iota // x == 0
    y = iota // y == 1
    z = iota // z == 2
    w        // 这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
    a       = iota //a=0
    b       = "B"
    c       = iota             //c=2
    d, e, f = iota, iota, iota //d=3,e=3,f=3
    g       = iota             //g = 4
)
```



## GO 数据类型

1. 布尔类型

   - var bool bool 和 var b bool = true 和 var b = false
   - 操作符 == 和 !=
   - 取反操作符：!b
   - &&(且)  ||(或) 操作符
   - bool 类型格式化输出占位符：%t

   

2. 整数和浮点数类型

   - Int8、int16、int32、int64
   - uint8、uint16、uint32、uint64(无符号的)
   - int 和 unit 和操作系统平台有关
   - float32 和 float64 浮点类型
   - 所有整数初始化为0，所有浮点数初始化为0.0，布尔类型初始化为 false
   - Go 是强类型语言，不同类型相加以及赋值是不允许的
   - 输出占位符：整数 %d，%x十六进制，%f浮点数

   

3. 字符串类型

   - var str string
   - var str string = "hello"
   - 字符串输出占位符%s
   - 万能输出占位符：%v

- 字符串常用操作
  - 长度：len(str)
  - 拼接：+， fmt.Sprintf
  - 分割：strings.Split
  - 包含：strings.Contains
  - 前缀或后缀判断：strings.HasPrefix，strings.HasSuffix
  - 字符串出现的位置：string.index()，strings.LastIndex()
  - join 操作：string.Join(a[]string, sep, string)

````
package main

import {
	"fmt",
	"strings"
}

func main()  {
	// len 的使用
	var a string = "hello"
	var strLen = len(a)
	fmt.Printf("str len of %v", strLen)

	// 字符串拼接
	var b string = "world"
	var c string = a + b
	fmt.Println("\n", c)

	// 方式二： 使用 fmt.Sprintf()
	var d = fmt.Sprintf("%s%s", a, b)
	fmt.Println(d)

	// 字符串分割 strings.Split()
	var ipStr = "127.0.0.1;127.0.0.0"
	var ipArray = strings.Split(ipStr, ";")
	fmt.Printf("str splic %v \n", ipArray)

	// 字符串包含
	var result = "hello world"
	fmt.Println(strings.Contains(result, "hello"))

	// 字符串前后缀判断 strings.HasPrefix/strings.HasSuffix
	var longStr = "long str"
	fmt.Println(strings.HasPrefix(longStr, "long"))
	fmt.Println(strings.HasSuffix(longStr, "str"))

	// 字符串索引 strings.Index()/string.LastIndex
	//strings.Index()/strings.LastIndex(): 找到目标字符串则返回目标字符串索引，如果找不到则返回 -1
	var indexStr = "index string"
	var index = strings.Index(indexStr, "i")
	fmt.Println(index)
	var lastIndex = strings.LastIndex(indexStr, "i")
	fmt.Println(lastIndex)

	// 字符串 join 操作
	var strList []string = []string{"h", "e", "l", "l", "o"}
	fmt.Println(strings.Join(strList, ";"))
}
````



## GO 操作符

1. 操作符
   - 逻辑操作符：== 、!=、<、<=、>=
   - 算数操作符：+、-、*、/、%



## GO 程序的基本机构

- 任何一个代码文件都是隶属于一个包

- import 关键字，引用其他包

- 开发可执行程序，package main 并且有且只有一个 main 入口函数

- 包中函数调用

  - 同一个包中函数直接用函数名调用
  - 不同包中函数通过报名+点+函数名进行调用

- 包访问控制规则：

  - 大写意味着这个函数/变量是可导出
  - 小写意味着这个函数/变量是私有的，包外部不能访问

  