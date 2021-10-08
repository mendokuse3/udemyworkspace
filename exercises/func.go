package main

import (
	"fmt"
)

func main() {
	p1 := person{
		name: "bob",
		age:  34,
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	p.name = "joey"
}
