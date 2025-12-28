package quads

import (
	"fmt"

	"github.com/01-edu/z01"
)

// QuadA prints a rectangle with:
// - 'o' at all four corners
// - '-' on top and bottom edges
// - '|' on left and right edges
// - ' ' (space) in the interior

func QuadA(x, y int) {
	// Validate input
	if x <= 0 || y <= 0 {
		fmt.Printf("Invalid input parameters\n")
		return
	}

	// Outer loop: iterate through each row (0 to y-1)
	for row := 0; row < y; row++ {
		// Inner loop: iterate through each column (0 to x-1)
		for col := 0; col < x; col++ {
			// CONDITION 1: All four corners (same character)
			// Top-left: (0, 0), Top-right: (0, x-1)
			// Bottom-left: (y-1, 0), Bottom-right: (y-1, x-1)
			if (row == 0 || row == y-1) && (col == 0 || col == x-1) {
				z01.PrintRune('o')

				// CONDITION 2: Top and bottom edges (not corners)
				// Top row: row == 0, Bottom row: row == y-1
			} else if row == 0 || row == y-1 {
				z01.PrintRune('-')

				// CONDITION 3: Left and right edges (not corners, not top/bottom)
				// Left column: col == 0, Right column: col == x-1
			} else if col == 0 || col == x-1 {
				z01.PrintRune('|')

				// CONDITION 4: Interior (everything else)
				// Not on any edge or corner
			} else {
				z01.PrintRune(' ')
			}
		}

		// End of row: print newline before next row
		z01.PrintRune('\n')
	}
}
