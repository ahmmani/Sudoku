package piscine

import "github.com/01-edu/z01"

// Print the solved Sudoku grid
func PrintGrid() {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			z01.PrintRune(rune(Grid[row][col] + '0'))
			if col < 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
