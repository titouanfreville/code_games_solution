package kata

import (
	"sort"
	"strconv"
)

type Val struct {
	val int
	ind int
}

func Smallest(n int64) []int64 {
	str := strconv.FormatInt(n, 10)

	var (
		number []Val
		sorted []Val

		isNotMinStart = false

		swapA Val
		swapB Val

		i = 0
	)

	for ind, nb := range str {
		intNb, _ := strconv.Atoi(string(nb))
		number = append(number, Val{val: intNb, ind: ind})
		sorted = append(sorted, Val{val: intNb, ind: ind})
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].val < sorted[j].val })

	swapB = sorted[0]
	for sorted[i].val == swapB.val {
		if sorted[i].ind > swapB.ind {
			swapB = sorted[i]
		}
	}

	i = 0

	for isNotMinStart {
		isNotMinStart = number[i].val > sorted[0].val
		i++
	}
	swapA = number[i]

	number[swapA.ind].val = swapB.val
	number[swapB.ind].val = swapA.val

	numRes := ""
	for _, val := range number {
		numRes += strconv.Itoa(val.val)
	}
	num, _ := strconv.ParseInt(numRes, 10, 64)
	return []int64{num, int64(swapA.ind), int64(swapB.ind)}
}
