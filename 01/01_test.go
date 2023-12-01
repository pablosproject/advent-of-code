package main

import "testing"

func TestProcessNumber(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, test := range tests {
		res := processNumber(test.input)
		if res != test.output {
			t.Errorf("processNumber(%s) = %d, want %d", test.input, res, test.output)
		}
	}
}
