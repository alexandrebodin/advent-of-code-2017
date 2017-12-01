package main

import "strconv"

// Day1_1 => exercise 1.1
func Day1_1(input []byte) int {
	sum := 0
	len := len(input)
	for i := 0; i < len; i++ {
		j := (i + 1) % (len)
		if input[i] == input[j] {
			val, _ := strconv.Atoi(string(input[i]))
			sum += val
		}
	}

	return sum
}
