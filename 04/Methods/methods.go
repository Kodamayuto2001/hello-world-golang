/*
	Goには、クラス(class)の仕組みはありませんが、型にメソッド(method)を定義できます。
	メソッドは、特別なレシーバー(receiver)引数を関数に撮ります。
	レシーバは、funcキーワードとメソッド名の間に自信の引数リストで表現します。
	この例では、Absメソッドはvという名前のVertex片野レシーバをもつことを意味しています。
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
