package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solve() {
	ans := 0

	fmt.Fprintln(w, ans)
}

func main() {
	defer w.Flush()
	t := 1
	//t = nextInt()
	for t > 0 {
		solve()
		t--
	}

}

var r *bufio.Reader
var w *bufio.Writer

func init() {
	in, err := os.Open("input.txt")
	if err != nil {
		in = os.Stdin
	}
	r = bufio.NewReader(in)
	out, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		out = os.Stdout
	}
	w = bufio.NewWriter(out)
}

func nextIntSlice() []int {
	line := nextLineSlice()
	subs := bytes.Fields(line)
	s := make([]int, 0, len(subs))
	for i := range subs {
		val, _ := strconv.Atoi(string(subs[i]))
		s = append(s, val)
	}
	return s
}

func nextLineSlice() []byte {
	line, hasMore, _ := r.ReadLine()
	if hasMore {
		tmp := slices.Clone(line)
		for hasMore {
			line, hasMore, _ = r.ReadLine()
			tmp = append(tmp, line...)
		}
		line = tmp
	}
	return line
}

func nextLine() []byte {
	line, _ := r.ReadBytes('\n')
	return bytes.TrimSpace(line)
}

func nextString() string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(string(s))
}

func nextInt() int {
	line := nextLineSlice()
	val, _ := strconv.Atoi(string(line))
	return val
}

func next2Int() (int, int) {
	line := nextLineSlice()
	nums := bytes.Fields(line)
	v1, _ := strconv.Atoi(string(nums[0]))
	v2, _ := strconv.Atoi(string(nums[1]))
	return v1, v2
}

func Min[T int | float64](b ...T) (min T) {
	min = b[0]
	for i := 1; i < len(b); i++ {
		if b[i] < min {
			min = b[i]
		}
	}
	return
}

func Max[T int | float64](b ...T) (max T) {
	max = b[0]
	for i := 1; i < len(b); i++ {
		if b[i] > max {
			max = b[i]
		}
	}
	return
}

func Abs[T int | float64](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
