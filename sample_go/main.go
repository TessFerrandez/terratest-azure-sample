package main

import (
	"fmt"
	"terratest-azure-sample/sample_go/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
