package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1_1(t *testing.T) {

	var tableTests = []struct {
		in       []byte
		expected int
	}{
		{[]byte("1122"), 3},
		{[]byte("1111"), 4},
		{[]byte("1234"), 0},
		{[]byte("91212129"), 9},
	}

	for _, tt := range tableTests {
		actual := Day1_1(tt.in)
		assert.Equal(t, tt.expected, actual)
	}

}
