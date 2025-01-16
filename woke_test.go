package funnysort

import "testing"

func TestWoke(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run("", func(t *testing.T) {
			r := Woke(cpy(stages[i].source))
			if !assertEqual(r, stages[i].woke) {
				t.Log(r)
				t.Fail()
			}
		})
	}
}

func BenchmarkWoke(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			var input []int
			for n := 0; n < b.N; n++ {
				input = append(input[:0], stages[i].source...)
				Woke(input)
			}
		})
	}
}
