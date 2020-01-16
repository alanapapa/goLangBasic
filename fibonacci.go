package piscine

func Fibonacci(index int) int {
	if index == 0 || index == 1 {
		return index
	} else if index > 1 {
		return Fibonacci(index-1) + Fibonacci(index-2)
	} else {
		return -1
	}
}
