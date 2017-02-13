package main

import (
	//aCounter "testCount/counterSync"
	aCounter "testCount/counter"
	"fmt"
	"time"
)

func increment (nTimes int, out chan<- int ) {
	i := 0
	for i = 0; i < nTimes; i++ {
		aCounter.Add()
	}
	out <- 0
}

func main() {

	channel := make (chan int)

	i := 0
	nTimes   := 1000000
	nThreads := 10

	// get current time
	initial  := time.Now()
	initialD := time.Since(initial)

	// create coroutines
	for i = 0; i < nThreads; i++ {
		go increment(nTimes, channel)
	}

	// wait for finishing of coroutines
	for i = 0; i < nThreads; i ++{
		<-channel
	}

	// get current time
	endD     := time.Since(initial)
	result := int64(time.Duration(endD)) - int64(time.Duration(initialD))

	fmt.Print("Total count: ", aCounter.Get())
	fmt.Print("; Should be: ", nTimes * nThreads)
	fmt.Println("; Duration: ", result / int64(time.Millisecond), "ms")
}
