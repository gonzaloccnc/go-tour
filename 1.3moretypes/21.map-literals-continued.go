package moretypes

import "fmt"

var mlc = map[string]VertexMap{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func MapLiteralsContinued() {
	fmt.Println(mlc)
}
