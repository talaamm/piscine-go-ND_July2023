package piscinego

func MakeRange(min, max int) []int {
	if min < max {
		s := max - min
		a := make([]int, s)

		for i := 0; i <= s-1; i++ {
			a[i] = min + i
		}
		return a
	} else {
		return nil
	}
}
