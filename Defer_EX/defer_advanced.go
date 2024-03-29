package main

import "fmt"

func main() {
    defer fmt.Println("Deferred in main")
    fmt.Println("Executing main")
    
    defer func() {
        fmt.Println("Deferred function 1")
    }()
    
    defer func() {
        fmt.Println("Deferred function 2")
    }()
    
    fo()
}

func fo() {
    defer fmt.Println("Deferred in foo")
    fmt.Println("Executing foo")
}
