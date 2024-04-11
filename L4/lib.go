package l4

import (
	"fmt"
	"strings"
)

var Verbose = false

func BM_BC_Search(txt, pat string) []int {
	t := []rune(txt)
	p := []rune(pat)
	tab := BM_BC_Shift(p)
	n := len(t)
	m := len(p)
	i := m - 1
	found := []int{}
	for i < n {
		j := m - 1
		k := i

		if Verbose {
			fmt.Println()
			fmt.Print(strings.Repeat(" ", i))
			fmt.Println("i")
			fmt.Println(txt)
			fmt.Print(strings.Repeat(" ", i-j))
			fmt.Println(pat)
			fmt.Print(strings.Repeat(" ", i))
			fmt.Println("j")
		}

		for j >= 0 && p[j] == t[k] {
			j--
			k--
		}
		if j < 0 {
			if Verbose {
				fmt.Println("pattern at ", k+2)
			}
			found = append(found, k+2)
			i++
		} else {
			if t, ok := tab[t[k]]; ok {
				i += t
			} else {
				i += m
			}
		}
	}
	return found
}

func BM_BC_Shift(pat []rune) map[rune]int {
	shift := map[rune]int{}
	m := len(pat)
	shift[pat[m-1]] = m
	for j := range m - 1 {
		if s, ok := shift[pat[j]]; ok {
			if m-j-1 < s {
				shift[pat[j]] = m - j - 1
			}
		} else {
			shift[pat[j]] = m - j - 1
		}
	}

	if Verbose {
		fmt.Println("Bad character heuristic:")
		u, d := string(""), string("")
		for i, v := range pat {
			u += string(v) + " "
			d += fmt.Sprint(shift[pat[i]]) + " "
		}
		fmt.Println("", u, "\n", d)

	}

	return shift
}

func BM_GS_Search(txt, pat string) []int {
	t := []rune(txt)
	p := []rune(pat)
	s := 0
	m := len(p)
	n := len(t)
	shift := BM_GS_Shift(p)
	found := []int{}
	for s <= n-m {
		j := m - 1
		for j >= 0 && p[j] == t[s+j] {
			j -= 1
		}
		if j < 0 {
			if Verbose {
				fmt.Println("pattern at ", s+1)
			}
			found = append(found, s+1)
			s += shift[0]
		} else {
			s += shift[j+1]
		}
	}
	return found
}

func BM_GS_Shift(pat []rune) []int {
	m := len(pat)
	bordr := make([]int, m+1)
	shift := make([]int, m+1)
	// find strong suffix
	i := m
	j := m + 1
	bordr[i] = j
	for i > 0 {
		for j <= m && pat[i-1] != pat[j-1] {
			if shift[j] == 0 {
				shift[j] = j - i
			}
			j = bordr[j]
		}
		i--
		j--
		bordr[i] = j
	}
	// find strong prefix
	j = bordr[0]
	for i := range m + 1 {
		if shift[i] == 0 {
			shift[i] = j
		}
		if i == j {
			j = bordr[j]
		}
	}
	return shift
}

func KMPSearch(txt, pat string) []int {
	t := []rune(txt)
	p := []rune(pat)
	lps := LPPrefix(p)
	n, m := len(t), len(p)
	i, j := 0, 0
	found := []int{}
	for m-j <= n-i {
		if j == m {
			if Verbose {
				fmt.Println()
				fmt.Print(strings.Repeat(" ", i))
				fmt.Println("i")
				fmt.Println(txt)
				fmt.Print(strings.Repeat(" ", i-j))
				fmt.Println(pat)
				fmt.Print(strings.Repeat(" ", i))
				fmt.Println("j")
				fmt.Println("pattern at ", i-j+1)
			}
			found = append(found, i-j+1)
			j = lps[j-1]
			continue
		}
		if p[j] == t[i] {
			i++
			j++
			continue
		}
		if Verbose {
			fmt.Println()
			fmt.Print(strings.Repeat(" ", i))
			fmt.Println("i")
			fmt.Println(txt)
			fmt.Print(strings.Repeat(" ", i-j))
			fmt.Println(pat)
			fmt.Print(strings.Repeat(" ", i))
			fmt.Println("j")
		}
		if j != 0 {
			j = lps[j-1]
		} else {
			i++
		}
	}
	return found
}

func LPPrefix(pat []rune) []int {
	m := len(pat)
	s := 0
	lps := make([]int, m)
	lps[0] = 0
	for i := 1; i < m; {
		if pat[i] == pat[s] {
			s++
			lps[i] = s
			i++
			continue
		}
		if s != 0 {
			s = lps[s-1]
		} else {
			lps[i] = 0
			i++
		}
	}

	if Verbose {
		fmt.Println("Longest possible prefix:")
		u, d := string(""), string("")
		for i, v := range pat {
			u += string(v) + " "
			d += fmt.Sprint(lps[i]) + " "
		}
		fmt.Println("", u, "\n", d)

	}

	return lps
}
