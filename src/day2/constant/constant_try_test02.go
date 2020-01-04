package constant

import "unsafe"

/**
 * 定义枚举常量
 */
const (
	a = "Hello"
	b = 12
	c = unsafe.Sizeof(a)
)

func main()  {
	println(a,b,c)

}