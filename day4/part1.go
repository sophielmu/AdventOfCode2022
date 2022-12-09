package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printPartOneAnswer() {
	allLines := readDay4Input()
	overlapCount := getNumberOfPairsWithFullyContainedSections(allLines)
	fmt.Printf("%s: %s", "\nDay 4 part 1 answer", strconv.Itoa(overlapCount))
}

func getNumberOfPairsWithFullyContainedSections(pairs []string) (overlapCount int) {
	for _, pair := range pairs {
		var sectionIds = strings.FieldsFunc(pair, Split)
		sectionIdsInt := make([]int, 0)

		for _, section := range sectionIds {
			var toInt, _ = strconv.Atoi(section)
			sectionIdsInt = append(sectionIdsInt, toInt)
		}

		if (sectionIdsInt[0] >= sectionIdsInt[2]) && (sectionIdsInt[1] <= sectionIdsInt[3]) {
			overlapCount++
			continue
		}
		if (sectionIdsInt[2] >= sectionIdsInt[0]) && (sectionIdsInt[3] <= sectionIdsInt[1]) {
			overlapCount++
		}
	}
	return
}

func Split(r rune) bool {
	return r == '-' || r == ','
}

func readDay4Input() (allLines []string) {
	file, _ := os.Open("./day4/day4input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		allLines = append(allLines, l)
	}
	return
}
