package main

import (
	"bufio"
	"flag"
	"fmt"
	s "github.com/snail-with-tea/AlgLabs/L3"
	"os"
	"strconv"
	"strings"
)

func getArr() []int {
	stdin := bufio.NewReader(os.Stdin)
	line, err := stdin.ReadString('\n')
	if err != nil {
		panic(err)
	}
	strs := strings.Split(line[0:len(line)-1], " ")
	nums := make([]int, len(strs))
	for i, str := range strs {
		if nums[i], err = strconv.Atoi(str); err != nil {
			panic(fmt.Sprint(str, " is not an integer"))
		}
	}
	return nums
}

func main() {
	rec_s := flag.Bool("r", false, "shorthand for recursive")
	rec_f := flag.Bool("recursive", false, "make merge/quick sort recursive")
	srt_s := flag.String("s", "", "shorthand for sorttype")
	srt_f := flag.String("sorttype", "", "select sort type [quick|merge|shell]")
	ver_s := flag.Int("v", 0, "shorthand for verbosity")
	ver_f := flag.Int("verbosity", 0, "how verbose output should be [0..2]")
	flag.Parse()
	a := flag.Arg(0)
	if a == "" {
		a = *srt_s
	}
	if a == "" {
		a = *srt_f
	}
	recursive := *rec_s || *rec_f
	s.Verbosity = max(*ver_f, *ver_s)

	sorter := func(arr []int) {}

	if a == "" {
		flag.Usage()
		return
	}
	msg := strings.Builder{}
	switch a[0] {
	case 'q':
		msg.WriteString("Sorting with QuickSort Hoare partitioning\n")
		if recursive {
			msg.WriteString("Using recursive approach\n")
			sorter = func(arr []int) {
				s.QuickRecSort(arr, 0, len(arr)-1)
			}
		} else {
			msg.WriteString("Using iterative approach\n")
			sorter = s.QuickItrSort
		}
	case 'm':
		msg.WriteString("Sorting with MergeSort\n")
		if recursive {
			msg.WriteString("Using recursive approach\n")
			sorter = func(arr []int) {
				s.MergeRecSort(arr, 0, len(arr)-1)
			}
		} else {
			msg.WriteString("Using iterative approach\n")
			sorter = s.MergeItrSort
		}
	case 's':
		msg.WriteString("Sorting with ShellSort\n")
		sorter = func(arr []int) {
			s.ShellSort(arr, len(arr))
		}
	default:
		flag.Usage()
		return
	}
	fmt.Println("Enter array separated by spaces:")
	arr := getArr()
	fmt.Print(msg.String())
	fmt.Println(arr)
	sorter(arr)
	fmt.Println(arr)
}
