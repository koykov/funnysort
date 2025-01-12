package funnysort

import "testing"

func TestLateSocialist(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run("", func(t *testing.T) {
			r := LateSocialist(cpy(stages[i].source))
			if !assertEqual(r, stages[i].socialistLate) {
				t.Log(r)
				t.Fail()
			}
		})
	}
}

func BenchmarkLateSocialist(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			var input []int
			for n := 0; n < b.N; n++ {
				input = append(input[:0], stages[i].source...)
				LateSocialist(input)
			}
		})
	}
}
