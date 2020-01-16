package piscine

func Capitalize(s string) string {
	array := []rune(s)
	count := 0
	for range array {
		count++
	}
	if array[0] >= 'a' && array[0] <= 'z' {
		array[0] -= 32
	}
	for i := 1; i < count; i++ {
		if array[i] >= 'A' && array[i] <= 'Z' {
			if (array[i-1] >= 'A' && array[i-1] <= 'Z') || (array[i-1] >= 'a' && array[i-1] <= 'z') || (array[i-1] >= '0' && array[i-1] <= '9') {
				array[i] += 32
			}
		}
		if array[i] >= 'a' && array[i] <= 'z' {
			if (array[i-1] >= 'A' && array[i-1] <= 'Z') || (array[i-1] >= 'a' && array[i-1] <= 'z') || (array[i-1] >= '0' && array[i-1] <= '9') {
				continue
			} else {
				array[i] -= 32
			}
		}
	}
	result := string(array)
	return result
}
