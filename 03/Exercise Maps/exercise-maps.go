package main

import (
	"fmt"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// if s[1] == ' ' {
	// 	fmt.Println("ok")
	// }
	// fmt.Println(string(s[0:3]))
	var a int = 0
	var b int = 0

	for i := range s {
		// fmt.Println(uint8(i))
		if s[i] == ' ' {
			if a == 0 {
				fmt.Println(s[0])
			} else {
				fmt.Println(s[a+1 : b])
			}

			a = int(i)
		} else {
			b = int(i + 1)
		}
	}

	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
