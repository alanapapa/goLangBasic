package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {

	flag := true
	flag1 := true

	for i := 0; i < Count(tab)-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			flag = false
		}
	}
	for i := 0; i < Count(tab)-1; i++ {
		if f(tab[i], tab[i+1]) < 0 {
			flag1 = false
		}
	}
	return flag || flag1

}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func Count(arg []int) int {
	count := 0
	for range arg {
		count++
	}
	return count
}
