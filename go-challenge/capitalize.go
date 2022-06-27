package piscine

func Capitalize(s string) string {
	tabrunes := []rune(s)
	for i, lettre := range tabrunes {
		if isNumberOrAlph(lettre) {
			if i == 0 || isNumberOrAlph(tabrunes[i-1]) == false {
				if tabrunes[i] >= 'a' && tabrunes[i] <= 'z' {
					tabrunes[i] = lettre - 32
				}
			} else {
				if tabrunes[i] >= 'A' && tabrunes[i] <= 'Z' {
					tabrunes[i] = lettre + 32
				}
			}
		}
	}
	return string(tabrunes)
}

func isNumberOrAlph(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}
