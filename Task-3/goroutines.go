package main

import (
	"fmt"
	
)
func main(){
	go greet("hello")
	greet("world")
	
}

func greet(s string){
	for i := 0; i <5; i++ {	
		fmt.Println(s)	
	}

}