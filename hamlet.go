package funnysort

import (
	"cmp"
	"math/rand"
	"slices"
)

// Hamlet sort array only if decided to do that.
func Hamlet[T cmp.Ordered](a []T) []T {
	const (
		notToBe = iota
		toBe
	)
	decision := rand.Intn(2)
	switch decision {
	case toBe:
		slices.Sort(a)
	case notToBe:
		fallthrough
	default:
		return a
	}
	return a
}
