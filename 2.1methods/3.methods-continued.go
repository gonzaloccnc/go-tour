package methods

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) AbsMehodsContinued() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func MethodsContinued() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.AbsMehodsContinued())
}
