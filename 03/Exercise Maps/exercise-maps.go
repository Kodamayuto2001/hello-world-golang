package main

import (
	"golang.org/x/tour/wc"
)

// func WordCount(s string) map[string]int {
// 	// if s[1] == ' ' {
// 	// 	fmt.Println("ok")
// 	// }
// 	// fmt.Println(string(s[0:3]))
// 	var a int = 0
// 	var b int = 0
// 	m := make(map[string]int)

// 	for i := range s {
// 		// fmt.Println(uint8(i))
// 		if s[i] == ' ' {
// 			if a == 0 {
// 				m[string(s[0])] += 1
// 				fmt.Println(string(s[0]))
// 			} else {
// 				m[string(s[a+1:b])] += 1
// 				fmt.Println(s[a+1 : b])
// 			}

// 			a = int(i)
// 		} else {
// 			b = int(i + 1)
// 		}
// 	}
// 	m[string(s[a+1:b])] += 1
// 	fmt.Println(s[a+1 : b])

// 	fmt.Println(m)

// 	return m
// }

func WordCount(s string) map[string]int {
	var a int = 0
	var b int = 0
	m := make(map[string]int)

	for i := range s {
		if s[i] == ' ' {
			b = i
			// fmt.Println(string(s[a:b]))
			m[string(s[a:b])] += 1
			a = b + 1
		}
	}
	// fmt.Println(string(s[a:]))
	m[string(s[a:])] += 1
	return m
}

func main() {
	wc.Test(WordCount)
}
