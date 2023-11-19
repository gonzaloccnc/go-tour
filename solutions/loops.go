package solutions

import (
	"fmt"
	"math"
)

const delta2 = 1e-6

func Sqrt2(x float64) float64 {
	z := x
	n := 0.0
	for math.Abs(n-z) > delta2 {
		n, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func LoopsSolve() {
	const x = 2
	mine, theirs := Sqrt2(x), math.Sqrt(x)
	fmt.Println(mine, theirs, mine-theirs)
}
