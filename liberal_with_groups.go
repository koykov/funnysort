package funnysort

// LiberalWithGroups also avoids sorting, but groups items according common attribute to achieve "no one felt isolated".
// Original order keeps inside the groups.
func LiberalWithGroups[T any](buf *LiberalBuffer[T], a []T, keyFunc func(T) int) []T {
	n := len(a)
	if n == 0 {
		return a
	}
	_ = a[n-1]
	for i := 0; i < n; i++ {
		buf.bufferize(keyFunc(a[i]), a[i])
	}
	return buf.appendTo(a[:0])
}

// LiberalBuffer buffers intermediate data to avoid allocations during sorting.
type LiberalBuffer[T any] struct {
	idx map[int]int
	buf [][]T
	off int
}

func (buf *LiberalBuffer[T]) bufferize(key int, value T) {
	if buf.idx == nil {
		buf.idx = make(map[int]int)
	}
	idx, ok := buf.idx[key]
	if !ok {
		if buf.off < len(buf.buf) {
			idx = buf.off
			buf.off++
		} else {
			buf.buf = append(buf.buf, []T{})
			buf.off++
			idx = buf.off - 1
			buf.idx[key] = idx
		}
	}
	buf.buf[idx] = append(buf.buf[idx], value)
}

func (buf *LiberalBuffer[T]) appendTo(dst []T) []T {
	for i := 0; i < len(buf.buf); i++ {
		dst = append(dst, buf.buf[i]...)
	}
	return dst
}

func (buf *LiberalBuffer[T]) Reset() {
	for k := range buf.buf {
		delete(buf.idx, k)
	}
	for i := 0; i < len(buf.buf); i++ {
		buf.buf[i] = buf.buf[i][:0]
	}
	buf.off = 0
}
