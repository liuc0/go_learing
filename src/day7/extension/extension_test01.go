package extension

import (
	"fmt"
	"testing"
)

type User struct {
	name string
}

func (u *User) sayHi() {
	u.sayName()
	u.sayType()
}

func (u *User) sayName() {
	fmt.Printf("I am %s.", u.name)
}

func (u *User) sayType() {
	fmt.Println("I am a user.")
}

func TestDog01(t *testing.T) {
	user := &User{name: "Chris"}
	user.sayHi()
}