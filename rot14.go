package piscinego

/*/////////////////////////////////////////////////
						solution 1
/////////////////////////////////////////////////*/

func Rot14(s string) string {
	var result string

	for _, r := range s {
		if (r >= 'a' && r <= 'l') || (r >= 'A' && r <= 'L') {
			r = r + 14
		} else if (r >= 'm' && r <= 'z') || (r >= 'M' && r <= 'Z') {
			r = r - 12
		}
		result += string(r)
	}
	return result
}

/*/////////////////////////////////////////////////
						solution 2
/////////////////////////////////////////////////*/

func rot14(s string) string {
	res := ""
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			r = (r-'a'+14)%26 + 'a'
		case r >= 'A' && r <= 'Z':
			r = (r-'A'+14)%26 + 'A'
		}
		res += string(r)
	}
	return res
}
