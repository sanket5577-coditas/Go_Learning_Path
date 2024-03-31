package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var mut sync.Mutex

	squaredChan := make(chan int)
	numbers := []int{2, 4, 6, 8, 10}
	var sum int

	wg.Add(2)

	go func(ch chan int, nums []int, wg *sync.WaitGroup) {
		defer wg.Done()
		for _, num := range nums {
			ch <- num * num
		}
		close(ch)
	}(squaredChan, numbers, &wg)

	go func(ch chan int, wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		// sum := 0
		for squaredNum := range ch {
			fmt.Println("Squared element:", squaredNum)
			mut.Lock()
			sum += squaredNum
			mut.Unlock()
		}
		fmt.Println("Sum of squared numbers:", sum)
	}(squaredChan, &wg, &mut)

	wg.Wait()
}
