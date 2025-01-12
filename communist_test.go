package funnysort

import "testing"

func TestCommunist(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run("", func(t *testing.T) {
			r := Communist(cpy(stages[i].source))
			if !assertEqual(r, stages[i].communist) {
				t.Fail()
			}
		})
	}
}

func BenchmarkCommunist(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			var input []int
			for n := 0; n < b.N; n++ {
				input = append(input[:0], stages[i].source...)
				Communist(input)
			}
		})
	}
}
