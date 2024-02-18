package piscinego

func Join(strs []string, sep string) string {
	res := ""
	for i, r := range strs {
		if i == len(strs)-1 {
			res += r
		} else {
			res += r + sep
		}
	}
	return res
}
