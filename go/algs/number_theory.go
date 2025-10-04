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
