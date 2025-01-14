package funnysort

import (
	"cmp"
	"slices"
)

// Robespierre sort array according principe - all items denied by terrorPolicyFunc removes from thr sorted array.
// The greatest element falls victim to terror afterward.
func Robespierre[T cmp.Ordered](a []T, terrorPolicyFunc func([]T) []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	slices.Sort(a)
	a = terrorPolicyFunc(a)
	if len(a) > 0 {
		a = a[:len(a)-1]
	}
	return a
}

func RobespierreTerrorPolicyExample_HasNoSiblings[T cmp.Ordered](a []T) []T {
	for i := 0; i < len(a)-1; i++ {
		var c, p int
		for j := i; j < len(a); j++ {
			if a[i] == a[j] {
				c++
				p = j
				continue
			}
			break
		}
		if c > 1 {
			a = append(a[:i], a[p+1:]...)
		} else {
			i = p
		}
	}
	return a
}

func RobespierreTerrorPolicyExample_HasSiblings[T cmp.Ordered](a []T) []T {
	for i := 0; i < len(a); i++ {
		var c, p int
		for j := i; j < len(a); j++ {
			if a[i] == a[j] {
				c++
				p = j
				continue
			}
			break
		}
		if c == 1 {
			a = append(a[:i], a[i+1:]...)
			i--
		} else {
			i = p
		}
	}
	return a
}
