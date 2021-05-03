package main

import "fmt"

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z, ErrNegativeSqrt(x)
}

// func Sqrt(x float64) float64 {
// 	var z float64 = 1
// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 	}
// 	return z
// }

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
	}
	return "nil"
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
