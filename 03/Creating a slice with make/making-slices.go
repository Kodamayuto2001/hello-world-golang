/*
	スライスは、組み込みのmake関数を使用して作成することができます。これは、動的サイズの配列を作成する方法です。

	make関数はゼロかされた配列を割り当て、その配列を指すスライスを返します。

	makeの3番目の引数に、スライスの容量を指定できます。cap(b)で、スライスの容量を返します。

*/

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
