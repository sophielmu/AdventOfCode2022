package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printPartTwoAnswer() {
	answer := calculatePart2Score()
	fmt.Printf("%s: %s", "Day 2 part 2 answer", strconv.Itoa(answer))
}

func calculatePart2Score() int {
	file, _ := os.Open("./day2/day2input.txt")
	scanner := bufio.NewScanner(file)

	total := 0

	results := make(map[string]int)
	results["A X"] = 3
	results["A Y"] = 1
	results["A Z"] = 2
	results["B X"] = 1
	results["B Y"] = 2
	results["B Z"] = 3
	results["C X"] = 2
	results["C Y"] = 3
	results["C Z"] = 1

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		} else {
			total += getOutcomeScore(l)
			total += results[l]
		}
	}
	return total
}

func getOutcomeScore(line string) int {
	if strings.Contains(line, "X") {
		return 0
	} else if strings.Contains(line, "Y") {
		return 3
	} else if strings.Contains(line, "Z") {
		return 6
	}
	return 0
}
