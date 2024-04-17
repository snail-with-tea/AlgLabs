package main

import (
	"fmt"

	l5 "github.com/snail-with-tea/AlgLabs/L5"
)

func main() {
	// size, conns := l5.EnterSimConn()
	// fmt.Println(conns)
	// mst_k := l5.MST_Kruskal(size, conns)
	// fmt.Println(mst_k)
	// matr := l5.SimConnToMatr(size, conns)
	// mst_p := l5.MST_Prim(matr)
	// fmt.Println(mst_p)
	m := [][]int{
		{0, 7, 15, 12, 0, 10, 0},
		{7, 0, 13, 9, 0, 0, 8},
		{15, 13, 0, 7, 15, 7, 0},
		{12, 9, 7, 0, 9, 0, 11},
		{0, 0, 15, 9, 0, 10, 0},
		{10, 0, 7, 0, 10, 0, 12},
		{0, 8, 0, 11, 0, 12, 0},
	}
	fmt.Println(l5.MST_Count(m))
}
