package funnysort

import "testing"

func TestRobespierre(t *testing.T) {
	t.Run("no siblings", func(t *testing.T) {
		for i := 0; i < len(stages); i++ {
			t.Run("", func(t *testing.T) {
				r := Robespierre(cpy(stages[i].source), RobespierreTerrorPolicyExample_HasNoSiblings[int])
				if !assertEqual(r, stages[i].robespierreNoSl) {
					t.Log(r)
					t.Fail()
				}
			})
		}
	})
	t.Run("has siblings", func(t *testing.T) {
		for i := 0; i < len(stages); i++ {
			t.Run("", func(t *testing.T) {
				r := Robespierre(cpy(stages[i].source), RobespierreTerrorPolicyExample_HasSiblings[int])
				if !assertEqual(r, stages[i].robespierreSl) {
					t.Log(r)
					t.Fail()
				}
			})
		}
	})
}

func BenchmarkRobespierre(b *testing.B) {
	b.Run("no siblings", func(b *testing.B) {
		for i := 0; i < len(stages); i++ {
			b.Run("", func(b *testing.B) {
				b.ReportAllocs()
				var input []int
				for n := 0; n < b.N; n++ {
					input = append(input[:0], stages[i].source...)
					Robespierre(input, RobespierreTerrorPolicyExample_HasNoSiblings[int])
				}
			})
		}
	})
	b.Run("has siblings", func(b *testing.B) {
		for i := 0; i < len(stages); i++ {
			b.Run("", func(b *testing.B) {
				b.ReportAllocs()
				var input []int
				for n := 0; n < b.N; n++ {
					input = append(input[:0], stages[i].source...)
					Robespierre(input, RobespierreTerrorPolicyExample_HasSiblings[int])
				}
			})
		}
	})
}
