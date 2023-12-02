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
	index := gameRex.FindStringSubmatchIndex(line)

	extractions := strings.Split(line[index[1]:], ";")
	max := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, extraction := range extractions {
		maxCubes(extraction, max)
	}

	return power(max)
}

func maxCubes(extraction string, max map[string]int) {
	colors := strings.Split(extraction, ",")

	for _, color := range colors {
		for _, cube := range cubes {
			if strings.Contains(color, cube.color) {
				numbers := strings.Split(color, cube.color)
				number, err := strconv.Atoi(strings.Trim(numbers[0], " "))
				if err != nil {
					panic(err)
				}
				if number > max[cube.color] {
					max[cube.color] = number
				}
			}
		}
	}

	fmt.Printf("%v\n", max)
}

func power(min map[string]int) int {
	return min["red"] * min["green"] * min["blue"]
}
