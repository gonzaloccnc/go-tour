package moretypes

import "fmt"

func MakingSlices() {
	a := make([]int, 5)
	printSliceWithMake("a", a)

	b := make([]int, 0, 5)
	printSliceWithMake("b", b)

	c := b[:2]
	printSliceWithMake("c", c)

	d := c[2:5]
	printSliceWithMake("d", d)
}

func printSliceWithMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
