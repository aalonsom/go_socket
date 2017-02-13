package main

import (
    "fmt"

    "user/counter"
    "user/stringutil"
)

func main() {
    fmt.Printf(stringutil.Reverse("!oG ,olleH"))
    fmt.Println(counter.Get())
}
