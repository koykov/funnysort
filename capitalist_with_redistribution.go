package funnysort

func CapitalistWithRedistribution[T OrderedNoString](a []T) []T {
	n := len(a)
	if n == 0 {
		return a
	}
	a = Capitalist[T](a)
	for i := 0; i < n-1; i++ {
		a[i] += a[i+1] / 2
		a[i+1] /= 2
	}
	return a
}
