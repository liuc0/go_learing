package extension

import (
	"fmt"
	"testing"
)

/**
 * go语言的扩展与复用
 */
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Dog) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao")
}
