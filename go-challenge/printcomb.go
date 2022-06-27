package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 1; i < 999; i++ {
		a := i%10 + 48
		b := i/10%10 + 48
		c := i/100%10 + 48
		if a > b && b > c {
			z01.PrintRune(rune(c))
			z01.PrintRune(rune(b))
			z01.PrintRune(rune(a))
			if i == 789 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
		}
	}
}

/*package piscine

import "github.com/01-edu/z01"

func PrintComb() {
    var a, b, c int
    for i := 0; i < 789; i++ {
        c++
        if c == 10 {
            b++
            c = 0
        }
        if b == 10 {
            a++
            b = 0
        }
        if c > b && c > a && b > a {
            z01.PrintRune(rune(48 + a))
            z01.PrintRune(rune(48 + b))
            z01.PrintRune(rune(48 + c))
            z01.PrintRune(rune(44))
            z01.PrintRune(32)
        }
        if i == 788 {
            z01.PrintRune(rune(48 + a))
            z01.PrintRune(rune(48 + b))
            z01.PrintRune(rune(48 + c))
            z01.PrintRune(rune('\n'))
            break
        }

    }
}
*/
