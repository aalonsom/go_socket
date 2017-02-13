package main

import (
	"fmt"
	"counter"
)

func main () {

	//counter.init()
	counter.Add()
	fmt.Println(counter.Get())
}
