package funnysort

import "math/rand"

// Bakunin makes array chaotic due to:
// * all items are equal, no one must take proper or privileged position
// * any attempt to make an order suppresses of elements freedom
// * the result must be randomized due to chaos is natural state of the system
func Bakunin[T any](a []T) []T {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}
