package piscinego

func StringToIntSlice(str string) []int {
	tt := []int{}

	if str == "" {
		return nil
	}

	for _, p := range str {
		tt = append(tt, int(p))
	}
	return tt
}
