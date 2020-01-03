package main

import "fmt"

// 指定变量类型，如果没有初始化，则变量默认为零值。
func main()  {
	var a int
	var b string
	var c float64
	var d bool
	fmt.Printf("%v %q %v %v \n",a,b,c,d )
}