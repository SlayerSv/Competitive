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

func write[T int | int64 | float64 | string](n T) {
	fmt.Fprintln(w, n)
}

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)
