package piscine

func MakeRange(min, max int) []int {
	var tableau []int
	if max > min {
		tableau = make([]int, max-min)
		for i := range tableau {
			tableau[i] = i + min
		}
	}
	return tableau
}
