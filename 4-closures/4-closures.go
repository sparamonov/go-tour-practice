package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	res, pr, nx := 0, 0, 1
		
	return func() int {
		res = pr
		pr, nx = nx, pr + nx
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
