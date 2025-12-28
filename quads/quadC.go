package quads

import (
	"fmt"

	"github.com/01-edu/z01"
)

// QuadC Pattern:
//
// 5x3 example:
//   ABBBA
//   B   B
//   CBBBC
//
// Character mapping:
//   'A' = Top-left corner, Top-right corner
//   'C' = Bottom-left corner, Bottom-right corner
//   'B' = All other edges
//   ' ' = Interior spaces

// QuadC prints a rectangle with:
// - 'A' at top-left and top-right corners
// - 'C' at bottom-left and bottom-right corners
// - 'B' on all other edges
// - ' ' (space) in the interior

func QuadC(x, y int) {
	// Validate input
	if x <= 0 || y <= 0 {
		fmt.Printf("Invalid input parameters\n")
		return
	}

	// Outer loop: iterate through each row (0 to y-1)
	for row := 0; row < y; row++ {
		// Inner loop: iterate through each column (0 to x-1)
		for col := 0; col < x; col++ {

			// CONDITION 1: Top-left corner
			// Position: (0, 0)
			if row == 0 && col == 0 {
				z01.PrintRune('A')

			// CONDITION 2: Top-right corner
			// Position: (0, x-1)
			} else if row == 0 && col == x-1 {
				z01.PrintRune('A')

			// CONDITION 3: Bottom-left corner
			// Position: (y-1, 0)
			} else if row == y-1 && col == 0 {
				z01.PrintRune('C')

			// CONDITION 4: Bottom-right corner
			// Position: (y-1, x-1)
			} else if row == y-1 && col == x-1 {
				z01.PrintRune('C')

			// CONDITION 5: All other edges (top, bottom, left, right)
			// Not a corner, but on the perimeter
			} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
				z01.PrintRune('B')

			// CONDITION 6: Interior (everything else)
			// Not on any edge or corner
			} else {
				z01.PrintRune(' ')
			}
		}

		// End of row: print newline before next row
		z01.PrintRune('\n')
	}
}