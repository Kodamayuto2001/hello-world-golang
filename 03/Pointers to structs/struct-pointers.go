//	structのフィールドは、structのポインタを通してアクセスすることもできます。

//	フィールドXをもつstructのポインタpがある場合、フィールドXにアクセスするには(*p)/Xのように書くことができます。しかし、この表記法は大変面倒ですので、Goでは代わりにp.Xと書くこともできます。

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
