package main

import "fmt"

// now here WE ARE TELLING GO THAT WHAT ACUTLLAY WE ARE GOING TO RETURN FROM THE FUNCTION AND BY DOING THIS WE DON'T NEED TO RETURN THETHINGS IT WILL AUTTOMATICALLY RETURN 
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func named() {
	x,y := split(17)
	fmt.Println(x,y)
}
