package funnysort

import "math"

// Woke sort array according principe - weak items robs the middle class, top items grabs bottom afterward.
func Woke[T OrderedNoString](a []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	a = Capitalist(a)

	topsz, midsz := int(math.Ceil(float64(n)*.1)), int(math.Ceil(float64(n)*.3))
	botsz := n - topsz - midsz

	top, mid, bot := a[0:topsz], a[topsz:topsz+midsz], a[topsz+botsz:]
	// middle -> bottom
	for i := 0; i < midsz; i++ {
		mid[i] -= 5
		bot[i%botsz] += 5
	}
	// bottom -> top
	for i := 0; i < midsz; i++ {
		bot[i] -= 6
		top[i%topsz] += 6
	}
	a = Capitalist(a)
	return a
}
