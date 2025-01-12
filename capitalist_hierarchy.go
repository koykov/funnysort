package funnysort

import (
	"math"
)

func CapitalistHierarchy[T OrderedNoString](a []T) []T {
	n := len(a)
	if n == 0 {
		return a
	}
	a = Capitalist(a)

	topsz, midsz := int(math.Ceil(float64(n)*.1)), int(math.Ceil(float64(n)*.3))
	botsz := n - topsz - midsz

	top, mid, bot := a[0:topsz], a[topsz:topsz+midsz], a[topsz+botsz:]
	// middle -> top
	for i := 0; i < midsz; i++ {
		mid[i] -= 5
		top[i%topsz] += 5
	}
	// middle -> bottom
	for i := 0; i < midsz; i++ {
		mid[i] -= 1
		bot[i%botsz] += 1
	}
	a = Capitalist(a)
	return a
}
