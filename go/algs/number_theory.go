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
		if b == 0 {
			return a
		}
		a %= b
		a, b = b, a
	}
}

func egcd(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x, y := egcd(b, a%b)
	return g, y, x - y*(a/b)
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

// returns NOT primes bool array
func sieve(n int) []bool {
	notPrimes := make([]bool, n+1)
	notPrimes[0], notPrimes[1] = true, true
	for i := 2; i*i <= n; i++ {
		if !notPrimes[i] {
			for j := i * i; j <= n; j += i {
				notPrimes[j] = true
			}
		}
	}
	return notPrimes
}
