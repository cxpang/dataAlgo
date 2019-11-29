package main

import (
	"fmt"
)

func executeFn(fn func() int) int {
	return fn()
}

func main() {
	a := 1
	b := 2
	c := executeFn(func() int {
		a += b
		return a
	})
	fmt.Printf("%d %d %d\n", a, b, c)
	user:=NewUser("pangchenxu",23)
	fmt.Println(user)

}
func SayHello(name string)  {
	fmt.Println(name)
}

type User struct {
	name string
	age int
}

func NewUser(name string,age int) *User {
	var user = new(User)
	user.name = name
	user.age = age
	return user
}