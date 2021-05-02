package main

import "fmt"

//	fibonacci is a function that returns a function that returns an int.

//	0,1,1,2,3,5,....
func fibonacci() func(int) int {
	old := 1
	old2 := 0
	var y int = 0
	return func(x int) int {
		if x == 0 {
			y = 0
		} else if x == 1 {
			y = 1
		} else {
			y = old2 + old
			old2 = old
			old = y
		}
		return y
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
