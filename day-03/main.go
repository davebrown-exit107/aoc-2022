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
		fmt.Println(err)
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

	//newTotalScore := NewScoreRounds(fileLines)
	//fmt.Printf("The new total score is %v\n", newTotalScore)

}
