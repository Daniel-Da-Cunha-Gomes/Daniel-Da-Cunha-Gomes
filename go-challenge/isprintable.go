package piscine

func IsPrintable(s string) bool {
	count := 0
	countr := 0
	for _, i := range s {
		if i >= 32 && i <= 126 {
			count++
		}
		countr++
	}
	if count == countr {
		return true
	} else {
		return false
	}
}
