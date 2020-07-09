package funcs

// 包内小写字母开头的为私有变量，包外无法使用
var name = "private name"
var Age = "public age"

func Sum(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}