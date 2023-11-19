package methods

import (
	"fmt"
	"math"
)

type F float64

func (f F) M() {
	fmt.Println(f)
}

func InterfaceValues() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
