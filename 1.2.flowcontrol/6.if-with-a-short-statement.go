package flowcontrol

import (
	"fmt"
	"math"
)

func powWithElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func IfWithAShortStament() {
	fmt.Println(
		powWithElse(3, 2, 10),
		powWithElse(3, 3, 20),
	)
}
