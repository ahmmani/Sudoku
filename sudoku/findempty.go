package piscine

// Find an empty cell (0 value), returns (row, col, found)
func findEmpty() (int, int, bool) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if Grid[row][col] == 0 {
				return row, col, true
			}
		}
	}
	return -1, -1, false // No empty cell found
}
