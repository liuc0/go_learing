package main

// 类型相同多个变量, 非全局变量
var x, y int
// 这种因式分解关键字的写法一般用于声明全局变量
var (
	a int
	b bool
)

var c, d int = 1, 2
// 和 python 很像,不需要显示声明类型，自动推断
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现，否则会导致编译错误
//g, h := 123, "hello"

func main(){
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}