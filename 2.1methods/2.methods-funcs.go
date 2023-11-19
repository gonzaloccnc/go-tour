package methods

import (
	"fmt"
	"math"
)

func AbsMethodsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MethodsFuncs() {
	v := Vertex{3, 4}
	fmt.Println(AbsMethodsFunc(v))
}
