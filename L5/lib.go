package l5

import (
	"fmt"
	l3 "github.com/snail-with-tea/AlgLabs/L3"
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
			g[x+y+1][y] = g[y][x+y+1]
		}
	}
	fmt.Println()
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
	size := len(matr)
	conns := []Conn{}
	for y := range size - 1 {
		for x := range size - y - 1 {
			c := Conn{x, y, matr[y][x+1]}
			conns = append(conns, c)
		}
	}
	return size, conns
}

func MST_Kruskal(size int, cons []Conn) []Conn {
	l3.QuickConf(cons, func(a, b int) bool {
		return cons[a].Len > cons[b].Len
	})

	mst := []Conn{}

	setindx := make([]int, size)
	setnext := 1
	c_len := len(cons)
	for c_len > 0 {
		c_len--
		con := cons[c_len]
		cons = cons[:c_len]
		z1 := setindx[con.Vert1] == 0
		z2 := setindx[con.Vert2] == 0
		if !z1 && !z2 {
			continue
		}
		mst = append(mst, con)
		if z1 && z2 {
			setindx[con.Vert1] = setnext
			setindx[con.Vert2] = setnext
			setnext++
		}
		if z1 && !z2 {
			setindx[con.Vert1] = setindx[con.Vert2]
		}
		if !z1 && z2 {
			setindx[con.Vert2] = setindx[con.Vert1]
		}

	}

	return mst

}

func MST_Prim(matr [][]int) []Conn {
	size := len(matr)
	used := make([]bool, size)
	dot := 42 % size
	used[dot] = true
	unused := size - 1
	conns := []Conn{}
	mst := []Conn{}
	for unused > 0 {
		con := Conn{0, 0, 0}

		for i := range size {
			l := matr[dot][i]
			if l == 0 || used[i] {
				continue
			}

			conns = append(conns, Conn{dot, i, l})
		}

		for _, c := range conns {
			if con.Len == 0 || c.Len < con.Len {
				con.Len = c.Len
				con.Vert1 = c.Vert1
				con.Vert2 = c.Vert2
			}
		}
		if con.Len == 0 {
			continue
		}

		used[con.Vert2] = true
		unused--
		dot = con.Vert2

		mst = append(mst, con)

		// clear up active connections
		// OPTIONAL
		con_len := len(conns)
		for i := 0; i < con_len; {
			if conns[i].Vert2 == con.Vert2 {
				if i+1 < con_len {
					conns = append(conns[:i], conns[i+1:]...)
				} else {
					conns = conns[:i]
				}
				con_len--
			} else {
				i++
			}
		}
	}
	// fmt.Println(mst)
	return mst
}

func MST_Count(matr [][]int) int {
	size := len(matr)
	kirh := make([][]int, size)
	for i := range size {
		kirh[i] = make([]int, size)
	}
	for y := range size {
		con_count := 0
		for x := range size {
			if matr[y][x] != 0 {
				con_count++
				kirh[y][x] = -1
			}
		}
		kirh[y][y] = con_count
	}
	for y := range size {
		fmt.Println(kirh[y])
	}
	m := Minor(kirh, size, 1, 1)
	return Det(m, size-1)
}

func Det(matr [][]int, size int) int {
	if size <= 0 {
		return 0
	}
	if size == 1 {
		return matr[0][0]
	}
	if size == 2 {
		return matr[0][0]*matr[1][1] - matr[0][1]*matr[1][0]
	}
	det := 0
	for x := range size {
		minr := Minor(matr, size, x, 0)
		m_det := Det(minr, size-1)
		sign := -1
		if x%2 == 0 {
			sign = 1
		}
		det += sign * matr[0][x] * m_det
	}
	return det
}

func Minor(matr [][]int, size, x, y int) [][]int {
	minr := make([][]int, size-1)
	i_m := 0
	for i_o := range size {
		if i_o == y {
			continue
		}
		j_m := 0
		minr[i_m] = make([]int, size-1)
		for j_o := range size {
			if j_o == x {
				continue
			}
			minr[i_m][j_m] = matr[i_o][j_o]
			j_m++
		}
		i_m++
	}
	return minr
}
