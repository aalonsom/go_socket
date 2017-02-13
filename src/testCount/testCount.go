package main

import (
	//counter "testCount/counterSync"
	counter "testCount/counter"
	"fmt"
	"time"
)

func increment (nTimes int) {
	i := 0
	for i = 0; i < nTimes; i++ {
		counter.Add()
	}
	//fmt.Println(" ", i, " ")
}

func main() {

	i := 0
	nTimes   := 1000000
	nThreads := 10
	initial  := time.Now()
	initialD := time.Since(initial)

	for i = 0; i < nThreads; i++ {
		go increment(nTimes)
	}

	//end  := time.Now()

	time.Sleep(5e9)

	endD     := time.Since(initial)

	//fmt.Println(int64(time.Duration(endD)))
	//fmt.Println(int64(time.Duration(initialD)))
	result := int64(time.Duration(endD)) - int64(time.Duration(initialD))
	fmt.Print("Total count: ", counter.Get())
	fmt.Print("; Should be: ", nTimes * nThreads)
	fmt.Println("; Duration: ", result / int64(time.Millisecond), "ms")
}
