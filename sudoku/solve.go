package piscine

// Solve the Sudoku puzzle using backtracking
func Solve() bool {
	// Find an empty cell
	row, col, found := findEmpty()
	if !found {
		return true // Puzzle is solved
	}

	// Try digits 1-9 for this empty cell
	for num := 1; num <= 9; num++ {
		if IsValid(row, col, num) {
			// Place the number
			Grid[row][col] = num

			// Recursively solve the rest of the puzzle
			if Solve() {
				return true
			}

			// If placing this number didn't work, backtrack
			Grid[row][col] = 0
		}
	}

	return false // Trigger backtracking
}
