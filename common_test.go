package funnysort

type stage struct {
	source           []int
	stalin           []int
	mercifulStalin   []int
	liberal          []int
	liberalGroups    []int
	capitalist       []int
	capitalistRedist []int
	capitalistHrchy  []int
}

var stages = []stage{
	{
		source:           []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		stalin:           []int{1, 2, 10, 15, 30},
		mercifulStalin:   []int{1, 2, 2, 3, 4, 6, 10, 15, 20, 30},
		liberal:          []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		liberalGroups:    []int{1, 3, 15, 2, 10, 2, 4, 6, 30, 20},
		capitalist:       []int{30, 20, 15, 10, 6, 4, 3, 2, 2, 1},
		capitalistRedist: []int{40, 17, 12, 8, 5, 3, 2, 2, 1, 0},
		capitalistHrchy:  []int{45, 14, 9, 6, 4, 4, 3, 3, 3, 2},
	},
	{
		source:           []int{78, 33, 100, 61, 65, 72, 11, 66, 89, 3},
		stalin:           []int{78, 100},
		mercifulStalin:   []int{3, 11, 33, 61, 65, 66, 72, 78, 89, 100},
		liberal:          []int{78, 33, 100, 61, 65, 72, 11, 66, 89, 3},
		liberalGroups:    []int{78, 100, 72, 66, 33, 61, 65, 11, 89, 3},
		capitalist:       []int{100, 89, 78, 72, 66, 65, 61, 33, 11, 3},
		capitalistRedist: []int{144, 83, 75, 69, 65, 62, 46, 21, 6, 1},
		capitalistHrchy:  []int{115, 83, 72, 66, 66, 65, 61, 34, 12, 4},
	},
	{
		source:           []int{2, 2, 3, 1, 10},
		stalin:           []int{2, 2, 3, 10},
		mercifulStalin:   []int{1, 2, 2, 3, 10},
		liberal:          []int{2, 2, 3, 1, 10},
		liberalGroups:    []int{2, 2, 10, 3, 1},
		capitalist:       []int{10, 3, 2, 2, 1},
		capitalistRedist: []int{11, 2, 2, 1, 0},
		capitalistHrchy:  []int{20, 3, 2, -3, -4},
	},
	{
		source:           []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		stalin:           []int{1, 2, 10, 15, 30},
		mercifulStalin:   []int{1, 2, 2, 3, 4, 6, 10, 15, 20, 30},
		liberal:          []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		liberalGroups:    []int{1, 3, 15, 2, 10, 2, 4, 6, 30, 20},
		capitalist:       []int{30, 20, 15, 10, 6, 4, 3, 2, 2, 1},
		capitalistRedist: []int{40, 17, 12, 8, 5, 3, 2, 2, 1, 0},
		capitalistHrchy:  []int{45, 14, 9, 6, 4, 4, 3, 3, 3, 2},
	},
	{
		source:           []int{1, 2, 2, 3, 2, 5},
		stalin:           []int{1, 2, 2, 3, 5},
		mercifulStalin:   []int{1, 2, 2, 2, 3, 5},
		liberal:          []int{1, 2, 2, 3, 2, 5},
		liberalGroups:    []int{1, 3, 5, 2, 2, 2},
		capitalist:       []int{5, 3, 2, 2, 2, 1},
		capitalistRedist: []int{6, 2, 2, 2, 1, 0},
		capitalistHrchy:  []int{15, 3, 2, 2, -3, -4},
	},
}

func cpy(a []int) (r []int) {
	r = append(r, a...)
	return
}

func assertEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
