package moretypes

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	return [][]uint8{
		{0, 0},
		{0, 0},
		{0, 0},
	}
}

func ExerciceSlices() {
	pic.Show(Pic)
}
