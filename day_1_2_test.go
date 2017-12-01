package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1_2(t *testing.T) {

	var tableTests = []struct {
		in       []byte
		expected int
	}{
		{[]byte("1212"), 6},
		{[]byte("1221"), 0},
		{[]byte("123425"), 4},
		{[]byte("123123"), 12},
		{[]byte("12131415"), 4},
	}

	for _, tt := range tableTests {
		actual := Day1_2(tt.in)
		assert.Equal(t, tt.expected, actual)
	}

}
