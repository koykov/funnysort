package funnysort

type entry uint64

func (e *entry) Encode(lo, hi uint32) entry {
	*e = entry(lo)<<32 | entry(hi)
	return *e
}

func (e *entry) Decode() (lo, hi uint32) {
	lo = uint32(*e >> 32)
	hi = uint32(*e)
	return
}

func (e *entry) Reset() { *e = 0 }
