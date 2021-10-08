package main

import "fmt"

// type person struct {
// 	firstname string
// 	lastname  string
// 	flavors   []string
// }

func main() {
	// x := person{
	// 	firstname: "john",
	// 	lastname:  "wick",
	// 	flavors:   []string{"vanilla", "revenge", "coffee"},
	// }
	// y := person{
	// 	firstname: "baba",
	// 	lastname:  "yaga",
	// 	flavors:   []string{"blood", "death", "headshots"},
	// }
	// fmt.Println(x.firstname, x.lastname)
	// for _, value := range x.flavors {
	// 	fmt.Println("Favorite flavors are: \t", value)
	// }
	// fmt.Println(y.firstname, y.lastname)
	// for _, value := range y.flavors {
	// 	fmt.Println("Favorite flavors are: \t", value)
	// }

	// mp := map[string]person{
	// 	x.lastname: x,
	// 	y.lastname: y,
	// }

	// fmt.Println(mp)

	// for key, value := range mp {
	// 	fmt.Println(key)
	// 	fmt.Println(value.firstname)
	// 	fmt.Println(value.lastname)
	// 	for _, value2 := range value.flavors {
	// 		fmt.Println(value2)
	// 	}
	// }

	// type vehicle struct {
	// 	doors string
	// 	color string
	// }
	// type truck struct {
	// 	vehicle
	// 	fourWheel bool
	// }
	// type sedan struct {
	// 	vehicle
	// 	luxury bool
	// }
	// tacoma := truck{
	// 	vehicle: vehicle{
	// 		doors: "four door",
	// 		color: "white",
	// 	},
	// 	fourWheel: true,
	// }
	// tesla := sedan{
	// 	vehicle: vehicle{
	// 		doors: "four door",
	// 		color: "nardo grey",
	// 	},
	// 	luxury: true,
	// }

	// fmt.Println(tesla)
	// fmt.Println(tacoma.doors, tacoma.color, tacoma.fourWheel)
	// fmt.Println(tesla.doors, tesla.color, tesla.luxury)
	x := struct {
		indev  bool
		rating int
	}{
		indev:  true,
		rating: 5,
	}

	fmt.Println(x)
}
