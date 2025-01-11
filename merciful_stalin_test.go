package funnysort

import "testing"

func TestMercifulStalinSort(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run("", func(t *testing.T) {
			var buf MercifulStalinBuffer[int]
			r := MercifulStalin(&buf, cpy(stages[i].source))
			if !assertEqual(r, stages[i].mercifulStalin) {
				t.Fail()
			}
		})
	}
}

func BenchmarkMercifulStalinSort(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		b.Run("", func(b *testing.B) {
			var (
				buf   MercifulStalinBuffer[int]
				input []int
			)
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				input = append(input[:0], stages[i].source...)
				buf.Reset()
				MercifulStalin(&buf, input)
			}
		})
	}
}
