/*
	mapはキーと値とを関連付けます。（マップします）
	マップのゼロ値はnilです。
	nilマップはキーをもっておらず、またキーを追加することもできません。
	make関数は指定された方のマップを初期化して、使用可能な状態で返します。
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
}
