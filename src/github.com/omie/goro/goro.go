package main

import (
    "fmt"
)

func f(id string, to int) {
    for i := 0; i<to; i++ {
        fmt.Println(id, ":", i)
    }
}
func main() {
    f("first", 5)

    go f("second", 30)

    go func() {
        for i := 0; i<10; i++ {
            fmt.Println("anon : ", i)
        }
    }()
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}

