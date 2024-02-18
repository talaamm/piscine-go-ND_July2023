package piscinego

/*/////////////////////////////////////////////////
						solution 1
/////////////////////////////////////////////////*/

func StrRev(s string) string {
	result := ""
	for _, r := range s {
		result = string(r) + result
	}
	return result
}

/*/////////////////////////////////////////////////
						solution 2
/////////////////////////////////////////////////*/

func strrev3(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

/*/////////////////////////////////////////////////
						solution 3
/////////////////////////////////////////////////*/

func StrRev2(s string) string {
	r := len(s) - 1
	j := 0
	newstr := make([]rune, len(s))
	for i := r; i >= 0; i-- {
		newstr[j] = rune(s[i])
		j++
	}
	return string(newstr)
}
