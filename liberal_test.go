package funnysort

import "testing"

func TestLiberal(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		t.Run("", func(t *testing.T) {
			r := Liberal(cpy(stages[i].source))
			if !assertEqual(r, stages[i].liberal) {
				t.Fail()
			}
		})
	}
}

func BenchmarkLiberal(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		b.Run("", func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				Liberal(stages[i].source)
			}
		})
	}
}
