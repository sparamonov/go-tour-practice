package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	res := 1.0
	prevRes := 2.0
	
	for math.Abs(prevRes - res) > 0.000000001 {
		prevRes = res
		res = res - ((math.Pow(res, 2) - x) / (2 * res))
		
		fmt.Printf("res = %g\nprevRes = %g\n", res, prevRes)
		fmt.Printf("math.Sqrt = %g\n", math.Sqrt(x))
	}

	return res, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
