package solutions

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkImpl2(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkImpl2(t.Left, ch)
	ch <- t.Value
	walkImpl2(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk2(t *tree.Tree, ch chan int) {
	walkImpl2(t, ch)
	// Need to close the channel here
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// NOTE: The implementation leaks goroutines when trees are different.
// See binarytrees_quit.go for a better solution.
func Same2(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)

	go Walk2(t1, w1)
	go Walk2(t2, w2)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func BinaryTreesSolve() {
	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}
