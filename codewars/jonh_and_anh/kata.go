package kata

var (
	annMemory  []int
	johnMemory []int
	nbDayAnn   int
	nbDayJohn  int
)

func ann(n int) int {
	if n < nbDayAnn {
		return annMemory[n]
	}

	if n == 0 {
		annMemory = append(annMemory, 1)
		nbDayAnn += 1
		return 1
	}

	val := n - john(ann(n-1))
	annMemory = append(annMemory, val)
	nbDayAnn += 1

	return val
}

func john(n int) int {
	if n < nbDayJohn {
		return johnMemory[n]
	}

	if n == 0 {
		johnMemory = append(johnMemory, 0)
		nbDayJohn += 1
		return 0
	}

	val := n - ann(john(n-1))
	johnMemory = append(johnMemory, val)
	nbDayJohn += 1

	return val
}

func Ann(n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, ann(i))
	}
	return
}

func John(n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, john(i))
	}
	return
}

func SumJohn(n int) int {
	res := 0
	toSum := John(n)
	for _, val := range toSum {
		res += val
	}
	return res
}

func SumAnn(n int) int {
	res := 0
	toSum := Ann(n)
	for _, val := range toSum {
		res += val
	}
	return res
}
