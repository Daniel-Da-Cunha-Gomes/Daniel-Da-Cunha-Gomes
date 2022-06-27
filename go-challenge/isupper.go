package piscine

func IsUpper(s string) bool {
	count := 0
	countr := 0
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
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
