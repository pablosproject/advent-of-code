package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type numbersWord struct {
	word  string
	value string
}

var numbers = []numbersWord{
	{"one", "1"},
	{"two", "2"},
	{"three", "3"},
	{"four", "4"},
	{"five", "5"},
	{"six", "6"},
	{"seven", "7"},
	{"eight", "8"},
	{"nine", "9"},
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	number := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineVal := processNumber(line)
		number += lineVal
	}

	fmt.Printf("%d\n", number)
}

type numberPos struct {
	position int
	number   string
}

func processNumber(line string) int {
	numbers := searchNumbersOccurrence(line)
	numbersWord := searchNumberWordOccurrence(line)

	numbers = append(numbers, numbersWord...)
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].position < numbers[j].position
	})

	last := numbers[len(numbers)-1]
	first := numbers[0]

	num := first.number + last.number

	res, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	return res
}

func searchNumberWordOccurrence(line string) []numberPos {
	res := []numberPos{}

	for _, numWord := range numbers {
		re := regexp.MustCompile(numWord.word)

		positions := re.FindAllStringIndex(line, -1)

		for _, pos := range positions {
			position := pos[0]
			value := numWord.value
			res = append(res, numberPos{position, value})
		}
	}
	return res
}

func searchNumbersOccurrence(line string) []numberPos {
	re := regexp.MustCompile("[0-9]")
	numbers := re.FindAllStringIndex(line, -1)

	res := []numberPos{}
	for _, numberLoc := range numbers {
		position := numberLoc[0]
		char := line[position]
		res = append(res, numberPos{position, string(char)})
	}
	return res
}
