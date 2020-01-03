package main

import "fmt"

// 省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，
// 格式：v_name := value

func main()  {
	f := "Runoob" // 将 var f string = "Runoob" 简写为 f := "Runoob"

	fmt.Println(f)
}