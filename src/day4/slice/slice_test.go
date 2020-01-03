package slice

import "testing"

/**
 * Go 语言切片是对数组的抽象。
 */
func TestSliceInit(t *testing.T) {
	// 声明一个未指定大小的数组来定义切片
	var s0 []int
	t.Log(len(s0), cap(s0))
	// 在切片中添加数据1
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	// 直接初始化切片，[]表示是切片类型，{1,2,3,4}初始化值依次是1,2,3,4.其cap=len=4
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// 通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片,3为长度len,5为容量cap
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}
/**
 * 遍历切片
 */
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	// 将year中从下标startIndex(3)到endIndex-1(6) 下的元素创建为一个新的切片
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// if a == b { //切片只能和nil比较
	// 	t.Log("equal")
	// }
	t.Log(a, b)
}
