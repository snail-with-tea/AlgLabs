package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	l5 "github.com/snail-with-tea/AlgLabs/L5"
)

func main() {
	v1 := flag.Bool("verbose", false, "output steps of search algorithm")
	v2 := flag.Bool("v", false, "shorthand for verbose")
	s1 := flag.String("searcher", "", "search algorithm [Kruskal|Prim|All]")
	s2 := flag.String("s", "", "shorthand for searcher")
	e1 := flag.String("entry", "", "graph entry method [Matrix|Connections]")
	e2 := flag.String("e", "", "shorthand for entry")
	flag.Parse()
	if *v1 || *v2 {
		l5.Verbose = true
	}
	if *s1 == "" {
		*s1 = *s2
	}
	if *s1 == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *e1 == "" {
		*e1 = *e2
	}
	if *e1 == "" {
		flag.Usage()
		os.Exit(1)
	}
	s := strings.ToLower(*s1)
	e := strings.ToLower(*e1)

	if (s[0] != 'k' && s[0] != 'p' && s[0] != 'a') || (e[0] != 'm' && e[0] != 'c') {
		flag.Usage()
		os.Exit(1)
	}

	graph := [][]int{}
	conns := []l5.Conn{}
	n_cnt := 0
	switch e[0] {
	case 'm':
		n_cnt, graph = l5.EnterSimMatr()
	case 'c':
		n_cnt, conns = l5.EnterSimConn()
	}
	if len(graph) == 0 {
		graph = l5.SimConnToSimMatr(n_cnt, conns)
		if l5.Verbose {
			fmt.Println("Graph as weight matrix:")
			l5.PrintMat(graph, n_cnt)
		}
	}
	if len(conns) == 0 {
		conns = l5.SimMatrToSimConn(n_cnt, graph)
		if l5.Verbose {
			fmt.Println("Graph as connections list:")
			l5.PrintCon(conns, false)
		}
	}

	mst_k := []l5.Conn{}
	mst_p := []l5.Conn{}
	weight_k := 0
	weight_p := 0
	switch s[0] {
	case 'k':
		mst_k, weight_k = l5.MST_Kruskal(n_cnt, conns)

		fmt.Println("MST weight = ", weight_k)
		fmt.Println("MST as connections:")
		l5.PrintCon(mst_k, false)
	case 'p':
		mst_p, weight_p = l5.MST_Prim(n_cnt, graph)

		fmt.Println("MST weight = ", weight_p)
		fmt.Println("MST as connections:")
		l5.PrintCon(mst_p, false)
	case 'a':
		mst_k, weight_k = l5.MST_Kruskal(n_cnt, conns)
		mst_p, weight_p = l5.MST_Prim(n_cnt, graph)
		forest := l5.ST_Count(graph)
		fmt.Println("Kruskal method:")
		fmt.Println("MST weight = ", weight_k)
		fmt.Println("MST as connections:")
		l5.PrintCon(mst_k, false)
		fmt.Println("Prim method:")
		fmt.Println("MST weight = ", weight_p)
		fmt.Println("MST as connections:")
		l5.PrintCon(mst_p, false)
		eq := ""
		if weight_k == weight_p {
			eq = "Weights are equal, solution is correct!"
		} else {
			eq = "Weights differ, there is error somewhere..."
		}
		fmt.Println(eq)

		if TEQ(mst_k, mst_p) {
			eq = "Kruskal & Prim methods found equal MST's"
		} else {
			eq = "Kruskal & Prim methods found different MST's"
		}
		fmt.Println("Total number of spanning trees(forest) = ", forest)
	}
}

func TEQ(mst_1, mst_2 []l5.Conn) bool {
	for i := range mst_1 {
		c_1 := mst_1[i]
		f := -1
		for j := range mst_2 {
			c_2 := mst_2[j]
			if (c_1.Vert1 == c_2.Vert1 && c_1.Vert2 == c_2.Vert2) || (c_1.Vert1 == c_2.Vert2 && c_1.Vert2 == c_2.Vert1) {
				f = j
				break
			}
		}
		if f < 0 {
			return false
		}
	}

	return true
}
