package funnysort

import "testing"

func TestHitler(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run("", func(t *testing.T) {
			r := Hitler(cpy(stages[i].source), func(i int) bool { return i%4 == 0 || i%7 == 0 })
			if !assertEqual(r, stages[i].hitler) {
				t.Fail()
			}
		})
	}
}

func BenchmarkHitler(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			var input []int
			for n := 0; n < b.N; n++ {
				input = append(input[:0], stages[i].source...)
				Hitler(input, func(i int) bool { return i%4 == 0 || i%7 == 0 })
			}
		})
	}
}
