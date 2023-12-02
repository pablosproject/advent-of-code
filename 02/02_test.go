package main

import "testing"

func TestProcessLine(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 1},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 2},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 0},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 0},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 5},
	}

	for _, test := range tests {
		res := processLine(test.input)
		if res != test.output {
			t.Errorf("processLine(%s) = %d, want %d", test.input, res, test.output)
		}
	}
}
