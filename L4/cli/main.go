package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	l4 "github.com/snail-with-tea/AlgLabs/L4"
)

func get_line() string {
	stdin := bufio.NewReader(os.Stdin)
	line, err := stdin.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.Trim(line, " \r\n\t")
	return line
}

func main() {
	v1 := flag.Bool("verbose", false, "output steps of search algorithm")
	v2 := flag.Bool("v", false, "shorthand for verbose")
	s1 := flag.String("searcher", "", "search algorithm [BoyerMoore|KnuthMorrisPratt]")
	s2 := flag.String("s", "", "shorthand for searcher")
	flag.Parse()
	if *v1 || *v2 {
		l4.Verbose = true
	}
	if *s1 == "" {
		*s1 = *s2
	}
	if *s1 == "" {
		flag.Usage()
		os.Exit(1)
	}

	alg := strings.ToLower(*s1)

	search := func(txt, pat string) []int { return []int{} }
	switch alg[0] {
	case 'b':
		search = l4.BM_BC_Search
	case 'k':
		search = l4.KMPSearch
	case 'g':
		search = l4.BM_GS_Search
	default:
		flag.Usage()
		os.Exit(1)
	}
	fmt.Println("Enter text where search:")
	txt := get_line()
	fmt.Println("Enter pattern to search:")
	pat := get_line()
	f := search(txt, pat)
	if len(f) > 0 {
		fmt.Println("Found pattern at", f)
	} else {
		fmt.Println("Pattern not found")
	}
}
