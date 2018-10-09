package recursive_algorithms_in_go

func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n - 2) + fibo(n - 1)
}