package piscine

func IsLower(s string) bool {
	count := 0
	countr := 0
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
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
