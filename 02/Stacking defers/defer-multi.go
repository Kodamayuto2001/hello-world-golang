/*
deferへ渡した関数が複数ある場合、その呼び出しはスタック(stack)されます。呼び出し元の関数がreturn するとき、deferへ渡した関数はLIFO(last-in-first-out)の順番で実行されます。
*/

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
