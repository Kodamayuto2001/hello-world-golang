/*
	Pic関数実装チャレンジ
		生成した画像が表示される

		引数
			dx
				各要素が8bitのunsigned int型のスライスの長さ

			dy
				長さ（スライス）
			dyにdxを割り当てる
*/

package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
