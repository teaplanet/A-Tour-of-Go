package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("error: %v", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}
	return 1.41421356, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
