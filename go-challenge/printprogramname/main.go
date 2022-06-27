package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, mots := range os.Args[0][2:] {
		z01.PrintRune(mots)
	}
	z01.PrintRune('\n')
}
