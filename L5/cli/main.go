package main

import (
	"fmt"

	l5 "github.com/snail-with-tea/AlgLabs/L5"
)

func main() {
	size, conns := l5.EnterSimConn()
	fmt.Println(conns)
	mst_k := l5.MST_Kruskal(size, conns)
	fmt.Println(mst_k)
	matr := l5.SimConnToMatr(size, conns)
	mst_p := l5.MST_Prim(matr)
	fmt.Println(mst_p)
}
