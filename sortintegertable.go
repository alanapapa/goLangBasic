package piscine

func SortIntegerTable(table []int) {
	count := 0
	for index := range table {
		if index >= 0 {
			count++
		}
	}
	for i := 0; i <= 5; i++ {
		for j := count; j > count; j-- {
			table[i], table[j] = table[j], table[i]
		}
	}
}
