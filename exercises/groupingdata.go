package main

import "fmt"

// "Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"
func main() {
	x := map[string][]string{
		"bond_james":      []string{"shaken, not stirred", "martinis", "Women"},
		"moneypenny_miss": []string{"james bond", "literature", "computer sciense"},
		"no_dor":          []string{"being evil", "ice cream", "sunsets"},
	}

	x["john wick"] = []string{"cars", "his wife", "his dog"}
	delete(x, "no_dor")

	for key, value := range x {
		fmt.Println("this is ", key)
		for i, value2 := range value {
			fmt.Println("\t", i, value2)
		}
	}
}
