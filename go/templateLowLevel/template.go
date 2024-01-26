package main

import (
	"os"
)

func solve() {

}

func main() {
	stats, _ := os.Stdin.Stat()
	data = make([]byte, stats.Size())
	os.Stdin.Read(data)
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
	left := 0
	right := len(bytes) - 1
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
}

func print(n int) {
	num := []byte{}
	for n >= 0 {
		num = append(num, byte(n%10+'0'))
		n /= 10
		if n == 0 {
			break
		}
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

var data []byte
var i int
var ans []byte
