package constant

import "fmt"

/**
 * 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
 * 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
 */
func main()  {
	const LENGHT int = 10
	const WIDTH int = 5
	var area int
	const a,b,c  = 1,false,"Hello" // 多重赋值

	area = LENGHT * WIDTH
	fmt.Println("面积为：",area)

	fmt.Println()
	fmt.Println(a,b,c)

	
}
