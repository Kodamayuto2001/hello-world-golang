/*
	Range

	for ループに利用するrangeは、スライスや、マップを一つずつ反復処理するために使います。
	スライスをrangeで繰り返す場合、rangeは反復枚に2つの変数を返します。
	一つ目の変数はインデックス(index)で、二つ目は、インデックスの場所の要素のコピーです。
*/

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
