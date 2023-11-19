package moretypes

import "fmt"

var ml = map[string]VertexMap{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func MapLiterals() {
	fmt.Println(ml)
}
