package funnysort

import "slices"

// LateSocialist sorts only half of array, but reports about success and surpassing the result of previous five-year plan.
func LateSocialist[T OrderedNoString](a []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	slices.Sort(a[:n/2])
	return a
}
