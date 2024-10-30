package main

import (
    "fmt"
    stack "stackpj/stack"
)

func main() {
    fmt.Println("Hello, World!");
    s := stack.NewStack()

    s.Push(1)
    s.Push(2)
    s.Push(3)
    
    for {
        val, err := s.Pop()
        if err != nil {
            break
        }
        fmt.Println(val)
    }
    fmt.Println("Done")
}
