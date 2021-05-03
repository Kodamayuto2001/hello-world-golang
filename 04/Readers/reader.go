/*
	ioパッケージは、データストリームを読むことを表現するio.Readerインターフェースを規定しています。

	Goの標準ライブラリには、ファイル、ネットワーク接続、圧縮、暗号化などで、個のインターフェースの多くの実装があります。

	io.ReaderインターフェースはReadメソッドをもちます。
	func(T) Read(b []byte) (n int, err error)

	Readは、データを与えられたバイトスライスへ入れ、入れたバイトのサイズとエラーの値を返します。ストリームの終端は、io.EOFのエラーで返します。package Readers例のコードは、strings.Readerを作成し、8byte毎に読みだしています。
*/

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
