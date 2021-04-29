package main

import "fmt"

func main() {
	// defer ステートメントは、deferへ渡した関数の実行を、呼び出し元の関数の終わり(return する)まで遅延させるものです。
	// deferへ渡した関数の引数は、すぐに評価されますが、その関数事態は呼び出し元の関数がreturnするまで実行されません。

	defer fmt.Println("world")

	fmt.Println("hello")
}
