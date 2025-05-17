package piscine

import "github.com/01-edu/z01"

// Print "Error" message
func PrintError() {
	for _, ch := range "Error" {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
