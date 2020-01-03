package array_test

import "testing"

/**
 * 数组的使用方式
 */
func TestArrayInit(t *testing.T) {
	// 初始化一个容量为3的数组
	var arr [3]int
	// 初始化一个容量为3，且指定数据的数组
	arr1 := [4]int{1, 2, 3, 4}
	// 初始化一个根据数据长度而定的数组
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}
/**
 * 遍历数组
 */
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	// 遍历方式一：普遍遍历
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	// 遍历方式二：普遍遍历
	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[:]
	t.Log(arr3_sec)
}
