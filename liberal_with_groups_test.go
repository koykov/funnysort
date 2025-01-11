package funnysort

import "testing"

func TestLiberalWithGroups(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run("", func(t *testing.T) {
			var buf LiberalBuffer[int]
			r := LiberalWithGroups(&buf, cpy(stages[i].source), func(x int) int { return x % 2 })
			if !assertEqual(r, stages[i].liberalGroups) {
				t.Fail()
			}
		})
	}
}

func BenchmarkLiberalWithGroups(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			var (
				buf   LiberalBuffer[int]
				input []int
			)
			for n := 0; n < b.N; n++ {
				input = append(input[:0], stages[i].source...)
				buf.Reset()
				LiberalWithGroups(&buf, input, func(x int) int { return x % 2 })
			}
		})
	}
}
