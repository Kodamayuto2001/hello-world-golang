package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// func (r rot13Reader) Read(p []byte) (n int, err error) {
// 	// fmt.Println(r)
// 	// {0xc00004a3c0}がたくさんでた
// 	// fmt.Println(p)
// 	//	0がたくさんでた

// 	for i := range p {
// 		if p[i] == 'L' {
// 			fmt.Println("Lがでた")
// 		}
// 	}

// 	return 0, nil
// 	//	固まった
// 	// return len(p), nil
// 	//	無限ループになった
// }

//	解答
func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	for i, v := range b {
		switch {
		case v >= 'A' && v < 'N', v >= 'a' && v < 'n':
			b[i] += 13
		case v >= 'N' && v <= 'Z', v >= 'n' && v <= 'z':
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
