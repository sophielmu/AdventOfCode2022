package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func printPartTwoAnswer() {
	answer := calculateTotalOfTop3Elves()
	fmt.Println(answer)
}

func calculateTotalOfTop3Elves() int {
	file, _ := os.Open("./day1/day1input.txt")
	scanner := bufio.NewScanner(file)

	currentElf := 0
	calorieTotals := make([]int, 0)

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			calorieTotals = append(calorieTotals, currentElf)
			currentElf = 0
		} else {
			calories, _ := strconv.Atoi(l)
			currentElf += calories
		}
	}
	sort.Ints(calorieTotals)
	return calorieTotals[len(calorieTotals)-1] + calorieTotals[len(calorieTotals)-2] + calorieTotals[len(calorieTotals)-3]
}
