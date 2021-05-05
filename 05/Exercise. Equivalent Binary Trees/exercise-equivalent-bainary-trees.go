// package main

// import (
// 	"fmt"

// 	"golang.org/x/tour/tree"
// )

// // type Tree struct {
// //     Left  *Tree
// //     Value int
// //     Right *Tree
// // }

// //	Walk walks the tree  sending all values
// //	from the tree to the channel ch.

// func Walk(t *tree.Tree, ch chan int) {
// 	//	最初のwalkの位置の値
// 	ch <- t.Value

// 	tmp := []*tree.Tree{}
// 	var i int = 0
// 	tmp[i] = t

// 	for {
// 		if t.Left != nil {
// 			t = t.Left
// 			ch <- t.Value
// 			i += 1
// 			tmp[i] = t
// 		} else if t.Right != nil {
// 			t = t.Right
// 			ch <- t.Value
// 			i += 1
// 			tmp[i] = t
// 		} else {
// 			if i < 9 {

// 			} else {

// 			}
// 		}
// 	}
// }

// //	Same determines whether the trees
// //	t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool {
// 	return false
// }

// func main() {
// 	ch := make(chan int)
// 	go Walk(tree.New(1), ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		switch {
		case !ok1, !ok2:
			return ok1 == ok2
		case v1 != v2:
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
