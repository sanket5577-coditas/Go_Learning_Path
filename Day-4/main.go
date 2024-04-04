package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var sum int
	var input int

	fmt.Println("Enter the number of elements you want to use:")
	fmt.Scan(&input)

	numbers := make([]int, input)

	fmt.Println("Enter the elements for the list:")
	for i := 0; i < input; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	chunkSize := (len(numbers) + 2 - 1) / 2
	ch := make(chan int)

	for i := 0; i < len(numbers); i += chunkSize {
		wg.Add(1)
		go elementChan(numbers[i:i+chunkSize], &wg, ch, &sync.Mutex{})
		go addition(ch, &wg, &sync.Mutex{}, &sum)
		wg.Wait()
	}
	wg.Wait()
	fmt.Println("Sum is:", sum)
}

func elementChan(element []int, wg *sync.WaitGroup, ch chan<- int, mutex *sync.Mutex) {
	defer wg.Done()
	for _, v := range element {
		mutex.Lock()
		ch <- v * v
		mutex.Unlock()
	}
}

func addition(ch <-chan int, wg *sync.WaitGroup, mutex *sync.Mutex, sum *int) {
	defer wg.Done()
	for v := range ch {
		mutex.Lock()
		*sum += v
		mutex.Unlock()
	}
}
