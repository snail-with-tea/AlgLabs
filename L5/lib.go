package l5

import (
	"fmt"
	"sort"
)

type Conn struct {
	Vert1 int
	Vert2 int
	Len   int
}

func EnterSimMatr() [][]int {
	fmt.Print("Enter node count: ")
	size := 0
	fmt.Scanln(&size)
	g := make([][]int, size)
	for y := range size {
		g[y] = make([]int, size)
	}
	fmt.Println("Enter connection matrix\n(0 means no connection):")
	fmt.Print("Y\\X  ")
	for x := range size {
		fmt.Print(x, "  ")
	}
	fmt.Println()
	for y := range size {
		fmt.Print("  ", y, " ")
		for range y + 1 {
			fmt.Print(" - ")
		}
		for x := range size - y - 1 {
			fmt.Scan(&g[y][x+y+1])
		}
	}
	return g
}

func EnterSimConn() (int, []Conn) {
	conns := []Conn{}
	fmt.Println("Enter connections (vert1,vert2,len):")
	size := 0
	for {
		c := Conn{0, 0, 0}
		fmt.Scan(&c.Vert1, &c.Vert2, &c.Len)
		if c.Len <= 0 || c.Vert1 == c.Vert2 {
			break
		}
		size = max(size, c.Vert1, c.Vert2)
		conns = append(conns, c)
	}
	return size + 1, conns
}

func SimConnToMatr(size int, conns []Conn) [][]int {
	g := make([][]int, size)
	for y := range size {
		g[y] = make([]int, size)
	}
	for _, c := range conns {
		g[c.Vert1][c.Vert2] = c.Len
		g[c.Vert2][c.Vert1] = c.Len
	}
	return g
}

func SimMatrToSimConn(matr [][]int) (int, []Conn) {
	size_y := len(matr)
	size_x := len(matr[0])
	conns := []Conn{}
	for y := range size_y - 1 {
		for x := range size_x - y - 1 {
			c := Conn{x, y, matr[y][x+1]}
			conns = append(conns, c)
		}
	}
	return max(size_x, size_y), conns
}

func MST_Kruskal(size int, conns []Conn) []Conn {
	sort.Slice(conns, func(a, b int) bool {
		return conns[a].Len < conns[b].Len
	})
	mst := []Conn{}

	return mst

}

func MST_Prim() {

}
