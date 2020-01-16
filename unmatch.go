package piscine

func Unmatch(arr []int) int {
	count := 0
	for range arr {
		count++
	}
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if arr[i] == arr[j] {
				arr[i] = -1
				arr[j] = -1
				break
			}
		}
	}
	for k := 0; k < count; k++ {
		if arr[k] != -1 {
			return arr[k]
		}
	}
	return -1
}
