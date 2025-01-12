package funnysort

// Communist sorts array by assigning to all elements an average value of all elements.
func Communist[T OrderedNoString](a []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	var s T
	_ = a[n-1]
	for i := 0; i < n; i++ {
		s += a[i]
	}
	avg := s / T(n)
	for i := 0; i < n; i++ {
		a[i] = avg
	}
	return a
}
