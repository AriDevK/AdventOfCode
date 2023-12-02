package main

import "strings"

// removeLetters removes all letters from a string.
func removeLetters(data string) []string {
	var output string
	dataList := strings.Split(data, "")

	for _, c := range dataList {
		if c >= "0" && c <= "9" {
			output += c
		}
	}

	return strings.Split(output, "")
}

// swapNameWithValue swaps the name of the number with the value of the number.
func swapNameWithValue(data string) string {
	var number string
	output := data
	splitData := strings.Split(data, "")
	numberList := map[string]string{
		"nine":  "9",
		"eight": "8",
		"seven": "7",
		"six":   "6",
		"five":  "5",
		"four":  "4",
		"three": "3",
		"two":   "2",
		"one":   "1",
	}

	for _, v := range splitData {
		number += v                               // build the string number
		numberKey := contains(numberList, number) // check if the number is in the map

		if numberKey != "" {
			// remove the last letter from the number to get the tricky number
			// the edge case is when the input is like this:
			// e.g. oneeight -> 1eight -> 18
			// so we need to remove the last letter to get the tricky number
			// really tricky isn't it :D
			trickyNumber := numberKey[0 : len(numberKey)-1]

			// replace the number with the value
			output = strings.Replace(output, trickyNumber, numberList[numberKey], 1)

			// reset the number to the last letter of the number
			number = numberKey[len(numberKey)-1 : len(numberKey)]
		}
	}

	return output
}

// contains checks if a map contains a key and returns the key if it does exist.
func contains(array map[string]string, value string) string {
	for k, _ := range array {
		if strings.HasSuffix(value, k) {
			return k
		}
	}

	return ""
}
