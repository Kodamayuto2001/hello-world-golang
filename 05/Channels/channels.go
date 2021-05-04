/*
	チャネル(Channel)型は、チャネルオペレータの<-を用いて値の送受信ができる通り道です。
	ch <- v		//	vをチャネルchへ送信する
	v := <- ch 	//	ch火r受信した変数をvへ割り当てる
	(データは、↩の方向に流れます)
	マップとスライスのように、チャネルを使う前に以下のように生成します。
	ch := make(chan int)

	通常、片方が準備できるまで送受信はブロックされます。これにより、明確なロックや条件変数がなくても、goroutineの動機を可能にします。
	サンプルコードは、スライス内の数値を合算し、2つのgoroutine間で作業を分配します。
	両方のgoroutineで計算が完了すると、最終結果が計算されます。
*/

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //receive from c

	fmt.Println(x, y, x+y)
}
