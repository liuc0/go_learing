package constant

import "fmt"

// 常量的定义格式
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
