package funnysort

// CapitalistWithRedistribution sorts array according principe - rich items growing his fortune, poor shrinks fortune.
func CapitalistWithRedistribution[T OrderedNoString](a []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	a = Capitalist[T](a)
	_ = a[n-1]
	for i := 0; i < n-1; i++ {
		a[i] += a[i+1] / 2
		a[i+1] /= 2
	}
	return a
}
