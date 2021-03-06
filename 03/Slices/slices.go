/*

	配列は固定長です。一方で、スラ出羽可変長です。より柔軟な配列とみなすこともできます。
	実際には、スライスは配列よりもより一般的です。

	型[]T は型Tのスライスを表します。

	転んで区切られた二つのインデックスlowとhighの境界を指定することによってスライスが形成されます。

	a[low : high]

	これは最初の要素は含むが、最後の要素は除いた半開区間を選択します。


	次の式はaの要素のうち1から3を含むスライスを作ります。
	a[1:4]
*/

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
