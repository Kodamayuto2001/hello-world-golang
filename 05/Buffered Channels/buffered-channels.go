/*
	チャネルは、バッファ(buffer)として使えます。バッファをもつチャネルを初期化するには、makeの2つ目の引数にバッファの長さを与えます。

	ch := make(chan int,100)

	バッファが詰まったときは、チャネルへの送信をブロックします。バッファが空の時は、チャネルの受信をブロックします。
	バッファが積んるようにサンプルコードを変更して、何が起きるのかを見てください。

*/

package main

import "fmt"

func main() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
