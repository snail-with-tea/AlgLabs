package main

import (
	"fmt"

	l4 "github.com/snail-with-tea/AlgLabs/L4"
)

func main() {
	txt := []rune("ABAABABACABABCABABCABAB")
	pat := []rune("ABA")
	l4.BM_BC_Search(txt, pat)
	fmt.Println(l4.BM_BC_Shift(pat))
}
