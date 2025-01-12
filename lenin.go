package funnysort

import "slices"

// Lenin sorts array according principe - better to have less amount of items, but only worthy.
func Lenin[T OrderedNoString](a []T, threshold T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	slices.Sort(a)
	var p int
	for p = 1; p < n; p++ {
		if a[p] > threshold {
			break
		}
	}
	a = append(a[:0], a[p:]...)
	return a
}
