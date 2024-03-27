package L3

import "fmt"

var Verbosity = 0

func ShellSort(arr []int, size int) {
	for gap := size / 2; gap > 0; gap /= 2 {
		for i := gap; i < size; i++ {
			if Verbosity > 1 {
				fmt.Println(arr)
			}
			for j := i - gap; j >= 0 && arr[j] > arr[j+gap]; j -= gap {
				if Verbosity > 1 {
					fmt.Println(j, j+gap)
				}
				arr[j], arr[j+gap] = arr[j+gap], arr[j]
			}
		}
		if Verbosity > 0 {
			fmt.Println(arr, "\n---")
		}
	}
}

func MergeRecSort(arr []int, l int, r int) {
	if l < r {
		m := l + (r-l)/2
		MergeRecSort(arr, l, m)
		MergeRecSort(arr, m+1, r)
		merge(arr, l, m, r)
		if Verbosity == 1 {
			fmt.Println(arr)
		}
		if Verbosity > 1 {
			fmt.Println(arr, l, m, r)
		}
	}
}

func MergeItrSort(arr []int) {
	size := len(arr)
	if size < 2 {
		return
	}
	for blk := 1; blk <= size; blk *= 2 {
		for l := 0; l < size; l += blk * 2 {
			if Verbosity > 1 {
				fmt.Println(arr)
			}
			r := min(l+blk*2-1, size-1)
			m := min(l+blk-1, size-1)
			merge(arr, l, m, r)
		}
		if Verbosity > 0 {
			fmt.Println(arr, "\n---")
		}
	}
}

func merge(arr []int, l int, m int, r int) {
	s_l := m - l + 1
	// s_r := r - m
	ltm := make([]int, s_l)
	copy(ltm, arr[l:m+1])
	/* mtr := make([]int, s_r)
	copy(mtr, arr[m+1:r+1]) */
	ins, lpt, rpt := l, 0, m+1

	for lpt < s_l && rpt < r+1 {
		if ltm[lpt] < arr[rpt] {
			arr[ins] = ltm[lpt]
			lpt++
		} else {
			arr[ins] = arr[rpt]
			rpt++
		}
		ins++
	}

	for lpt < s_l {
		arr[ins] = ltm[lpt]
		lpt++
		ins++
	}

	for rpt < r+1 {
		arr[ins] = arr[rpt]
		rpt++
		ins++
	}
}

func QuickRecSort(arr []int, l int, r int) {
	if l < r {
		m := hoarePart(arr, l, r)
		if Verbosity == 1 {
			fmt.Println(arr)
		}
		if Verbosity > 1 {
			fmt.Println(arr, l, m, r)
		}
		QuickRecSort(arr, l, m)
		QuickRecSort(arr, m+1, r)
	}
}

func QuickItrSort(arr []int) {
	size := len(arr)
	if size < 2 {
		return
	}
	stack := []int{}
	stack = append(stack, 0, size-1)
	sttop := 1
	for sttop >= 0 {
		r := stack[sttop]
		l := stack[sttop-1]
		sttop -= 2
		stack = stack[:sttop+1]
		m := hoarePart(arr, l, r)
		if Verbosity == 1 {
			fmt.Println(arr)
		}
		if Verbosity > 1 {
			fmt.Println(arr, l, m, r)
		}
		if m > l {
			stack = append(stack, l, m)
			sttop += 2
		}
		if m+1 < r {
			stack = append(stack, m+1, r)
			sttop += 2
		}
	}
}

func hoarePart(arr []int, l int, r int) int {
	m := l + (r-l)/2
	lpt, rpt := l, r
	pivot := arr[m]
	for true {
		for arr[lpt] < pivot {
			lpt++
		}
		for arr[rpt] > pivot {
			rpt--
		}
		if lpt >= rpt {
			return rpt
		}
		arr[lpt], arr[rpt] = arr[rpt], arr[lpt]
		lpt++
		rpt--
	}
	return rpt
}
