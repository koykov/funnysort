package funnysort

import "cmp"

// Stalin sorts array by removing all items breaking the order.
func Stalin[T cmp.Ordered](a []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	buf, p, mx := a, 1, a[0]
	_ = a[n-1]
	for i := 1; i < n; i++ {
		if a[i] >= mx {
			mx = a[i]
			buf[p] = mx
			p++
		}
	}
	return buf[:p]
}
