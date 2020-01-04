package _interface

import "testing"

/**
 * go语言的接口与依赖
 * 接口是非入侵性，实现不依赖于接口定义
 * 所以接口的定义可以包含在接口使用者包内
 */

 // 接口定义
type Programmer interface {
	WriteHelloWorld() string
}

// 接口实现
type GoProgrammer struct {
}
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
