package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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

func processNumber(number string) int {
	re := regexp.MustCompile("[0-9]")
	numbers := re.FindAllString(number, -1)

	last := numbers[len(numbers)-1]
	first := numbers[0]

	num := first + last

	res, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	return res
}
