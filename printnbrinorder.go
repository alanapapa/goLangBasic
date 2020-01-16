package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var arr []int
	count := 0
	if n > 0 {
		for i := n; i > 0; i = i / 10 {
			arr = append(arr, i%10)
		}
		for len := range arr {
			if len >= 0 {
				count++
			}
		}
		for i := 0; i < count; i++ {
			for j := 0; j < count; j++ {
				if arr[i] < arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
		for _, i := range arr {
			z01.PrintRune(rune(i + 48))
		}
	} else {
		z01.PrintRune('0')
	}
}
