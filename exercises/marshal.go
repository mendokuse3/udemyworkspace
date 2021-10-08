package main

import (
	"encoding/json"
	"fmt"
)

// type person struct {
// 	First string
// 	Last  string
// 	Age   int
// }

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"james","Last":"bond","Age":32},{"First":"miss","Last":"moneypenny","Age":27}]`
	bs := []byte(s)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\nperson number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	// p1 := person{
	// 	First: "james",
	// 	Last:  "bond",
	// 	Age:   32,
	// }
	// p2 := person{
	// 	First: "miss",
	// 	Last:  "moneypenny",
	// 	Age:   27,
	// }

	// people := []person{p1, p2}

	// fmt.Println(people)

	// bs, err := json.Marshal(people)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(bs))
}
