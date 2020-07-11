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
