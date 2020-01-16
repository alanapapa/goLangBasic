package piscine

func StrRev(s string) string {
	var test string
	for _, element := range s {
		test = string(element) + test
	}
	return string(test)
}
