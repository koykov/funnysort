package funnysort

import "math"

// Corporate sorts array according principe - top management appropriates all the income, because we're family.
func Corporate[T OrderedNoString](a []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	topsz := int(math.Ceil(float64(n) * .1))
	if topsz == 0 {
		topsz = 1
	}
	a = Capitalist(a)
	_ = a[n-1]
	for i := topsz + 1; i < n; i++ {
		a[i%topsz] += a[i]
		a[i] = 0
	}
	return a
}
