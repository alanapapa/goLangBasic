package piscine

func Map(f func(int) bool, arr []int) []bool {
	len := 0
	for range arr {
		len++
	}
	result := make([]bool, len)
	for i := range arr {
		result[i] = f(arr[i])
	}
	return result
}
