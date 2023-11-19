package moretypes

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	return func() int {
		return 0
	}
}

func ExerciseWithFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
