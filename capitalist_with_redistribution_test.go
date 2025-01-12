package funnysort

import "testing"

func TestCapitalistWithRedistribution(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run("", func(t *testing.T) {
			r := CapitalistWithRedistribution(cpy(stages[i].source))
			if !assertEqual(r, stages[i].capitalistRedist) {
				t.Fail()
			}
		})
	}
}

func BenchmarkCapitalistWithRedistribution(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			var input []int
			for n := 0; n < b.N; n++ {
				input = append(input[:0], stages[i].source...)
				CapitalistWithRedistribution(input)
			}
		})
	}
}
