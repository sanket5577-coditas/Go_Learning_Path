package main

import (
	"fmt"

)

func squareWorker(list []int, element chan int){

	for _, v := range list {
		element <- v * v
		fmt.Println(element)
		
		
	}
	

}
