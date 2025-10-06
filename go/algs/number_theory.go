package algs

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for {
		if a == 0 {
			return b
		}
		b %= a
		a, b = b, a
	}
}

func phi(n int) int {
	ans := n
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			n /= i
		}
		ans -= ans / i
	}
	if n > 1 {
		ans -= ans / n
	}
	return ans
}

func digitSum(n int) int {
	ans := 0
	for n > 0 {
		ans += n % 10
		n /= 10
	}
	return ans
}
