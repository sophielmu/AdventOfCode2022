package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printPartTwoAnswer() {
	allRucksacks := readDay3Input()
	commonItems := getCommonItemForAllGroups(allRucksacks)

	totalPriority := 0

	for i := 0; i < len(commonItems); i++ {
		totalPriority += getPriorityForChar(commonItems[i])
	}

	fmt.Printf("%s: %s", "Day 3 part 2 answer", strconv.Itoa(totalPriority))
}

func getCommonItemForAllGroups(allRucksacks []string) (commonItems []byte) {

	for i := 0; i < len(allRucksacks); i++ {
		if i%3 != 0 {
			continue
		}

		for j := 0; j < len(allRucksacks[i]); j++ {
			var currentChar = string(allRucksacks[i][j])
			if strings.Contains(allRucksacks[i+1], currentChar) {
				if strings.Contains(allRucksacks[i+2], currentChar) {
					commonItems = append(commonItems, allRucksacks[i][j])
					break
				}
			}
		}
	}
	return commonItems
}

func readDay3Input() (allRucksacks []string) {
	file, _ := os.Open("./day3/day3input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		} else {
			allRucksacks = append(allRucksacks, l)
		}
	}
	return allRucksacks
}
