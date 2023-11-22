package main

import (
	"fmt"
)

func main() {
	s := 0
	for i := 0; i < 100; i++ {
		s += i
	}
	fmt.Println(s)
}

//export sum
func sum(a int, b int) int {
	return a + b
}

//export fibArray
func fibArray(n int32) int32 {
	arr := make([]int32, n)
	for i := int32(0); i < n; i++ {
		switch {
		case i < 2:
			arr[i] = i
		default:
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return arr[n-1]
}

//export fib
func fib(n int32) int32 {
	switch {
	case n < 2:
		return n
	default:
		return fib(n-1) + fib(n-2)
	}
}
