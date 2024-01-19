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

func write(n int) {
	fmt.Fprintln(w, n)
}

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)
