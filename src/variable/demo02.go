package main

import "fmt"

// 指定变量类型，如果没有初始化，则变量默认为零值。
func main()  {
	// 声明一个变量并初始化
	var a = "Hello"
	fmt.Println(a)
	// 如果没有初始化则为0
	var b int
	fmt.Println(b)
	// bool 零值为 false
	var c bool
	fmt.Println(c)
	//

}
