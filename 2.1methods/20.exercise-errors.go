package methods

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func ExerciceErrors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
