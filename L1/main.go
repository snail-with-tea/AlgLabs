package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	m "math"
	"os"
	"strconv"
	"strings"
)

func t1(x float64, y float64) float64 {
	u := m.Sqrt(m.Abs(x-1)) - m.Sqrt(m.Abs(y))
	d := 1 + x*x/2 + y*y/4
	return u / d
}

func t2(x float64) float64 {
	y := m.Cos(x) + 0.2
	switch {
	case x < 0:
		y = x*x*x - 1.5
	case x > m.Pi/2:
		y = x*x + x*2
	}

	return y
}

func t3(x float64, y float64) int {
	dst := m.Sqrt(x*x + y*y)
	switch {
	case dst <= 2:
		return 1
	case dst <= 4:
		return 2
	default:
		return 0
	}
}

func t4(arr []int) int {
	acc := 0
	for _, e := range arr {
		acc += e

	}
	return acc
}

func t5(arr []int) (int, int) {
	acc := 0
	mul := 1
	for _, e := range arr {
		if e > 0 {
			acc += e
			mul *= e
		}
	}
	return acc, mul
}

func t6(x float64) (float64, error) {
	if !(0 < x && x <= 4) {
		return 0, errors.New("X must satisfy 0<X≤4")
	}
	acc := 0.0
	for i := range 7 {
		acc += m.Pow(x, float64(i+1))
	}
	return acc, nil
}

func t7(x float64) (float64, error) {
	if !(m.Pi/3 < x && x <= m.Pi) {
		return 0.0, errors.New("X must satisfy π/3<X≤π")
	}
	acc := 0.0
	add := m.Cos(x)
	itr := 1
	for m.Abs(add) > m.Pow10(-4) {
		acc += add
		itr += 1
		add = m.Cos(x*float64(itr)) / float64(itr)
	}
	return acc, nil
}

func h1() {
	var x, y float64
	fmt.Print("Enter x,y (separated by space): ")
	fmt.Scanln(&x, &y)
	r := t1(x, y)
	fmt.Println("Output:", r)

}
func h2() {
	var x float64
	fmt.Print("Enter x: ")
	fmt.Scanln(&x)
	r := t2(x)
	fmt.Println("Output:", r)

}
func h3() {
	var x, y float64
	fmt.Print("Enter x,y (separated by space): ")
	fmt.Scanln(&x, &y)
	r := t3(x, y)
	fmt.Println("Output:", r)

}

func h4() {
	fmt.Print("Enter array separated by spaces: ")
	arr := getArr()
	r := t4(arr)
	fmt.Println("Output:", r)
}

func h5() {
	fmt.Print("Enter array separated by spaces: ")
	arr := getArr()
	acc, mul := t5(arr)
	fmt.Println("Sum of positives:", acc, "\nProduct of positives:", mul)
}

func h6() {
	var x float64
	fmt.Print("Enter x (0<X≤4): ")
	fmt.Scanln(&x)
	r, err := t6(x)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Output:", r)
}

func h7() {
	var x float64
	fmt.Print("Enter x (π/3<X≤π): ")
	fmt.Scanln(&x)
	r, err := t7(x)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Output:", r)
}

func getArr() []int {
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
	tf := flag.String("task", "", "task to run [1..=7]")
	ts := flag.String("t", "", "task to run [1..=7]")
	flag.Parse()
	a := flag.Arg(0)
	if a == "" {
		a = *ts
	}
	if a == "" {
		a = *tf
	}
	task := 0
	err := errors.New("")
	if task, err = strconv.Atoi(a); err != nil {
		flag.Usage()
		return
	}

	switch task {
	case 1:
		h1()
	case 2:
		h2()
	case 3:
		h3()
	case 4:
		h4()
	case 5:
		h5()
	case 6:
		h6()
	case 7:
		h7()
	default:
		flag.Usage()
	}

}
