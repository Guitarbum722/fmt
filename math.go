package box

// IsPrime returns true if input n is a prime number
func IsPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	if n < 2 {
		return false
	}
	return true
}
