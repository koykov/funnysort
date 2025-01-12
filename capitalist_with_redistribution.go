package funnysort

import (
	"math"
)

func CapitalistWithRedistribution[T OrderedNoString](a []T) []T {
	n := len(a)
	if n == 0 {
		return a
	}
	a = Capitalist[T](a)
	for i := 0; i < n-1; i++ {
		switch x := any(a[i]).(type) { // keep in sync with OrderedNoString type
		case int:
			y := any(a[i+1]).(int)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		case int8:
			y := any(a[i+1]).(int8)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		case int16:
			y := any(a[i+1]).(int16)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		case int32:
			y := any(a[i+1]).(int32)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		case int64:
			y := any(a[i+1]).(int64)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		case uint:
			y := any(a[i+1]).(uint)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		case uint8:
			y := any(a[i+1]).(uint8)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		case uint16:
			y := any(a[i+1]).(uint16)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		case uint32:
			y := any(a[i+1]).(uint32)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		case uint64:
			y := any(a[i+1]).(uint64)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		case float32:
			y := any(a[i+1]).(float32)
			a[i] = T(math.Float32bits(x + y/2))
			a[i+1] = T(math.Float32bits(y / 2))
		case float64:
			y := any(a[i+1]).(float64)
			a[i] = T(math.Float64bits(x + y/2))
			a[i+1] = T(math.Float64bits(y / 2))
		case uintptr:
			y := any(a[i+1]).(uintptr)
			a[i] = T(x + y/2)
			a[i+1] = T(y / 2)
		default:
			// case string: sorry, no support strings even it includes to cmp.Ordered
		}
	}
	return a
}
