package condition

import "testing"

/**
 * switch主要用法
 */
func TestSwitchMultiCase(t *testing.T)  {
	// 初始化语句;条件判断;变量修改
	// switch判断整数类型
	for i := 0; i< 5 ; i++ {
		switch i {
		case 0,2:
			t.Log("Even")
		case 1,3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	// switch判断boolean类型
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}