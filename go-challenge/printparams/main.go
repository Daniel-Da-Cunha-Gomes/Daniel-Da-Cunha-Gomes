package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for t := 1; t < len(os.Args); t++ {
		for _, mots := range os.Args[t] {
			z01.PrintRune(mots)
		}
		z01.PrintRune('\n')
	}
}
