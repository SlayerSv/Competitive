package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() {

}

func main() {
	t := 1
	t = next()
	for t > 0 {
		solve()
		t--
	}
	w.Flush()
}

func next() int {
	var n int
	fmt.Fscan(r, &n)
	return n
}

func nexts() string {
	var n string
	fmt.Fscan(r, &n)
	return n
}

func nextf() float64 {
	var n float64
	fmt.Fscan(r, &n)
	return n
}

func print(n ...interface{}) {
	fmt.Fprintln(w, n...)
}

func min[T int | int64 | float64 | string](b ...T) (min T) {
	min = b[0]
	for i := 1; i < len(b); i++ {
		if b[i] < min {
			min = b[i]
		}
	}
	return
}

func max[T int | int64 | float64 | string](b ...T) (max T) {
	max = b[0]
	for i := 1; i < len(b); i++ {
		if b[i] > max {
			max = b[i]
		}
	}
	return
}

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)
