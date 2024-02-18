package piscinego

func BasicAtoi2(s string) int {
	res := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			res = int(r-'0') + res*10
		} else {
			return 0
		}
	}
	return res
}
