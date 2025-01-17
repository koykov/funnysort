package funnysort

import (
	"cmp"
	"math/rand"
	"slices"
)

// Hamlet sort array only if decided to do that.
func Hamlet[T cmp.Ordered](a []T) []T {
	toBeOrNotToBe := rand.Intn(2) == 1
	if toBeOrNotToBe {
		slices.Sort(a)
	}
	return a
}
