package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printPartOneAnswer() {
	answer := calculateMostCaloriesCarried()
	fmt.Println(answer)
}

func calculateMostCaloriesCarried() int {
	file, _ := os.Open("./day1/day1input.txt")
	scanner := bufio.NewScanner(file)

	currentElf := 0
	maxElf := 0

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			if currentElf > maxElf {
				maxElf = currentElf
			}
			currentElf = 0
		} else {
			calories, _ := strconv.Atoi(l)
			currentElf += calories
		}
	}
	return maxElf
}
