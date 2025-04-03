package main

import "fmt"
 
 	func main() {
		// Slice
		var number = []int{11, 22, 33, 44, 55}
 		fmt.Println(number)
 		number = append(number, 66, 77, 88)
 		fmt.Println(number, len(number), cap(number))
 	}
	