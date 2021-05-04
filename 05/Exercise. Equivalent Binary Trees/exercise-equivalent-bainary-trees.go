package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

//	Walk walks the tree  sending all values
//	from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int) {
	//	最初のwalkの位置の値
	ch <- t.Value

	tmp := []*tree.Tree{}
	tmp[0] = t
	for {

	}
}

//	Same determines whether the trees
//	t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
