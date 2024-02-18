package piscinego

func BasicAtoi(s string) int {
	res := 0
	for _, r := range s {
		res = int(r-'0') + res*10
	}
	return res
}
