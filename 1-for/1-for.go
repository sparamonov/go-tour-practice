package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	res := 1.0
	prevRes := 2.0
	
	for math.Abs(prevRes - res) > 0.000000001 {
	//for i := 0; i < 10; i++ {
		prevRes = res
		res = res - ((math.Pow(res, 2) - x) / (2 * res))
		
		fmt.Printf("res = %g\nprevRes = %g\n", res, prevRes)
		fmt.Printf("math.Sqrt = %g\n", math.Sqrt(x))
	}

	return res
}

func main() {
	fmt.Println(Sqrt(5))
}
