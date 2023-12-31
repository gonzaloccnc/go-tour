package solutions

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f, g := 1, 0
	return func() int {
		f, g = g, f+g
		return f
	}
}

func FibSolve() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
