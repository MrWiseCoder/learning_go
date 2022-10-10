package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Hello go")
    var s string
    var v int
    var r = rand.Int() % 10
    for i:= 0; i < 3; i++ {
        fmt.Println(i)
        fmt.Print("Please enter a val: ")
        fmt.Scanln(&s)
        v, _ = strconv.Atoi(s)
        if v == r {
            fmt.Println(" \t\t *** OK *** ")
            break
        } else {
            fmt.Println(" \t\t --- suck --- ")
        }
    }
    fmt.Println("it was", r)
}
