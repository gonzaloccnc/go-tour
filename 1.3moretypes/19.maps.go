package moretypes

import "fmt"

type VertexMap struct {
	Lat, Long float64
}

var m map[string]VertexMap

func Maps() {
	m = make(map[string]VertexMap)
	m["Bell Labs"] = VertexMap{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
