package piscinego

/*/////////////////////////////////////////////////
						solution 1
/////////////////////////////////////////////////*/

func FindNextPrimePISCINE(nb int) int {
	if nb <= 1 {
		return 2
	}

	for nb > 1 {
		IsPrime := true

		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {

				IsPrime = false
				break
			}
		}
		if IsPrime {
			return nb
		}
		nb++
	}
	return nb
}

/*/////////////////////////////////////////////////
						solution 2
/////////////////////////////////////////////////*/

func FindNextPrime(nb int) int {
	for {
		if prime(nb) {
			return nb
		} else {
			nb++
		}
	}
}

func prime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
