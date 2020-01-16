package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb < 21 {
		for i := 1; i <= nb; i++ {
			result *= i
		}
		return result
	} else {
		return 0
	}
}
