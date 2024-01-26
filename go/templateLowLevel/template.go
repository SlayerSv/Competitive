package main

import (
	"os"
)

func solve() {

}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		file = os.Stdin
	}
	defer file.Close()
	stats, _ := file.Stat()
	data = make([]byte, stats.Size())
	file.Read(data)
	t := 1
	t = next()
	ans = make([]byte, 0, t*10)
	for t > 0 {
		solve()
		t--
	}
	os.Stdout.Write(ans)
}

func next() int {
	num := 0
	for (data[i] == 10 || data[i] == 13 || data[i] == 32) && i < len(data) {
		i++
	}
	sign := 1
	if data[i] == 45 {
		sign = -1
		i++
	}
	for data[i] != 32 && data[i] != 10 && data[i] != 13 && i < len(data) {
		num = num*10 + int(data[i]-'0')
		i++
	}
	num *= sign
	return num
}

func reverse(bytes []byte) {
	n := len(bytes)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
	}
}

func print(n int) {
	num := []byte{}
	if n == 0 {
		ans = append(ans, '0')
		return
	}
	for n > 0 {
		num = append(num, byte(n%10+'0'))
		n /= 10
	}
	reverse(num)
	num = append(num, ' ')
	ans = append(ans, num...)
}

func min[T int | float64](a ...T) (min T) {
	min = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return
}

func max[T int | float64](a ...T) (max T) {
	max = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return
}

func abs[T int | float64](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

var data []byte
var i int
var ans []byte
