package main

// import "fmt"

func main(){
	element := make(chan int)
	list:=[]int{2,4,6,8,10}
	// fmt.Println(squareWorker(list,element))
	squareWorker(list,element)
}