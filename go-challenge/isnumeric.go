package piscine

func IsNumeric(s string) bool {
	count := 0
	countr := 0
	for _, i := range s {
		if i >= '0' && i <= '9' {
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
