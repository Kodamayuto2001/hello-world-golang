/*
	チャネルが、gorutine間で通信するための素晴らしい方法であることを見てきました。

	しかし、通信が必要ない場合はどうでしょう？コンフリクトを避けるために、一度に１つのgorutineだけが変数にアクセスできるようにしたい場合はどうでしょうか？

	このコンセプトは、排他制御(mutual exclusion)と呼ばれ、このデータ構造を指す一般的な名前は、mutexです。

	Goの標準ライブラリは、排他制御をsync.Mutexと次の二つのメソッドで提供します。

	Lock
	Unlock

	Incメソッドにあるように、LockとUnlockｄ囲むことで排他制御で実行するコードを定義できます。

	Valueメソッドのように、mutexがUnlockされることを保証するためにdeferを使うこともできます。
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

//	SafeCounter si safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

//	Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	//	Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

//	Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	//	Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
