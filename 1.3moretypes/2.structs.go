package moretypes

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Structs() {
	fmt.Println(VertexMap{1, 2})
}
