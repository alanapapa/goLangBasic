package piscine

func SplitWhiteSpaces(str string) []string {
	l := 0
	ns := -1
	str1 := []rune(str)

	for k, i := range str {
		if k != 0 {
			if i == 9 || i == 10 || i == 32 {
				if str[k-1] < 9 || str[k-1] > 10 && str[k-1] < 32 || str[k-1] > 32 {
					l++
				}
			}
		}
	}

	for range str {
		ns++
	}

	if str1[ns] == 9 || str1[ns] == 10 || str1[ns] == 32 {
		l--
	}

	res := make([]string, l+1)
	m := 0
	n := 0

	for i, a := range str1 {
		if i != 0 && i < ns {
			if a == 9 || a == 10 || a == 32 {
				if str1[i-1] == 9 || str1[i-1] == 10 || str1[i-1] == 32 {
					n = i + 1
				} else if str1[i-1] < 9 || str1[i-1] > 10 && str1[i-1] < 32 || str1[i-1] > 32 {
					res[m] = str[n:i]
					m++
					n = i + 1
				}
			}
		} else if i == ns {
			if a < 9 || a > 10 && a < 32 || a > 32 {
				res[m] = str[n:]
			}
		}
	}

	return res
}
