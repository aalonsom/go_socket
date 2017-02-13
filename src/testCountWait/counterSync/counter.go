package counter

import "sync"

var value int = 0
var mutex sync.Mutex

func init() {
	value = 0
}

func Add() {
	mutex.Lock()
	value ++;
	mutex.Unlock()
}

func Get() int {
	return value;
}

func Set (i int) {
	value = i
}


