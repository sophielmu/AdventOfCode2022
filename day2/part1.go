package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printPartOneAnswer() {
	answer := calculateScore()
	fmt.Printf("%s: %s", "Day 2 part 1 answer", strconv.Itoa(answer))
}

func calculateScore() int {
	file, _ := os.Open("./day2/day2input.txt")
	scanner := bufio.NewScanner(file)

	total := 0

	results := make(map[string]int)
	results["A X"] = 3
	results["A Y"] = 6
	results["A Z"] = 0
	results["B X"] = 0
	results["B Y"] = 3
	results["B Z"] = 6
	results["C X"] = 6
	results["C Y"] = 0
	results["C Z"] = 3

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		} else {
			total += getShapeScore(l)
			total += results[l]
		}
	}
	return total
}

func getShapeScore(line string) int {
	if strings.Contains(line, "X") {
		return 1
	} else if strings.Contains(line, "Y") {
		return 2
	} else if strings.Contains(line, "Z") {
		return 3
	}
	return 0
}
