package main

import (
	"os"
)

func solve() {

}

func main() {
	t := 1
	t = next()
	ans = make([]byte, 0, t*10)
	for t > 0 {
		solve()
		t--
	}
	os.Stdout.Write(ans)
}

var data []byte
var i int
var ans []byte

func init() {
	file, err := os.Open("in.txt")
	if err != nil {
		file = os.Stdin
	}
	defer file.Close()
	stats, _ := file.Stat()
	data = make([]byte, stats.Size())
	file.Read(data)
}

func nexts() []byte {
	word := make([]byte, 0, 12)
	for data[i] == 10 || data[i] == 13 || data[i] == 32 {
		i++
	}
	for data[i] != 32 && data[i] != 10 && data[i] != 13 {
		word = append(word, data[i])
		i++
	}
	return word
}

func next() int {
	return toInt(nexts())
}

func toInt(word []byte) int {
	num := 0
	i := 0
	sign := 1
	if word[i] == 45 {
		sign = -1
		i++
	}
	for i < len(word) {
		num = num*10 + int(word[i]-'0')
		i++
	}
	return num * sign
}

func toByte(num int) []byte {
	word := make([]byte, 0, 12)
	if num == 0 {
		return append(word, '0')
	}
	neg := false
	if num < 0 {
		num = -num
		neg = true
	}
	for num > 0 {
		word = append(word, byte(num%10+'0'))
		num /= 10
	}
	if neg {
		word = append(word, 45)
	}
	reverse(word)
	return word
}

func reverse(bytes []byte) {
	n := len(bytes)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
	}
}

func print(args ...interface{}) {
	for i := 0; i < len(args); i++ {
		switch val := args[i].(type) {
		case int:
			ans = append(ans, toByte(val)...)
		case string:
			ans = append(ans, val...)
		case []byte:
			ans = append(ans, val...)
		case []int:
			for i := 0; i < len(val); i++ {
				ans = append(ans, toByte(val[i])...)
				ans = append(ans, ' ')
			}
		}
		ans = append(ans, ' ')
	}
	ans = append(ans, '\n')
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
