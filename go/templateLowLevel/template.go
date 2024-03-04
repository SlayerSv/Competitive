package main

import (
	"os"
)

func solve() {

}

func main() {
	defer out.Close()
	t := 1
	//t = next()
	for t > 0 {
		solve()
		t--
	}
	out.Write(ans)
}

var data []byte
var i int
var ans []byte
var buff []byte
var size int
var in *os.File
var out *os.File

func init() {
	ans = make([]byte, 0, 200)
	buff = make([]byte, 30)
	var err error
	in, err = os.Open("input.txt")
	if err != nil {
		in = os.Stdin
	}
	defer in.Close()
	out, err = os.OpenFile("outs.txt", os.O_WRONLY|os.O_TRUNC, 0222)
	if err != nil {
		out = os.Stdout
	}
	stats, _ := in.Stat()
	size = int(stats.Size())
	data = make([]byte, size)
	in.Read(data)
}

func skipWS() {
	for data[i] == 10 || data[i] == 13 || data[i] == 32 {
		i++
	}
}

func next() []byte {
	skipWS()
	j := 0
	for i < size && data[i] != 32 && data[i] != 13 && data[i] != 10 {
		buff[j] = data[i]
		i++
		j++
	}
	return buff[:j]
}

func nexts() []byte {
	w := make([]byte, 0)
	skipWS()
	for i < size && data[i] != 32 && data[i] != 13 && data[i] != 10 {
		w = append(w, data[i])
		i++
	}
	return w
}

func nexti() int {
	return toInt(next())
}

func nextf() float64 {
	return toFloat(next())
}

func nextSi(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = nexti()
	}
	return s
}

func toInt(word []byte) int {
	num := 0
	j := 0
	sign := 1
	if word[j] == 45 {
		sign = -1
		j++
	}
	for j < len(word) {
		num = num*10 + int(word[j]-'0')
		j++
	}
	return num * sign
}

func toFloat(word []byte) float64 {
	m := 0.0
	d := 0
	j := 0
	sign := 1.0
	if word[j] == 45 {
		sign = -1.0
		j++
	}
	n := len(word)
	for j < n {
		if word[j] == '.' {
			d = j
			j++
			continue
		}
		m = m*10 + float64(word[j]-'0')
		j++
	}
	m *= sign
	if d > 0 {
		for n-d > 1 {
			m /= 10
			d++
		}
	}
	return m
}

func toByte(num int) []byte {
	if num == 0 {
		buff[0] = '0'
		return buff[:1]
	}
	neg := false
	if num < 0 {
		num = -num
		neg = true
	}
	j := 0
	for ; num > 0; j++ {
		buff[j] = byte(num%10 + '0')
		num /= 10
	}
	if neg {
		buff[j] = 45
		j++
	}
	reverse(buff[:j])
	return buff[:j]
}

func reverse(bytes []byte) {
	for l, r := 0, len(bytes)-1; l < r; l, r = l+1, r-1 {
		bytes[l], bytes[r] = bytes[r], bytes[l]
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
		case []string:
			for i := 0; i < len(val); i++ {
				ans = append(ans, val[i]...)
				ans = append(ans, ' ')
			}
		}
		ans = append(ans, ' ')
	}
	ans = append(ans, '\n')
}

func Min[T int | float64](a ...T) (min T) {
	min = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return
}

func Max[T int | float64](a ...T) (max T) {
	max = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
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
