package l4

import (
	"fmt"
)

var Verbose = false

func BM_BC_Search[T comparable](txt, pat []T) {
	tab := BM_BC_Shift(pat)
	n := len(txt)
	m := len(pat)
	i := m - 1
	for i < n {
		j := m - 1
		k := i
		for j >= 0 && pat[j] == txt[k] {
			j--
			k--
		}
		if j < 0 {
			fmt.Println("pattern at ", k+1)
			i++
		} else {
			if t, ok := tab[txt[k]]; ok {
				i += t
			} else {
				i += m
			}
		}
	}

}

func BM_BC_Shift[T comparable](pat []T) map[T]int {
	shift := map[T]int{}
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
	return shift
}

func BM_GS_Search[T comparable](txt, pat []T) {
	s := 0
	m := len(pat)
	n := len(txt)
	shift := BM_GS_Shift(pat)
	for s <= n-m {
		j := m - 1
		for j >= 0 && pat[j] == txt[s+j] {
			j -= 1
		}
		if j < 0 {
			fmt.Println("pattern at ", s)
			s += shift[0]
		} else {
			s += shift[j+1]
		}
	}
}

func BM_GS_Shift[T comparable](pat []T) []int {
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

func KMPSearch[T comparable](txt, pat []T) {
	lps := LPSuffix(pat)
	n, m := len(txt), len(pat)
	i, j := 0, 0
	for m-j <= n-i {
		if j == m {
			fmt.Println("pattern at ", i-j)
			j = lps[j-1]
			continue
		}
		if pat[j] == txt[i] {
			i++
			j++
			continue
		}
		if j != 0 {
			j = lps[j-1]
		} else {
			i++
		}

	}
}

func LPSuffix[T comparable](pat []T) []int {
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
	return lps
}
