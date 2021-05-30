package kata

func fibSumb(n int) int {
	sum := 0
	var n1, n2 int
	for i := 0; i < n; i++ {
		if i == 0 {
			n1 = 1
			sum += 1
		} else if i == 1 {
			n2 = 1
			sum += 1
		} else {
			sum += n1 + n2
			n1 = n2
			n2 = i
		}
	}
	return sum
}

func Perimeter(n int) int {
	return 4 * fib(n+1)
}
