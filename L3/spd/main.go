package main

import (
	"bufio"
	"flag"
	"io/fs"
	"sync"

	// "flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	s "github.com/snail-with-tea/AlgLabs/L3"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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
		if nums[i], err = strconv.Atoi(str); err != nil || nums[i] < 1 {
			panic(fmt.Sprint(str, " is not a positive integer"))
		}
	}
	return nums
}

type record struct {
	T float64
	P int
}

func timeAll(n uint) [5]float64 {
	funcs := make([]func([]int), 5)
	funcs[0] = func(arr []int) {
		s.ShellSort(arr, len(arr)-1)
	}
	funcs[1] = s.MergeItrSort
	funcs[2] = func(arr []int) {
		s.MergeRecSort(arr, 0, len(arr)-1)
	}
	funcs[3] = s.QuickItrSort
	funcs[4] = func(arr []int) {
		s.QuickRecSort(arr, 0, len(arr)-1)
	}

	chn := make(chan record, 5)
	reslt := make([]float64, 5)
	wg := new(sync.WaitGroup)

	for ind, fun := range funcs {
		wg.Add(1)
		go timeOne(n, ind, fun, wg, chn)
	}

	go func(wg *sync.WaitGroup, chn chan record) {
		wg.Wait()
		close(chn)
	}(wg, chn)

	for rec := range chn {
		reslt[rec.P] = rec.T
	}

	return [5]float64(reslt)
}

var tip = 0

func timeOne(num uint, ind int, fun func([]int), wg *sync.WaitGroup, chn chan record) {
	defer wg.Done()
	arr := make([]int, num)
	switch tip {
	case 0:
		for i := range num {
			arr[i] = rand.Int()
		}
	case 1:
		for i := range num {
			arr[i] = int(i)
		}
	case 2:
		for i := range num {
			arr[i] = int(num - i)
		}
	}
	start := time.Now()
	fun(arr)
	timed := max(time.Now().Sub(start).Seconds(), 0.0)
	chn <- record{timed, ind}
}

func main() {
	n_test := flag.Uint("n", 10, "number of tests for each plot point")
	flag.Parse()

	fmt.Println("array of points to test")
	elements := getArr()
	size := len(elements)
	pt_sh := make([]plotter.XYer, size)
	pt_mi := make([]plotter.XYer, size)
	pt_mr := make([]plotter.XYer, size)
	pt_qi := make([]plotter.XYer, size)
	pt_qr := make([]plotter.XYer, size)
	err := os.Mkdir("./assets", fs.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
	for t := range 3 {
		tip = t
		for i, el := range elements {
			xy_sh := make(plotter.XYs, *n_test)
			xy_mi := make(plotter.XYs, *n_test)
			xy_mr := make(plotter.XYs, *n_test)
			xy_qi := make(plotter.XYs, *n_test)
			xy_qr := make(plotter.XYs, *n_test)
			pt_sh[i] = xy_sh
			pt_mi[i] = xy_mi
			pt_mr[i] = xy_mr
			pt_qi[i] = xy_qi
			pt_qr[i] = xy_qr

			for j := range *n_test {
				res := timeAll(uint(el))
				// fmt.Println(res)
				xy_sh[j].X = float64(el)
				xy_mi[j].X = float64(el)
				xy_mr[j].X = float64(el)
				xy_qi[j].X = float64(el)
				xy_qr[j].X = float64(el)
				xy_sh[j].Y = res[0]
				xy_mi[j].Y = res[1]
				xy_mr[j].Y = res[2]
				xy_qi[j].Y = res[3]
				xy_qr[j].Y = res[4]
			}
		}
		plt := plot.New()

		mean95_sh, err_sh := plotutil.NewErrorPoints(plotutil.MeanAndConf95, pt_sh...)

		mean95_mi, err_mi := plotutil.NewErrorPoints(plotutil.MeanAndConf95, pt_mi...)
		mean95_mr, err_mr := plotutil.NewErrorPoints(plotutil.MeanAndConf95, pt_mr...)

		mean95_qi, err_qi := plotutil.NewErrorPoints(plotutil.MeanAndConf95, pt_qi...)
		mean95_qr, err_qr := plotutil.NewErrorPoints(plotutil.MeanAndConf95, pt_qr...)
		switch {
		case err_sh != nil:
			panic(err_sh)
		case err_mi != nil:
			panic(err_mi)
		case err_mr != nil:
			panic(err_mr)
		case err_qi != nil:
			panic(err_qi)
		case err_qr != nil:
			panic(err_qr)
		}
		plotutil.AddLinePoints(plt,
			"Shell\nSort", mean95_sh,
			"Merge\nIterative", mean95_mi,
			"Merge\nRecursive", mean95_mr,
			"Quick\nIterative", mean95_qi,
			"Quick\nRecursive", mean95_qr,
		)
		plotutil.AddYErrorBars(plt, mean95_sh, mean95_mi, mean95_mr, mean95_qi, mean95_qr)
		plt.Legend.Top = true
		plt.Legend.Left = true
		plt.X.Scale = plot.LogScale{}
		plt.X.Label.Text = "Elements"
		plt.X.Tick.Marker = plot.LogTicks{}
		plt.Y.Scale = plot.LogScale{}
		plt.Y.Label.Text = "Seconds"
		plt.Y.Tick.Marker = plot.LogTicks{}
		switch tip {
		case 0:
			plt.Title.Text = "Random numbers"
			plt.Save(16*vg.Centimeter, 8*vg.Centimeter, "./assets/random.svg")
			plt.Save(vg.Points(800), vg.Points(400), "./assets/random.png")
		case 1:
			plt.Title.Text = "Already sorted"
			plt.Save(16*vg.Centimeter, 8*vg.Centimeter, "./assets/sorted.svg")
			plt.Save(vg.Points(800), vg.Points(400), "./assets/sorted.png")
		case 2:
			plt.Title.Text = "Reverse sorted"
			plt.Save(16*vg.Centimeter, 8*vg.Centimeter, "./assets/reversed.svg")
			plt.Save(vg.Points(800), vg.Points(400), "./assets/reversed.png")
		}

	}
}
