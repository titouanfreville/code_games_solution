package kata

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var weightRe = regexp.MustCompile(`\d+`)

type Weight struct {
	Val string
	Sum int
}

type Weights []Weight

func (w Weight) Lower(w2 Weight) bool {
	if w.Sum == w2.Sum {
		return w.Val < w2.Val
	}

	return w.Sum < w2.Sum
}

func (wl Weights) Sort() Weights {
	sort.Slice(wl, func(i, j int) bool { return wl[i].Lower(wl[j]) })
	return wl
}

func (wl Weights) String() string {
	res := ""
	for _, w := range wl {
		res += w.Val + " "
	}
	return res
}

func WeightFromString(s string) Weight {
	sum := 0
	for _, c := range s {
		v, _ := strconv.Atoi(string(c))
		sum += v
	}

	return Weight{Val: s, Sum: sum}
}

func OrderWeight(s string) string {
	candidates := weightRe.FindAllString(s, -1)
	var weights Weights

	for _, c := range candidates {
		weights = append(weights, WeightFromString(c))
	}

	return strings.TrimSpace(weights.Sort().String())
}
