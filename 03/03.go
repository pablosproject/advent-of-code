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

type gear struct {
	row, col int
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
	var gears []gear

	for scanner.Scan() {
		line := scanner.Text()
		// Number matrix creation (for easy reference)
		row := make([]rune, len(line))
		for i, char := range line {
			row[i] = char
		}
		fileMap = append(fileMap, row)

		// Extract all numbers and gears
		numbers = append(numbers, findNumbers(line, len(fileMap)-1)...)
		gears = append(gears, findGears(line, len(fileMap)-1)...)
	}

	sumGearRatio := 0
	for _, gear := range gears {
		sumGearRatio += findGearRatio(gear, numbers)
	}

	fmt.Printf("%d\n", sumGearRatio)
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

func findGears(line string, row int) []gear {
	var gears []gear
	gearRegex := regexp.MustCompile(`\*`)
	for _, match := range gearRegex.FindAllStringSubmatchIndex(line, -1) {
		gears = append(gears, gear{row, match[0]})
	}

	return gears
}

func findGearRatio(gear gear, numbers []number) int {
	var closeNumbers []number
	for _, number := range numbers {
		sameRow := number.row == gear.row || number.row == gear.row-1 || number.row == gear.row+1
		sameCol := gear.col >= number.start-1 && gear.col <= number.end+1
		if sameRow && sameCol {
			closeNumbers = append(closeNumbers, number)
		}
	}

	if len(closeNumbers) == 2 {
		return closeNumbers[0].value * closeNumbers[1].value
	}

	return 0
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
