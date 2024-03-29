package main

import "fmt"

func foo() int {
    defer fmt.Println("Deferred in foo")
    defer fmt.Println("Another deferred in foo")
    fmt.Println("Executing foo")
    return 5
}

func bar() {
    defer fmt.Println("Deferred in bar")
    fmt.Println("Executing bar")
}

// func deferExample() {
//     defer fmt.Println("Deferred in main")
//     fmt.Println("Executing main")
//     result := foo()
//     fmt.Println("Result from foo:", result)
//     bar()
// }

