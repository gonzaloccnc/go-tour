package moretypes

import "fmt"

func SliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSliceOtherWay(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSliceOtherWay(s)

	// Extend its length.
	s = s[:4]
	printSliceOtherWay(s)

	// Drop its first two values.
	s = s[2:]
	printSliceOtherWay(s)
}

func printSliceOtherWay(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}