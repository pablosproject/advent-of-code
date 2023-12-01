package main

import "testing"

func TestProcessNumber(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
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
