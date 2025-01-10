package funnysort

import (
	"cmp"
)

type MercifulStalinBuffer[T cmp.Ordered] struct {
	fw_, rfw_, bw_, rbw_, r_, r1_ entryMFB

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

type entryMFB uint64

func (e *entryMFB) Encode(lo, hi uint32) entryMFB {
	*e = entryMFB(lo)<<32 | entryMFB(hi)
	return *e
}

func (e *entryMFB) Decode() (lo, hi uint32) {
	lo = uint32(*e >> 32)
	hi = uint32(*e)
	return
}

func (e *entryMFB) Reset() { *e = 0 }
