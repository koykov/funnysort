package funnysort

import "testing"

func TestLenin(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run("", func(t *testing.T) {
			r := Lenin(cpy(stages[i].source), 3)
			if !assertEqual(r, stages[i].lenin) {
				t.Log(r)
				t.Fail()
			}
		})
	}
}

func BenchmarkLenin(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			var input []int
			for n := 0; n < b.N; n++ {
				input = append(input[:0], stages[i].source...)
				Lenin(input, 3)
			}
		})
	}
}
