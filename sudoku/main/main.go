package main

import (
	"os"

	"piscine"
)

// Main function that handles program execution
func main() {
	// Check if we have 9 arguments (9 rows of the Sudoku grid)
	if len(os.Args) != 10 {
		piscine.PrintError()
		return
	}

	// Parse the input and initialize the grid
	if !piscine.ParseInput() {
		piscine.PrintError()
		return
	}

	// Check if the initial grid is valid
	if !piscine.IsValidGrid() {
		piscine.PrintError()
		return
	}

	// Solve the Sudoku
	if piscine.Solve() {
		// Print the solution if found
		piscine.PrintGrid()
	} else {
		// Print error if no solution exists
		piscine.PrintError()
	}
}
