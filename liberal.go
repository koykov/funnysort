package funnysort

// Liberal avoids sorting because of all elements are equal and has right to be in its place.
func Liberal[T any](a []T) []T {
	return a
}
