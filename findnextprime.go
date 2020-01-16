package piscine

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	for i := 5; i*i <= nb; i = i + 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(Nb int) int {
	if Nb <= 1 {
		return 2
	}

	var prime int = Nb
	var found bool = false

	if isPrime(prime) != found {
		return Nb
	}

	for !found {
		prime++
		if isPrime(prime) {
			found = true
		}
	}
	return prime
}
