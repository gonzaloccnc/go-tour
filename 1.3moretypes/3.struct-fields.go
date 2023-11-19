package moretypes

import "fmt"

func StructsFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
