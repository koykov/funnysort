package funnysort

import (
	"slices"
	"sort"
)

// Hitler sorts array according principe - all items groups by checkSkullFormFunc and the values of isolated items
// redistributes between survived items.
func Hitler[T OrderedNoString](a []T, checkSkullFormFunc func(T) bool) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	sort.Slice(a, func(i, j int) bool {
		skullI, skullJ := checkSkullFormFunc(a[i]), checkSkullFormFunc(a[j])
		switch {
		case skullI && !skullJ:
			return false
		case !skullI && skullJ:
			return true
		}
		return true
	})
	var (
		p = -1
		s T
	)
	for i := 0; i < n; i++ {
		if checkSkullFormFunc(a[i]) {
			if p == -1 {
				p = i
			}
		} else {
			s += a[i]
		}
	}
	if p == -1 {
		slices.Sort(a)
		return a
	}
	dist := s / T(p)
	for i := 0; i < p; i++ {
		a[i] += dist
	}
	slices.Sort(a[:p])
	return a[:p]
}
