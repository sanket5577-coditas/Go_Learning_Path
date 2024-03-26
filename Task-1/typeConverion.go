package main

import (
	"fmt"
	"math"
)

func typeconversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	// const:= 2  can't do this 

	const number = 2   
	fmt.Println(number,x, y, z)
}
