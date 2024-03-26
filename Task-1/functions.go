package main

import "fmt"

// the function will return and sum of two number whic type will be int and it take two parameter both as int only  
func add(x int, y int) int {
	return x + y
}

// ifwe have same datatype as input we can write as follows 
// func add(x , y int) int {
// 	return x + y
// }

func function() {
	// fuction call
	fmt.Println("The addition of two number is :",add(42, 13))
}
