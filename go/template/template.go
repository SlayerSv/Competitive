package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func runTestCase() {
	n := nextInt()

	ans := solve(n)

	fmt.Fprintln(w, ans)
}

func solve(n int) int {
	ans := n

	return ans
}

func main() {
	defer w.Flush()
	t := 1
	//t = nextInt()
	for t > 0 {
		runTestCase()
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
	stats, _ := in.Stat()
	r = bufio.NewReaderSize(in, int(stats.Size()))
	out, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		out = os.Stdout
	}
	w = bufio.NewWriter(out)
}

func nextIntSlice() []int {
	line := nextLineSlice()
	subs := bytes.Fields(line)
	s := make([]int, len(subs))
	i, n := 0, len(subs)
	for i < n {
		val, err := strconv.Atoi(string(subs[i]))
		if err != nil {
			panic(err)
		}
		s[i] = val
		i++
	}
	return s
}

func nextLineSlice() []byte {
	line, hasMore, err := r.ReadLine()
	if err != nil && err.Error() != "EOF" {
		panic(err)
	}
	if hasMore {
		tmp := slices.Clone(line)
		for hasMore {
			line, hasMore, err = r.ReadLine()
			if err != nil {
				panic(err)
			}
			tmp = append(tmp, line...)
		}
		line = tmp
	}
	return line
}

func nextLine() []byte {
	line, err := r.ReadBytes('\n')
	if err != nil && err.Error() != "EOF" {
		panic(err)
	}
	return line
}

func nextString() string {
	s, err := r.ReadString('\n')
	if err != nil && err.Error() != "EOF" {
		panic(err)
	}
	return strings.TrimSpace(string(s))
}

func nextInt() int {
	line := nextLineSlice()
	val, err := strconv.Atoi(string(line))
	if err != nil {
		panic(err)
	}
	return val
}

func next2Int() (int, int) {
	line := nextLineSlice()
	nums := bytes.Fields(line)
	if len(nums) != 2 {
		panic("next2int expects exactly 2 values")
	}
	v1, err := strconv.Atoi(string(nums[0]))
	if err != nil {
		panic(err)
	}
	v2, err := strconv.Atoi(string(nums[1]))
	if err != nil {
		panic(err)
	}
	return v1, v2
}

func next3Int() (int, int, int) {
	line := nextLineSlice()
	nums := bytes.Fields(line)
	if len(nums) != 3 {
		panic("next3int expects exactly 3 values")
	}
	v1, err := strconv.Atoi(string(nums[0]))
	if err != nil {
		panic(err)
	}
	v2, err := strconv.Atoi(string(nums[1]))
	if err != nil {
		panic(err)
	}
	v3, err := strconv.Atoi(string(nums[2]))
	if err != nil {
		panic(err)
	}
	return v1, v2, v3
}

func Abs[T int | float64](a T) T {
	if cmp.Less(a, 0) {
		a = -a
	}
	return a
}
