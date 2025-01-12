package funnysort

import (
	"cmp"
	"math/rand"
	"slices"
)

// Yeltsin sorts array by combining random, liberalism, populism and reformation principles.
func Yeltsin[T cmp.Ordered](a []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	rand.Shuffle(n, func(i, j int) { a[i], a[j] = a[j], a[i] })
	reformStart := n / 3
	slices.Sort(a[reformStart:])
	return a
}
