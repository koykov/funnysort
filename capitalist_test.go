package funnysort

import "testing"

func TestCapitalist(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run("", func(t *testing.T) {
			r := Capitalist(cpy(stages[i].source))
			if !assertEqual(r, stages[i].capitalist) {
				t.Fail()
			}
		})
	}
}

func BenchmarkCapitalist(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			var input []int
			for n := 0; n < b.N; n++ {
				input = append(input[:0], stages[i].source...)
				Capitalist(input)
			}
		})
	}
}
