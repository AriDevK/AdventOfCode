package main

import (
	"fmt"
	"strconv"
	"strings"
)

// partOne: sum all the numbers that are adjacent to a special character
func partOne() {
	var output int
	rows := strings.Split(input, "\n")
	m := make([][]string, len(rows))

	// fill the matrix with the input
	for i, row := range rows {
		m[i] = strings.Split(row, "")
	}

	// iterate over the rows
	for i := 0; i < len(rows); i++ {
		var aux string
		var isValid bool

		// iterate over the columns
		for j := 0; j < len(rows[i]); j++ {
			char := m[i][j]
			if isANumber(char) {
				aux += char
				if !isValid {
					if i == 0 {
						if j == 0 {
							isValid = isAdjacent("EGH", m, i, j)
						}

						if j == len(rows[i])-1 {
							isValid = isAdjacent("DFG", m, i, j)
						}

						if j > 0 && j < len(rows[i])-1 {
							isValid = isAdjacent("DEFGH", m, i, j)
						}
					} else if i == len(rows)-1 {
						if j == 0 {
							isValid = isAdjacent("BCE", m, i, j)
						}

						if j == len(rows[i])-1 {

							isValid = isAdjacent("ABD", m, i, j)
						}

						if j > 0 && j < len(rows[i])-1 {
							isValid = isAdjacent("ABCDE", m, i, j)
						}
					} else {
						if j == 0 {
							isValid = isAdjacent("BCEGH", m, i, j)
						}

						if j == len(rows[i])-1 {
							isValid = isAdjacent("ABDFG", m, i, j)
						}

						if j > 0 && j < len(rows[i])-1 {
							isValid = isAdjacent("*", m, i, j)
						}
					}
				}
			}

			if (char == "." || isSpecial(char) || j == len(rows[i])-1) && isValid {
				auxVal, _ := strconv.Atoi(aux)
				output += auxVal
				isValid = false
				aux = ""
			} else if char == "." && !isValid {
				aux = ""
			}
		}
	}

	fmt.Println(output)
}
