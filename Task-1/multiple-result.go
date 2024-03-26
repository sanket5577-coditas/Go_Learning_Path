package main

import "fmt"

// here we have created a function that accept two parameter both as string and return two string 
func swap(x, y string) (string, string) {
	return y, x
}

func multipleReturn() {
	
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
