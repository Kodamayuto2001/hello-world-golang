/*
	Goのプログラムは、errorの状態をerror値で表現します。
	error型は、fmt.Stringerに似た組み込みのインターフェースです：
	type error interface {
		Error() string
	}
	(fmt.Stringerと同様に、fmtパッケージは、変数を文字列で出力する際にerrorインターフェースを確認します。)
	よく、関数はerror変数を返します。そして呼び出し元はエラーがnilかどうかを確認することでerrorをハンドル(取り扱い)します。
*/

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
