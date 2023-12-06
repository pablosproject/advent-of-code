package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type number struct {
	row        int
	start, end int
	value      int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fileMap [][]rune
	var numbers []number

	for scanner.Scan() {
		line := scanner.Text()
		// Number matrix creation (for easy reference)
		row := make([]rune, len(line))
		for i, char := range line {
			row[i] = char
		}
		fileMap = append(fileMap, row)

		// Extract all numbers
		numbers = append(numbers, findNumbers(line, len(fileMap)-1)...)
	}

	// Sum all numbers with a correct number closenumbers.
	sum := 0
	for _, number := range numbers {
		if findCloseDigit(number, fileMap) {
			sum += number.value
		}
	}

	fmt.Println(sum)
}

func findNumbers(line string, row int) []number {
	var numbers []number
	numRegex := regexp.MustCompile(`(\d+)`)
	for _, match := range numRegex.FindAllStringIndex(line, -1) {
		value, _ := strconv.Atoi(line[match[0]:match[1]])
		numbers = append(numbers, number{row, match[0], match[1] - 1, value})
	}

	return numbers
}

func findCloseDigit(number number, fileMap [][]rune) bool {
	columsStart := number.start - 1
	columsEnd := number.end + 1
	rowsStart := number.row - 1
	rowsEnd := number.row + 1

	for row := rowsStart; row <= rowsEnd; row++ {
		if row < 0 || row >= len(fileMap) {
			continue
		}

		for colum := columsStart; colum <= columsEnd; colum++ {
			if colum < 0 || colum >= len(fileMap[row]) {
				continue
			}

			if fileMap[row][colum] == '.' || isDigit(fileMap[row][colum]) {
				continue
			}

			return true
		}
	}

	return false
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
