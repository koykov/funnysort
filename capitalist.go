package funnysort

import (
	"cmp"
	"slices"
)

// Capitalist sorts array according capitalist principles - rich items has first positions.
func Capitalist[T cmp.Ordered](a []T) []T {
	slices.Sort(a)
	slices.Reverse(a)
	return a
}
