package main

import (
	"fmt"
	"day02/funcs"
	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("hello world")
	sum := funcs.Sum(7, 3)
	fmt.Println(sum)

	// 由于 funcs sub 是小写属于私有变量只能在 funcs 包内使用
	//sub := funcs.sub(7, 3)
	//fmt.Println(sub)

	fmt.Println(funcs.Age)

	fmt.Println(funcs.Gender)

	// 第三方模块
	quantity := decimal.NewFromInt(3)
	fmt.Println(quantity)

	fee, err := decimal.NewFromString("136.02")
	fmt.Println(fee, err)
}