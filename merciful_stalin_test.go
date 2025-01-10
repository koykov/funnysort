package funnysort

import "testing"

var stagesMFS = []struct {
	a, r []int
}{
	{
		a: []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		r: []int{1, 2, 2, 3, 4, 6, 10, 15, 20, 30},
	},
	{
		a: []int{78, 33, 100, 61, 65, 72, 11, 66, 89, 3},
		r: []int{3, 11, 33, 61, 65, 66, 72, 78, 89, 100},
	},
	{
		a: []int{2, 2, 3, 1, 10},
		r: []int{1, 2, 2, 3, 10},
	},
	{
		a: []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		r: []int{1, 2, 2, 3, 4, 6, 10, 15, 20, 30},
	},
	{
		a: []int{1, 2, 2, 3, 2, 5},
		r: []int{1, 2, 2, 2, 3, 5},
	},
}

func TestMercifulStalinSort(t *testing.T) {
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
	for i := 0; i < len(stagesMFS); i++ {
		t.Run("", func(t *testing.T) {
			var buf MercifulStalinBuffer[int]
			r := MercifulStalin(&buf, stagesMFS[i].a)
			if !eq(r, stagesMFS[i].r) {
				t.Fail()
			}
		})
	}
}

func BenchmarkMercifulStalinSort(b *testing.B) {
	for i := 0; i < len(stagesMFS); i++ {
		b.Run("", func(b *testing.B) {
			var buf MercifulStalinBuffer[int]
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				buf.Reset()
				MercifulStalin(&buf, stagesMFS[i].a)
			}
		})
	}
}
