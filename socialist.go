package funnysort

import "slices"

// Socialist sorts array according principe - help to weak items, but keep the hierarchy.
func Socialist[T OrderedNoString](a []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	slices.Sort(a)
	for i := n - 1; i > 0; i-- {
		d := (a[i] - a[i-1]) / 2
		a[i] -= d
		a[i-1] += d
	}
	return a
}
