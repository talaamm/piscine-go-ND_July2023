package piscinego

func Unmatch(a []int) int {
	var i int

	for _, r := range a {
		i = 0
		for _, v := range a {
			if v == r {
				i++
			}
		}
		if i%2 != 0 {
			return r
		}
	}
	return -1
}
