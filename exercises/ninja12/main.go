package main

import (
	"fmt"
	"udemyWorkspace/exercises/ninja12/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	spot := canine{
		name: "spot",
		age:  dog.Years(12),
	}
	fmt.Println(spot)
}
