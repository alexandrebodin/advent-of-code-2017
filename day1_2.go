package main

import "strconv"

// Day1_2 => exercise 1.2
func Day1_2(input []byte) int {
	sum := 0
	len := len(input)
	offset := len / 2
	for i := 0; i < len; i++ {
		j := (i + offset) % len
		if input[i] == input[j] {
			val, _ := strconv.Atoi(string(input[i]))
			sum += val
		}
	}

	return sum
}
