package flowcontrol

import (
	"fmt"
	"math"
)

func sqrtWithElse(x float64) string {
	if x < 0 {
		return sqrtWithElse(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func If() {
	fmt.Println(sqrtWithElse(2), sqrtWithElse(-4))
}
