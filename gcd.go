package recursive_algorithms_in_go

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a % b)
}