package main

import (
	"fmt"
	"strconv"
	"strings"
)

// partTwo: multiply all the numbers that are adjacent to an engine (*)
func partTwo() {
	var cnt int
	var output int
	rows := strings.Split(input, "\n")
	m := make([][]string, len(rows))

	for i, row := range rows {
		m[i] = strings.Split(row, "")
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(m[i]); j++ {
			var numbers []int
			isEngine := m[i][j] == "*"
			if isEngine {
				ce := checkEngine(m, i, j)
				if len(ce) == 2 {
					for _, c := range ce {
						if i == 0 {
							var aux string

							if j == 0 {
								if c == "E" {
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}
								if c == "G" {
									jx := j
									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}
								if c == "H" {
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}
							}

							if j == len(rows[i])-1 {
								if c == "D" {
									// is on top on before column to the right
									jx := j - 1
									var valBf string

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux)
									val, _ := strconv.Atoi(valBf)
									numbers = append(numbers, val)
								}

								if c == "F" {
									// is on top on before column to the right
									jx := j - 1
									var valNx string
									var valBf string

									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									// remove first character because is duplicated on reverse
									valNx = aux

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux[1:])
									val, _ := strconv.Atoi(valBf + valNx)
									numbers = append(numbers, val)
								}

								if c == "G" {
									val, _ := strconv.Atoi(m[i+1][j])
									numbers = append(numbers, val)
								}
							}

							if j > 0 && j < len(rows[i])-1 {

								if c == "D" {
									// is on top on before column to the right
									jx := j - 1
									var valBf string

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux)
									val, _ := strconv.Atoi(valBf)
									numbers = append(numbers, val)
								}

								if c == "E" {
									// IS TO THE RIGHT
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "F" {
									// is on top on before column to the right
									jx := j - 1
									var valNx string
									var valBf string

									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									// remove first character because is duplicated on reverse
									valNx = aux

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux[1:])
									val, _ := strconv.Atoi(valBf + valNx)
									numbers = append(numbers, val)
								}

								if c == "G" {
									jx := j
									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "H" {
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

							}

						} else if i == len(rows)-1 {

							var aux string

							if j == 0 {

								if c == "B" {
									// is on top to the right on same column
									jx := j
									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}
									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "C" {
									// is on top to the right on next column
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}
									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "E" {
									// is on same row to the right
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

							}

							if j == len(rows[i])-1 {

								if c == "A" {
									// is on top on before column to the right
									jx := j - 1
									var valNx string
									var valBf string

									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									// remove first character because is duplicated on reverse
									valNx = aux

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux[1:])
									val, _ := strconv.Atoi(valBf + valNx)
									numbers = append(numbers, val)
								}

								if c == "B" {
									// is on the same column to the top, is the last column so we only need to check the value at the top
									val, _ := strconv.Atoi(m[i-1][j])
									numbers = append(numbers, val)
								}

								if c == "D" {
									// is on top on before column to the right
									jx := j - 1
									var valBf string

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux)
									val, _ := strconv.Atoi(valBf)
									numbers = append(numbers, val)
								}

							}

							if j > 0 && j < len(rows[i])-1 {
								if c == "A" {
									// is on top on before column to the right
									jx := j - 1
									var valNx string
									var valBf string

									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									// remove first character because is duplicated on reverse
									valNx = aux

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux[1:])
									val, _ := strconv.Atoi(valBf + valNx)
									numbers = append(numbers, val)
								}

								if c == "B" {
									// is on top on before column to the right
									jx := j
									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}
									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "C" {
									// is on top on before column to the right
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}
									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "D" {
									// is on top on before column to the right
									jx := j - 1
									var valBf string

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux)
									val, _ := strconv.Atoi(valBf)
									numbers = append(numbers, val)
								}

								if c == "E" {
									// is on same row to the right
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}
										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}
							}

						} else {

							var aux string

							if j == 0 {

								// THIS WILL CHECK B C E G H
								if c == "B" {
									// is on top on before column to the right
									jx := j
									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}
									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "C" {
									// is on top on before column to the right
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}
									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "E" {
									// is on same row to the right
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "G" {
									// is on bottom on same column
									jx := j
									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "H" {
									// is on bottom on next column
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}
							}

							if j == len(rows[i])-1 {

								// THIS WILL CHECK A B D F G
								if c == "A" {
									// is on top on before column to the right
									jx := j - 1
									var valNx string
									var valBf string

									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									// remove first character because is duplicated on reverse
									valNx = aux

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux[1:])
									val, _ := strconv.Atoi(valBf + valNx)
									numbers = append(numbers, val)
								}

								if c == "B" {
									// THIS IS ON TOP ON THE SAME COLUMN
									val, _ := strconv.Atoi(m[i-1][j])
									numbers = append(numbers, val)
								}

								if c == "D" {
									// is on top on before column to the right
									jx := j - 1
									var valBf string

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux)
									val, _ := strconv.Atoi(valBf)
									numbers = append(numbers, val)

								}

								if c == "F" {
									// is on top on before column to the right
									jx := j - 1
									var valNx string
									var valBf string

									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									// remove first character because is duplicated on reverse
									valNx = aux

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux[1:])
									val, _ := strconv.Atoi(valBf + valNx)
									numbers = append(numbers, val)
								}

								if c == "G" {
									// IS ON THE SAME COLUMN ON NEXT ROW
									val, _ := strconv.Atoi(m[i+1][j])
									numbers = append(numbers, val)
								}
							}

							if j > 0 && j < len(rows[i])-1 {

								aux = ""

								// THIS WILL CHECK A B C D E F G H
								if c == "A" {
									// is on top on before column to the right
									jx := j - 1
									var valNx string
									var valBf string

									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									// remove first character because is duplicated on reverse
									valNx = aux

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux[1:])
									val, _ := strconv.Atoi(valBf + valNx)
									numbers = append(numbers, val)
								}

								if c == "B" {
									// is on top on before column to the right
									jx := j
									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}
									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "C" {
									// is on top on before column to the right
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i-1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}
									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "D" {
									// is on top on before column to the right
									jx := j - 1
									var valBf string

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux)
									val, _ := strconv.Atoi(valBf)
									numbers = append(numbers, val)
								}

								if c == "E" {
									// is on same row to the right
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "F" {
									// is on bot on before column to the right
									jx := j - 1

									var valNx string
									var valBf string

									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									valNx = aux

									// now on reverse
									aux = ""
									jx = j - 1
									for jx >= 0 {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx--
									}

									valBf = Reverse(aux[1:])
									val, _ := strconv.Atoi(valBf + valNx)
									numbers = append(numbers, val)
								}

								if c == "G" {
									jx := j
									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}

								if c == "H" {
									jx := j + 1
									for jx < len(rows[i]) {
										val := m[i+1][jx]
										if isANumber(val) {
											aux += val
										} else if isSpecial(val) || val == "." {
											break
										}

										jx++
									}

									val, _ := strconv.Atoi(aux)
									numbers = append(numbers, val)
								}
							}
						}
					}

					output += numbers[0] * numbers[1]
					cnt++
				}
			}
		}

	}

	fmt.Println(output)
}
