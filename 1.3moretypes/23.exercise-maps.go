package moretypes

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func ExerciceMaps() {
	wc.Test(WordCount)
}
