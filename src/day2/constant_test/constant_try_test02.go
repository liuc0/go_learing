package constant

import "unsafe"

// 枚举常量
const (
	a = "Hello"
	b = 12
	c = unsafe.Sizeof(a)
)

func main()  {
	println(a,b,c)

}