package funnysort

import "testing"

var stagesStalin = []struct {
	a, r []int
}{
	{
		a: []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		r: []int{1, 2, 10, 15, 30},
	},
	{
		a: []int{78, 33, 100, 61, 65, 72, 11, 66, 89, 3},
		r: []int{78, 100},
	},
	{
		a: []int{2, 2, 3, 1, 10},
		r: []int{2, 2, 3, 10},
	},
	{
		a: []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		r: []int{1, 2, 10, 15, 30},
	},
	{
		a: []int{1, 2, 2, 3, 2, 5},
		r: []int{1, 2, 2, 3, 5},
	},
}

func TestStalin(t *testing.T) {
	eq := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(stagesStalin); i++ {
		t.Run("", func(t *testing.T) {
			r := Stalin(stagesStalin[i].a)
			if !eq(r, stagesStalin[i].r) {
				t.Fail()
			}
		})
	}
}

func BenchmarkStalin(b *testing.B) {
	for i := 0; i < len(stagesStalin); i++ {
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				Stalin(stagesStalin[i].a)
			}
		})
	}
}
