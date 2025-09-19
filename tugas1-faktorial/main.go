package main

import (
	"fmt"
	"math"
)

func calculateFactorial(n int) float64 {
	if n == 0 {
		return 1.0
	}

	var result float64 = 1.0
	for i := 1; i <= n; i++ {
		result *= float64(i)
	}
	return result
}

func F(n int) float64 {
	faktorialN := calculateFactorial(n)

	pangkat2N := math.Pow(2, float64(n))

	hasilBagi := faktorialN / pangkat2N

	hasilAkhir := math.Ceil(hasilBagi)

	return hasilAkhir
}

func main() {
	n := 10
	fmt.Printf("F(%d) = %.0f\n", n, F(n))

	n = 5
	fmt.Printf("F(%d) = %.0f\n", n, F(n))
}