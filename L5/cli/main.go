package main

import (
	"fmt"

	l5 "github.com/snail-with-tea/AlgLabs/L5"
)

func main() {
	size, conns := l5.EnterSimConn()
	m := l5.SimConnToMatr(size, conns)
	for _, l := range m {
		fmt.Println(l)
	}
}
