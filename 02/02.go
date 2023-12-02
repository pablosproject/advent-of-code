package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cubeNumber struct {
	color string
	max   int
}

var cubes = []cubeNumber{
	{"red", 12},
	{"green", 13},
	{"blue", 14},
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
		number += processLine(line)
	}

	fmt.Printf("%d\n", number)
}

func processLine(line string) int {
	gameRex := regexp.MustCompile(`Game (\d+):`)
	matches := gameRex.FindStringSubmatch(line)
	index := gameRex.FindStringSubmatchIndex(line)

	game, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}

	extractions := strings.Split(line[index[1]:], ";")
	for _, extraction := range extractions {
		if exceedLimit(extraction) {
			return 0
		}
	}

	return game
}

func exceedLimit(extraction string) bool {
	colors := strings.Split(extraction, ",")

	for _, color := range colors {
		for _, cube := range cubes {
			if strings.Contains(color, cube.color) {
				numbers := strings.Split(color, cube.color)
				number, err := strconv.Atoi(strings.Trim(numbers[0], " "))
				if err != nil {
					panic(err)
				}
				if number > cube.max {
					return true
				}
			}
		}
	}

	return false
}
