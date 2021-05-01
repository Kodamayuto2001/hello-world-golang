/*
	map mの操作を見ていきましょう。

	mへ要素(elem)の挿入や更新：
	m[key] = elem

	要素の取得
	elem = m[key]

	要素の削除
	delete(m,key)

	キーに対する要素が存在するかどうかは、2つ目の値で確認します。
	elem,ok = m[key]

	もし、mにkeyがあれば、変数okはtrueとなり、存在しなければ、okはfalseになります。

	なお、mapにkeyが存在しない場合、elemはmapの要素の型のゼロ値となります。

	Note:
		もし、elemやokをまだ宣言していなければ、次のように:=で短く宣言できます

		elem,ok := m[key]
*/

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Anser"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
