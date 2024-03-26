package main

import "fmt"

func loops() {

	// the increment part is optional in go
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
