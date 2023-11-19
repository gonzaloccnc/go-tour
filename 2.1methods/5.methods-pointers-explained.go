package methods

import (
	"fmt"
	"math"
)

func AbsMethodsPointersExplained(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MethodsPointersExplained() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(AbsMethodsPointersExplained(v))
}
