package main

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0
	for i:=0; i<10; i++ {
		s := z
		z -= ((z * z) - x) / (2 * z)
		fmt.Println(s, z)
		if (z - s < 0.0000000000000000001) {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
