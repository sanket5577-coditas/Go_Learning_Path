package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	
	element := make(chan int)
	list := []int{2, 4, 6, 8, 10}

	wg.Add(2)

	go func(ch chan int, list []int, wg *sync.WaitGroup) {
		defer wg.Done()
		for _, v := range list {
			ch <- v * v
		}
		close(ch) 
	}(element, list, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		sum :=0
		for val := range ch {
			sum+=val
		}
		fmt.Println("The Sum of list is :",sum)
	}(element, wg)
	wg.Wait()
}
