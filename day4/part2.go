package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func printPartTwoAnswer() {
	var allLines = readDay4Input()
	answer := getNumberOfOverlaps(allLines)
	fmt.Printf("%s: %s", "Day 4 part 2 answer", strconv.Itoa(answer))
}

func getNumberOfOverlaps(pairs []string) (overlapCount int) {
	for _, pair := range pairs {
		var sectionIds = strings.FieldsFunc(pair, Split)
		sectionIdsInt := make([]int, 0)

		for _, section := range sectionIds {
			var toInt, _ = strconv.Atoi(section)
			sectionIdsInt = append(sectionIdsInt, toInt)
		}

		if overlapExists(sectionIdsInt[0], sectionIdsInt[1], sectionIdsInt[2], sectionIdsInt[3]) {
			overlapCount++
		}
	}
	return
}

func overlapExists(x1 int, x2 int, y1 int, y2 int) bool {
	return (x1 >= y1 && x1 <= y2) ||
		(x2 >= y1 && x2 <= y2) ||
		(y1 >= x1 && y1 <= x2) ||
		(y2 >= x1 && y2 <= x2)
}
