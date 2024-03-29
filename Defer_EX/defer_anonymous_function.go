package main

import "fmt"

func main1() {
    defer fmt.Println("Deferred in main")
    fmt.Println("Executing main")

    defer func() {
        defer fmt.Println("Deferred inside anonymous function")
        fmt.Println("Executing anonymous function")
    }()

    oo()
}

func oo() {
    defer fmt.Println("Deferred in foo")
    fmt.Println("Executing foo")
}
