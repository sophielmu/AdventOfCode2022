package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printPartOneAnswer() {
	allCompartments := splitStringsIntoCompartments()
	priorityTotal := calculatePriorityTotal(allCompartments)
	fmt.Printf("%s: %s", "\nDay 3 part 1 answer", strconv.Itoa(priorityTotal))
}

func getPriorityForChar(char byte) int {
	var priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < len(priorities); i++ {
		if priorities[i] == char {
			return i + 1
		}
	}
	panicMessage := fmt.Sprintf("Cannot find matching priority character for %s", string(char))
	panic(panicMessage)
}

func calculatePriorityTotal(allCompartments []string) int {

	allMatches := make([]byte, 0)

	for i := 0; i < len(allCompartments); i++ {
		if i%2 == 1 {
			continue
		} else {
			matchingCharacter := getMatchingCharacter(allCompartments[i], allCompartments[i+1])
			allMatches = append(allMatches, matchingCharacter)
		}
	}

	var priorityTotal = 0

	for i := 0; i < len(allMatches); i++ {
		priorityTotal += getPriorityForChar(allMatches[i])
	}
	return priorityTotal
}

func getMatchingCharacter(compartment1 string, compartment2 string) byte {
	for j := 0; j < len(compartment1); j++ {
		for k := 0; k < len(compartment1); k++ {
			if compartment1[j] == compartment2[k] {
				// WE FOUND IT
				return compartment1[j]
			}
		}
	}
	panicMessage := fmt.Sprintf("No character matched for compartments %s and %s", compartment1, compartment2)
	panic(panicMessage)
}

func splitStringsIntoCompartments() []string {
	file, _ := os.Open("./day3/day3input.txt")
	scanner := bufio.NewScanner(file)

	allCompartments := make([]string, 0)

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		} else {
			currentCompartment := make([]rune, 0)
			for i := 0; i < len(l); i++ {
				currentCompartment = append(currentCompartment, rune(l[i]))
				if (i+1) == len(l)/2 || (i+1 == len(l)) {
					allCompartments = append(allCompartments, string(currentCompartment))
					currentCompartment = make([]rune, 0)
				}
			}
		}
	}
	return allCompartments
}
