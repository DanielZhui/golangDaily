# Golang 接口

## 一、接口介绍

### 1. 显示生活中的接口

现实生活中手机、相机、U盘都可以和电脑的 USB 建立连接，我们不需要关注 USB 卡槽大小是否是一样，因为所有的 USB 接口都是按照统一的标准来设计的。

那在我们编程中所说的接口是什么？通俗的讲接口就是一个标准，它是对一个对象的行为和规范进行约定，约定实现接口的对象必须得按照接口的规范。



### 2.Golang 接口定义

​	在 Golang 中接口（interface）是一种类型，一种抽象类型。接口（interface）是一组函数 method 的集合，Golang 中的接口不能包含任何其他变量。

​	在 Golang 中接口中的所有方法都没有方法体，接口定义了一个对象行为规范，只定义规范不实现。接口体现了程序设计的多态和内聚低耦合的思想。

​	Golang 中的接口也是一种数据类型，不需要显示实现。只需要一个变量含有接口类型中的所有方法，那么这个变量就实现了这个接口。

​	Golang 中每个接口有数个方法组成，接口的定义格式如下：

```
type 接口名 interface {
	方法名1(参数列表1) 返回值列表 1
	方法名2(参数列表2) 返回值列表 2
	...
}
```

其中：

- 接口名：使用 type 讲接口定义为自定义的类型名。Go 语言的接口在命名时，一般会在单词后面添加 er,如有写操作的接口叫 Write,有字符串功能的接口叫 Stringer 等。接口名最好要突出该接口的类型含义。
- 方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。



```go
package main

import "fmt"

type Usber interface {
	start()
	stop()
}

// 如果接口里面有方法必须要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	Name string
}

func (p Phone)  run() {
	fmt.Println(p.Name, "phone is running...")
}

func (p Phone)  start() {
	fmt.Println(p.Name, "phone is start...")
}

func (p Phone) stop()  {
	fmt.Println(p.Name, "phone stop ...")
}

type Camera struct {
	Name string
}

func (c Camera) start()  {
	fmt.Println(c.Name, "Camera is busy")
}

func (c Camera) stop()  {
	fmt.Println(c.Name, "Camera stop")
}

type Computer struct {
	Name string
}

func (c Computer) work(usb Usber)  {
	usb.start()
	usb.stop()
}

func main()  {
	p := Phone {
		Name: "iphone",
	}
	// p为定义接口, 这里不会报错
	p.run()
	p.stop()

	// p1 定义 Usber 接口类型，需要实现 start() stop() 方法
	var p1 Usber

	p1 = p
	// 由于 Usber 中没有定义 run(),这里会报错
	//p1.run()

	p1.start()

	c := Camera{
		"Nikon",
	}
	// c 定义为 Usber 接口类型
	var c1 Usber = c
	c1.start()
	c1.stop()

	mac := Computer{
		Name:"mac",
	}

	iphone := Phone{
		Name: "iphone",
	}

	camera := Camera{
		Name: "Canon",
	}

	mac.work(iphone)
	mac.work(camera)
}
```



### 3. 空接口

Golang 中接口可以不定义任何方法，没有任何方法的接口就是空接口。空接口没有任何约束，因此任何类型变量都可以实现空接口。

空接口在实际项目中用的是非常多，用空接口可以表示任意数据类型。



- 空接口作为函数的参数

使用空接口实现可以接收任意类型的函数参数。

```
// 空接口作为函数参数，可接受任何类型参数
func show(args interface{}) {
	fmt.Printf("值：%v 类型：%T\n", args, args)
}
```



- map 的值实现空接口

使用空接口实现可以保存任意值的字典。

```
// 空接口作为 map 值
var studentInfo = make(map[string]interface{})
studentInfo['name'] = 'tom'
studentInfo['age'] = 18
studentInfo['male'] = true
fmt.Println(studentInfo)
```



- 切片实现空接口

```
var slice = []interface{}["tom", 20, true, 1.66]
fmt.Println(slice)
```

```
package main

import "fmt"

type Usber interface {
	start()
	stop()
}

// 如果接口里面有方法必须要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	Name string
}

func (p Phone)  run() {
	fmt.Println(p.Name, "phone is running...")
}

func (p Phone)  start() {
	fmt.Println(p.Name, "phone is start...")
}

func (p Phone) stop()  {
	fmt.Println(p.Name, "phone stop ...")
}

type Camera struct {
	Name string
}

func (c Camera) start()  {
	fmt.Println(c.Name, "Camera is busy")
}

func (c Camera) stop()  {
	fmt.Println(c.Name, "Camera stop")
}

type Computer struct {
	Name string
}

func (c Computer) work(usb Usber)  {
	usb.start()
	usb.stop()
}

// 空接口
type None interface {

}

// 空接口作为函数参数
func show(args interface{})  {
	fmt.Printf("值：%v 类型：%T\n", args, args)
}

func main()  {
	p := Phone {
		Name: "iphone",
	}
	// p为定义接口, 这里不会报错
	p.run()
	p.stop()

	// p1 定义 Usber 接口类型，需要实现 start() stop() 方法
	var p1 Usber

	p1 = p
	// 由于 Usber 中没有定义 run(),这里会报错
	//p1.run()

	p1.start()

	c := Camera{
		"Nikon",
	}
	// c 定义为 Usber 接口类型
	var c1 Usber = c
	c1.start()
	c1.stop()

	mac := Computer{
		Name:"mac",
	}

	iphone := Phone{
		Name: "iphone",
	}

	camera := Camera{
		Name: "Canon",
	}

	mac.work(iphone)
	mac.work(camera)

	var none None

	var str = "hello golang"
	none = str
	fmt.Printf("值：%v 类型：%T\n", none, none)

	show("hello world")
	show(2020)
	show(false)

	// map 值为空接口实现 value 为任意值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "tom"
	studentInfo["age"] = 18
	studentInfo["male"] = true
	fmt.Println(studentInfo)

	// 切片为任意值实现切片里的值可为任意值
	var slice = []interface{}{"tom", 20, true, 1.66}
	fmt.Println(slice)
}

```