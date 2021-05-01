/*
	スライスするときは、それらの既定値の代わりに使用することで上限または、加減を省略することができます。
	既定値は、下限が０で、上限はスライスの長さです。

*/
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
