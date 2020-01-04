package main

import (
	"os"
	"fmt"
)

/**
 * 安装环境后，go语言的第一个程序
 */
func main()  {
	if len(os.Args) > 1 {
		fmt.Print("hello world",os.Args[1])
	}
}
