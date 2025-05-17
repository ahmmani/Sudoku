package piscine

func IsValidGrid() bool {
	// Check rows
	for row := 0; row < 9; row++ {
		seen := [10]bool{} // 1-indexed, ignore index 0
		for col := 0; col < 9; col++ {
			num := Grid[row][col]
			if num != 0 {
				if seen[num] {
					return false // Duplicate in row
				}
				seen[num] = true
			}
		}
	}

	// Check columns
	for col := 0; col < 9; col++ {
		seen := [10]bool{}
		for row := 0; row < 9; row++ {
			num := Grid[row][col]
			if num != 0 {
				if seen[num] {
					return false // Duplicate in column
				}
				seen[num] = true
			}
		}
	}

	// Check 3x3 boxes
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			seen := [10]bool{}
			for row := boxRow * 3; row < (boxRow+1)*3; row++ {
				for col := boxCol * 3; col < (boxCol+1)*3; col++ {
					num := Grid[row][col]
					if num != 0 {
						if seen[num] {
							return false // Duplicate in box
						}
						seen[num] = true
					}
				}
			}
		}
	}

	return true
}

// Check if placing 'num' at position (row, col) is valid
func IsValid(row, col, num int) bool {
	// Check row
	for i := 0; i < 9; i++ {
		if Grid[row][i] == num {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if Grid[i][col] == num {
			return false
		}
	}

	// Check 3x3 box
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if Grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}
