package piscine

func TrimAtoi(s string) int {
	var tableau []int
	result := 0
	miniindex := -1
	firstindex := 0
	index := 0
	tableaucompteur := 0
	for _, rune := range s {
		if rune == '-' {
			miniindex = index
		}
		if isDigit(rune) {
			if firstindex == 0 {
				firstindex = index
			}
			tableau = append(tableau, int(rune-'0'))
		}
		index++
	}
	for count := range tableau {
		tableaucompteur = count + 1
	}
	for i := 0; i < tableaucompteur; i++ {
		result = result*10 + tableau[i]
	}
	if miniindex < firstindex && miniindex != -1 {
		result = result * -1
	}
	return result
}

func isDigit(digit rune) bool {
	for a := '0'; a <= '9'; a++ {
		if digit == a {
			return true
		}
	}
	return false
}
