package funnysort

// MilitaryCommunist sorts array according principe - the strongest item expropriates values of all other items due to
// circumstances.
func MilitaryCommunist[T OrderedNoString](a []T) []T {
	n := len(a)
	if n < 2 {
		return a
	}
	var (
		p  = -1
		mx T
	)
	_ = a[n-1]
	for i := 0; i < n; i++ {
		if a[i] > mx {
			mx = a[i]
			p = i
		}
	}
	if p == -1 {
		p = n - 1
	}
	for i := 0; i < n; i++ {
		if i != p {
			a[p] += a[i]
			a[i] = 0
		}
	}
	return a
}
