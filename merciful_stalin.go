package funnysort

import (
	"cmp"
	"slices"
)

// MercifulStalin sorts array applies four steps:
// * apply Stalin sort forward
// * apply Stalin sort backward
// * merge results
// * recursively apply MercifulStalin to the rest and merge
func MercifulStalin[T cmp.Ordered](buf *MercifulStalinBuffer[T], a []T) []T {
	n := len(a)
	if n <= 1 {
		return a
	}
	buf.prealloc(uint32(n))
	fw, rfw, bw, rbw, r, r1 := buf.fw(), buf.rfw(), buf.bw(), buf.rbw(), buf.r(), buf.r1()

	// forward sort
	mx := a[0]
	fw = append(fw, mx)
	_ = a[n-1]
	for i := 1; i < n; i++ {
		if a[i] >= mx {
			mx = a[i]
			fw = append(fw, mx)
		} else {
			rfw = append(rfw, a[i])
		}
	}

	// backward sort
	if m := len(rfw); m > 0 {
		mn := rfw[m-1]
		bw = append(bw, mn)
		_ = rfw[:m-1]
		for i := m - 2; i >= 0; i-- {
			if rfw[i] <= mn {
				mn = rfw[i]
				bw = append(bw, mn)
			} else {
				rbw = append(rbw, rfw[i])
			}
		}
		slices.Reverse(bw)
	}

	// merge
	r = mergeMFS(r, fw, bw)
	if len(rbw) == 0 {
		return r
	}

	// apply recursion to the rest
	rem := MercifulStalin[T](buf.next(), rbw)
	r1 = mergeMFS(r1, r, rem)
	return append(a[:0], r1...)
}

func mergeMFS[T cmp.Ordered](dst, a, b []T) []T {
	n, m := len(a), len(b)
	var i, j int
	for i < n && j < m {
		if a[i] <= b[j] {
			dst = append(dst, a[i])
			i++
		} else {
			dst = append(dst, b[j])
			j++
		}
	}
	dst = append(dst, a[i:]...)
	dst = append(dst, b[j:]...)
	return dst
}

// MercifulStalinBuffer buffers intermediate data to avoid allocations during sorting.
type MercifulStalinBuffer[T cmp.Ordered] struct {
	fw_, rfw_, bw_, rbw_, r_, r1_ entry

	buf []T
	n   *MercifulStalinBuffer[T]
}

func (buf *MercifulStalinBuffer[T]) prealloc(n uint32) {
	if uint32(cap(buf.buf)) < n*6 {
		buf.buf = make([]T, n*6)
	}
	buf.buf = buf.buf[: n*6 : n*6]
	buf.fw_.Encode(0, n)
	buf.rfw_.Encode(n, n*2)
	buf.bw_.Encode(n*2, n*3)
	buf.rbw_.Encode(n*3, n*4)
	buf.r_.Encode(n*4, n*5)
	buf.r1_.Encode(n*5, n*6)
}

func (buf *MercifulStalinBuffer[T]) fw() []T {
	lo, hi := buf.fw_.Decode()
	return buf.buf[lo:hi][:0]
}

func (buf *MercifulStalinBuffer[T]) rfw() []T {
	lo, hi := buf.rfw_.Decode()
	return buf.buf[lo:hi][:0]
}

func (buf *MercifulStalinBuffer[T]) bw() []T {
	lo, hi := buf.bw_.Decode()
	return buf.buf[lo:hi][:0]
}

func (buf *MercifulStalinBuffer[T]) rbw() []T {
	lo, hi := buf.rbw_.Decode()
	return buf.buf[lo:hi][:0]
}

func (buf *MercifulStalinBuffer[T]) r() []T {
	lo, hi := buf.r_.Decode()
	return buf.buf[lo:hi][:0]
}

func (buf *MercifulStalinBuffer[T]) r1() []T {
	lo, hi := buf.r1_.Decode()
	return buf.buf[lo:hi][:0]
}

func (buf *MercifulStalinBuffer[T]) next() *MercifulStalinBuffer[T] {
	if buf.n == nil {
		buf.n = &MercifulStalinBuffer[T]{}
	}
	return buf.n
}

func (buf *MercifulStalinBuffer[T]) Reset() {
	buf.fw_.Reset()
	buf.rfw_.Reset()
	buf.bw_.Reset()
	buf.rbw_.Reset()
	buf.r_.Reset()
	buf.r1_.Reset()
	if buf.n != nil {
		buf.n.Reset()
	}
}
