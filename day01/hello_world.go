package main

import (
	"fmt"
	"strings"
)

/*
hello world 实例
func main()  {
	fmt.Println("hello world")
}
 */


/*
并发实例
func calc()  {
	for i:=0; i < 10; i++ {
		time.Sleep(1*time.Second)
		fmt.Println( "current time", i)
	}
}


func main() {
	go calc()
	fmt.Println("hello world")
}
 */


/*
d多返回值
func add(a int, b int) (int, int) {
	return a + b, a - b
}

func main()  {
	sum, sub := add(3, 4)
	fmt.Println(sum, sub)
}
 */


/*
func main()  {
	var (
		a int = 1
		b string = "go"
		c bool = true
	)

	fmt.Printf("a=%d b=%s c=%t", a, b, c)
}
// 输出：a=1 b=go c=true
 */


/*
字符串操作
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
 */
