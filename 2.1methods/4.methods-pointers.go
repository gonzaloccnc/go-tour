package methods

import (
	"fmt"
	"math"
)

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MethodsPointers() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
