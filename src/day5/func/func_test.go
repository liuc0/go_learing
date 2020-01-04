package fn_test

import (
	"math/rand"
	"time"
	"fmt"
	"testing"
)

/**
 * go语言函数使用
 */
func returnMultiValues() (int, int) {
	// 返回随机数
	return rand.Intn(10), rand.Intn(20)
}

// 输入是函数类型，返回也是函数类型，用于计算时间
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	// _表示忽略其中一个返回值
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 计算数组之和
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

// panic是中断程序，抛出
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
