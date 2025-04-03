package main

import "fmt"
 
 	func main() {
		// Slice
		var names = []string{"Almofadinha", "Farofinha", "Bugiu", "Caruzo", "Salsicha"}
 		fmt.Println(names)
		rangeOne := names[:2]
		fmt.Println(rangeOne)
		rangeTwo := names[3:]
		fmt.Println(rangeTwo)
		rangeThree := names[2]
		fmt.Println(rangeThree)
 		
 	}
	