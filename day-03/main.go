package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func intValueOf(letter string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(alphabet, letter) + 1
}

func dedupeArray(arrIn []string) []string {
	seen := []string{}
	for _, val := range arrIn {
		if !arrayContains(seen, val) {
			seen = append(seen, val)
		}
	}
	return seen
}

func arrayContains(arrIn []string, contains string) bool {
	for _, val := range arrIn {
		if val == contains {
			return true
		}
	}
	return false
}

func SumBadges(rucksacks []string) int {
	curGroup := make([]string, 0)
	badges := make([]string, 0)
	for _, sack := range rucksacks {
		// we're looking for groupings of three sacks to compare against
		curGroup = append(curGroup, sack)
		if len(curGroup)%3 == 0 {
			// we have three, start to scan
			// prep the groups for comparison
			dedupeSubSackOne := dedupeArray(strings.Split(curGroup[0], ""))
			dedupeSubSackTwo := dedupeArray(strings.Split(curGroup[1], ""))
			dedupeSubSackThree := dedupeArray(strings.Split(curGroup[2], ""))
			// scan through sack one, checking the other two sacks at the same time
			for _, item := range dedupeSubSackOne {
				if arrayContains(dedupeSubSackTwo, item) == true && arrayContains(dedupeSubSackThree, item) == true {
					// if you're a member of all three, you're a badge
					badges = append(badges, item)
				}
			}
			curGroup = []string{}
		}
	}
	// sum up the badges
	totals := 0
	for _, item := range badges {
		totals += intValueOf(item)
	}
	return totals
}

func SumDuplicates(rucksacks []string) int {
	totals := 0
	dupes := make([]string, 0)
	for _, rucksack := range rucksacks {
		// prep the array for comparison
		curLen := len(rucksack)
		halfsies := (curLen) / 2
		// split the rucksack in half and assign to compartments
		compOne := rucksack[0:halfsies]
		compTwo := rucksack[halfsies:curLen]
		// break each compartment into its component items and dedupe it
		compOneSlice := dedupeArray(strings.Split(compOne, ""))
		compTwoSlice := dedupeArray(strings.Split(compTwo, ""))
		// check for dupes between compartments
		for _, item := range compOneSlice {
			if arrayContains(compTwoSlice, item) == true {
				dupes = append(dupes, item)
			}
		}
	}
	// sum up the duplicates
	for _, item := range dupes {
		totals += intValueOf(item)
	}
	return totals
}

func main() {
	// File opening boilerplate
	readFile, err := os.Open("input.txt")
	if err != nil {
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	// AoC work starts here
	extrasTotal := SumDuplicates(fileLines)
	fmt.Printf("Sum of extras: %v\n", extrasTotal)

	badgesTotal := SumBadges(fileLines)
	fmt.Printf("Sum of badges per 3 elves: %v\n", badgesTotal)

}
