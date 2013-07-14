package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := complex128(1.0)
	for diff := z; cmplx.Abs(diff) > 1e-10; {
		diff = (cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2))
		z -= diff
	}
	return z
}

func main() {
	r := Cbrt(2)
	fmt.Println(r)
	fmt.Println(cmplx.Pow(r, 3))
}
