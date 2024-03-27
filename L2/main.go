package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
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
			if arr[j] < arr[min] {
				min = j
			}
			arr[min], arr[i] = arr[i], arr[min]
		}

	}
	return arr
}

func bubSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

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
	ts := flag.String("s", "", "shorthand for sorttype")
	tf := flag.String("sorttype", "", "select sort type [insert|select|bubble]")
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
