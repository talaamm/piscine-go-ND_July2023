package piscinego

func Map(f func(int) bool, a []int) []bool {
	r := make([]bool, len(a))

	for i, v := range a {
		r[i] = f(v)
	}

	return r
}
