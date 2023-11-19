package methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) AbsMethods() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods() {
	v := Vertex{3, 4}
	fmt.Println(v.AbsMethods())
}
