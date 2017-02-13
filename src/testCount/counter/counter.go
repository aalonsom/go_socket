package counter

var value int = 0

func init() {
	value = 0
}

func Add() {
	value ++;
}

func Get() int {
	return value;
}

func Set (i int) {
	value = i
}


