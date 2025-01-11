package funnysort

type stage struct {
	source         []int
	stalin         []int
	mercifulStalin []int
	liberal        []int
}

var stages = []stage{
	{
		source:         []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		stalin:         []int{1, 2, 10, 15, 30},
		mercifulStalin: []int{1, 2, 2, 3, 4, 6, 10, 15, 20, 30},
		liberal:        []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
	},
	{
		source:         []int{78, 33, 100, 61, 65, 72, 11, 66, 89, 3},
		stalin:         []int{78, 100},
		mercifulStalin: []int{3, 11, 33, 61, 65, 66, 72, 78, 89, 100},
		liberal:        []int{78, 33, 100, 61, 65, 72, 11, 66, 89, 3},
	},
	{
		source:         []int{2, 2, 3, 1, 10},
		stalin:         []int{2, 2, 3, 10},
		mercifulStalin: []int{1, 2, 2, 3, 10},
		liberal:        []int{2, 2, 3, 1, 10},
	},
	{
		source:         []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		stalin:         []int{1, 2, 10, 15, 30},
		mercifulStalin: []int{1, 2, 2, 3, 4, 6, 10, 15, 20, 30},
		liberal:        []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
	},
	{
		source:         []int{1, 2, 2, 3, 2, 5},
		stalin:         []int{1, 2, 2, 3, 5},
		mercifulStalin: []int{1, 2, 2, 2, 3, 5},
		liberal:        []int{1, 2, 2, 3, 2, 5},
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
