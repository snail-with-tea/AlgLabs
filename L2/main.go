package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Verbosity = 0

// 87 80 84 17 66 94 69 23 79 15
// 15 1 91 30 57 91 80 13 79 96

func insSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			if Verbosity > 1 {
				fmt.Println(arr, "|", arr[j], "<", arr[j-1])
			}
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
		if Verbosity > 0 {
			fmt.Println(arr, "\n---")
		}
	}
	return arr
}

func selSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] > arr[min] {
				continue
			}
			min = j
			if Verbosity > 1 {
				fmt.Println(arr, "| min:", arr[min], "at", min)
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
		if Verbosity > 0 {
			fmt.Println(arr, "\n---")
		}
	}
	return arr
}

func bubSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				if Verbosity > 1 {
					fmt.Println(arr, "|", arr[j+1], ">", arr[j])
				}
			}
		}
		if Verbosity > 0 {
			fmt.Println(arr, "\n---")
		}
	}
	return arr
}

func getArr() []int {
	fmt.Println("Enter array separated by spaces")
	stdin := bufio.NewReader(os.Stdin)
	line, err := stdin.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.Trim(line, " \r\n\t")
	strs := strings.Split(line, " ")
	nums := []int{}
	num := 0
	for _, str := range strs {
		if num, err = strconv.Atoi(str); err != nil {
			fmt.Println(fmt.Sprint(str, " is not an integer: skipping"))
		} else {
			nums = append(nums, num)
		}
	}
	return nums
}

func main() {
	ts := flag.String("s", "", "shorthand for sorttype")
	tf := flag.String("sorttype", "", "select sort type [insert|select|bubble]")
	vs := flag.Int("v", 0, "shorthand for verbosity")
	vf := flag.Int("verbosity", 0, "how verbose the output should be [0..2]")
	flag.Parse()
	a := flag.Arg(0)
	if a == "" {
		a = *ts
	}
	if a == "" {
		a = *tf
	}
	if a == "" {
		flag.Usage()
		return
	}

	Verbosity = max(Verbosity, *vs, *vf)
	sorter := func(arr []int) []int {
		return arr
	}
	switch a[0] {
	case 'i':
		sorter = insSort
	case 's':
		sorter = selSort
	case 'b':
		sorter = bubSort
	default:
		flag.Usage()
		return
	}
	arr := getArr()
	arr = sorter(arr)
	fmt.Println(arr)
}
