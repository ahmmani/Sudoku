package piscine

import "os"

// Parse command line arguments into our grid
func ParseInput() bool {
	for i := 1; i <= 9; i++ {
		row := os.Args[i]

		// Each row must be exactly 9 characters
		if len(row) != 9 {
			return false
		}

		for j := 0; j < 9; j++ {
			if row[j] == '.' {
				Grid[i-1][j] = 0 // Empty cell
			} else if row[j] >= '1' && row[j] <= '9' {
				Grid[i-1][j] = int(row[j] - '0') // Convert rune to int
			} else {
				return false // Invalid character
			}
		}
	}
	return true
}
