package piscinego

/*/////////////////////////////////////////////////
						solution 1
/////////////////////////////////////////////////*/

func AtoiJenan(s string) int {
	neg := false
	pos := false

	res := 0
	for _, r := range s {
		if s[0] == '-' && !neg && !pos {
			neg = true
		} else if s[0] == '+' && !neg && !pos {
			pos = true
		} else if r >= '0' && r <= '9' {
			res = int(r-'0') + res*10
		} else {
			return 0
		}
	}
	if neg {
		res = -res
	}

	return res
}

/*/////////////////////////////////////////////////
						solution 2
/////////////////////////////////////////////////*/

func Atoi(s string) int {
	wiwi := 0

	if s[0] == '-' {
		for _, r := range s[1:] {
			if r >= '0' && r <= '9' {
				wiwi = wiwi*10 + int(r-48)
			} else {
				return 0
			}
		}
		return wiwi * -1
	} else if s[0] == '+' {
		for _, r := range s[1:] {
			if r >= '0' && r <= '9' {
				wiwi = wiwi*10 + int(r-48)
			} else {
				return 0
			}
		}
		return wiwi
	} else {
		for _, r := range s {
			if r >= '0' && r <= '9' {
				wiwi = wiwi*10 + int(r-48)
			} else {
				return 0
			}
		}
	}
	return wiwi
}
