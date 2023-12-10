package main

import (
	"strconv"
	"strings"
)

/***************************************************
        GENERAL FUNCTIONS FOR BOTH PARTS
***************************************************/

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// isANumber checks if a string is a number
func isANumber(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil && char != "."
}

// isSpecial checks if a string is a special character
func isSpecial(char string) bool {
	return !isANumber(char) && char != "."
}

/***************************************************
		MAIN FUNCTIONS FOR BOTH PARTS


					 NOTE
----------------------------------------------------
> The isAdjacent and isAdjacentToNumber functions are
  based on the following matrix:
		A B C
		D * E
		F G H
  where * is the current number being evaluated.

***************************************************/

// checkEngine checks if an engine (*) is adjacent to a number and returns the coordinates of the numbers
func checkEngine(matrix [][]string, x int, y int) []string {
	var coordinates []string

	if x == 0 {
		if y == 0 {
			// CHECK E BECAUSE IS THE ONLY ONE THAT CAN BE ADJACENT
			if isAdjacentToNumber("E", matrix, x, y) {
				coordinates = append(coordinates, "E")
			}

			if isAdjacentToNumber("G", matrix, x, y) {
				coordinates = append(coordinates, "G")
			} else if isAdjacentToNumber("H", matrix, x, y) {
				coordinates = append(coordinates, "H")
			}

		} else if y == len(matrix[x])-1 {
			if isAdjacentToNumber("D", matrix, x, y) {
				coordinates = append(coordinates, "D")
			}

			if isAdjacentToNumber("F", matrix, x, y) {
				coordinates = append(coordinates, "F")
			} else if isAdjacentToNumber("G", matrix, x, y) {
				coordinates = append(coordinates, "G")
			}
		} else {
			if isAdjacentToNumber("D", matrix, x, y) {
				coordinates = append(coordinates, "D")
			}

			if isAdjacentToNumber("E", matrix, x, y) {
				coordinates = append(coordinates, "E")
			}

			// check bottom
			if !isAdjacentToNumber("G", matrix, x, y) {
				if isAdjacentToNumber("F", matrix, x, y) {
					coordinates = append(coordinates, "F")
				}

				if isAdjacentToNumber("H", matrix, x, y) {
					coordinates = append(coordinates, "H")
				}
			} else if isAdjacentToNumber("F", matrix, x, y) {
				coordinates = append(coordinates, "F")
			} else if isAdjacentToNumber("G", matrix, x, y) {
				coordinates = append(coordinates, "G")
			} else if isAdjacentToNumber("H", matrix, x, y) {
				coordinates = append(coordinates, "H")
			}
		}
	} else if x == len(matrix)-1 {
		if y == 0 {
			// check value at right
			if isAdjacentToNumber("E", matrix, x, y) {
				coordinates = append(coordinates, "E")
			}

			// check value at top
			if isAdjacentToNumber("B", matrix, x, y) {
				coordinates = append(coordinates, "B")
			} else if isAdjacentToNumber("C", matrix, x, y) {
				coordinates = append(coordinates, "C")
			}
		} else if y == len(matrix[x])-1 {
			// check value at left
			if isAdjacentToNumber("D", matrix, x, y) {
				coordinates = append(coordinates, "D")
			}

			// check value at top
			if isAdjacentToNumber("A", matrix, x, y) {
				coordinates = append(coordinates, "A")
			} else if isAdjacentToNumber("B", matrix, x, y) {
				coordinates = append(coordinates, "B")
			}
		} else {
			// check valus side by side
			if isAdjacentToNumber("D", matrix, x, y) {
				coordinates = append(coordinates, "D")
			}

			if isAdjacentToNumber("E", matrix, x, y) {
				coordinates = append(coordinates, "E")
			}

			// check top
			if !isAdjacentToNumber("B", matrix, x, y) {
				if isAdjacentToNumber("A", matrix, x, y) {
					coordinates = append(coordinates, "A")
				}

				if isAdjacentToNumber("C", matrix, x, y) {
					coordinates = append(coordinates, "C")
				}
			} else if isAdjacentToNumber("A", matrix, x, y) {
				coordinates = append(coordinates, "A")
			} else if isAdjacentToNumber("B", matrix, x, y) {
				coordinates = append(coordinates, "B")
			} else if isAdjacentToNumber("C", matrix, x, y) {
				coordinates = append(coordinates, "C")
			}
		}
	} else {
		if y == 0 {
			// check value at right
			if isAdjacentToNumber("E", matrix, x, y) {
				coordinates = append(coordinates, "E")
			}

			// check value at top
			if isAdjacentToNumber("B", matrix, x, y) {
				coordinates = append(coordinates, "B")
			} else if isAdjacentToNumber("C", matrix, x, y) {
				coordinates = append(coordinates, "C")
			}

			// check value at bottom
			if isAdjacentToNumber("G", matrix, x, y) {
				coordinates = append(coordinates, "G")
			} else if isAdjacentToNumber("H", matrix, x, y) {
				coordinates = append(coordinates, "H")
			}
		} else if y == len(matrix[x])-1 {
			// check value at left
			if isAdjacentToNumber("D", matrix, x, y) {
				coordinates = append(coordinates, "D")
			}

			// check value at top
			if isAdjacentToNumber("A", matrix, x, y) {
				coordinates = append(coordinates, "A")
			} else if isAdjacentToNumber("B", matrix, x, y) {
				coordinates = append(coordinates, "B")
			}

			// check value at bottom
			if isAdjacentToNumber("F", matrix, x, y) {
				coordinates = append(coordinates, "F")
			} else if isAdjacentToNumber("G", matrix, x, y) {
				coordinates = append(coordinates, "G")
			}
		} else {
			// check left
			if isAdjacentToNumber("D", matrix, x, y) {
				coordinates = append(coordinates, "D")
			}

			// check right
			if isAdjacentToNumber("E", matrix, x, y) {
				coordinates = append(coordinates, "E")
			}

			// check top
			if !isAdjacentToNumber("B", matrix, x, y) {
				if isAdjacentToNumber("A", matrix, x, y) {
					coordinates = append(coordinates, "A")
				}

				if isAdjacentToNumber("C", matrix, x, y) {
					coordinates = append(coordinates, "C")
				}
			} else if isAdjacentToNumber("A", matrix, x, y) {
				coordinates = append(coordinates, "A")
			} else if isAdjacentToNumber("B", matrix, x, y) {
				coordinates = append(coordinates, "B")
			} else if isAdjacentToNumber("C", matrix, x, y) {
				coordinates = append(coordinates, "C")
			}

			// check bottom
			if !isAdjacentToNumber("G", matrix, x, y) {
				if isAdjacentToNumber("F", matrix, x, y) {
					coordinates = append(coordinates, "F")
				}

				if isAdjacentToNumber("H", matrix, x, y) {
					coordinates = append(coordinates, "H")
				}
			} else if isAdjacentToNumber("F", matrix, x, y) {
				coordinates = append(coordinates, "F")
			} else if isAdjacentToNumber("G", matrix, x, y) {
				coordinates = append(coordinates, "G")
			} else if isAdjacentToNumber("H", matrix, x, y) {
				coordinates = append(coordinates, "H")
			}
		}
	}

	return coordinates
}

// isAdjacent checks if a coordinate is adjacent to a special character
func isAdjacent(coordinate string, matrix [][]string, x int, y int) bool {
	if strings.Contains(coordinate, "A") {
		if isSpecial(matrix[x-1][y-1]) {

			return true

		}
	}

	if strings.Contains(coordinate, "B") {
		if isSpecial(matrix[x-1][y]) {

			return true

		}
	}

	if strings.Contains(coordinate, "C") {
		if isSpecial(matrix[x-1][y+1]) {

			return true

		}
	}

	if strings.Contains(coordinate, "D") {
		if isSpecial(matrix[x][y-1]) {

			return true

		}
	}

	if strings.Contains(coordinate, "E") {
		if isSpecial(matrix[x][y+1]) {

			return true

		}
	}

	if strings.Contains(coordinate, "F") {
		if isSpecial(matrix[x+1][y-1]) {

			return true

		}
	}

	if strings.Contains(coordinate, "G") {
		if isSpecial(matrix[x+1][y]) {

			return true

		}
	}

	if strings.Contains(coordinate, "H") {
		if isSpecial(matrix[x+1][y+1]) {

			return true

		}
	}

	if coordinate == "*" {
		return isSpecial(matrix[x-1][y-1]) ||
			isSpecial(matrix[x-1][y]) ||
			isSpecial(matrix[x-1][y+1]) ||
			isSpecial(matrix[x][y-1]) ||
			isSpecial(matrix[x][y+1]) ||
			isSpecial(matrix[x+1][y-1]) ||
			isSpecial(matrix[x+1][y]) ||
			isSpecial(matrix[x+1][y+1])
	}

	return false
}

// isAdjacentToNumber checks if a coordinate is adjacent to a number
func isAdjacentToNumber(coordinate string, matrix [][]string, x int, y int) bool {
	if strings.Contains(coordinate, "A") {
		if isANumber(matrix[x-1][y-1]) {

			return true

		}
	}

	if strings.Contains(coordinate, "B") {
		if isANumber(matrix[x-1][y]) {

			return true

		}
	}

	if strings.Contains(coordinate, "C") {
		if isANumber(matrix[x-1][y+1]) {

			return true

		}
	}

	if strings.Contains(coordinate, "D") {
		if isANumber(matrix[x][y-1]) {

			return true

		}
	}

	if strings.Contains(coordinate, "E") {
		if isANumber(matrix[x][y+1]) {

			return true

		}
	}

	if strings.Contains(coordinate, "F") {
		if isANumber(matrix[x+1][y-1]) {

			return true

		}
	}

	if strings.Contains(coordinate, "G") {
		if isANumber(matrix[x+1][y]) {

			return true

		}
	}

	if strings.Contains(coordinate, "H") {
		if isANumber(matrix[x+1][y+1]) {

			return true

		}
	}

	return false
}
