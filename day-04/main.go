package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isOverlapping(groupA, groupB string) bool {
	// parse out the ranges, there's got to be a better way to do this
	groupAPairs := strings.Split(groupA, "-")
	groupBPairs := strings.Split(groupB, "-")
	groupAStart, err := strconv.Atoi(groupAPairs[0])
	if err != nil {
		os.Exit(1)
	}
	groupAEnd, err := strconv.Atoi(groupAPairs[1])
	if err != nil {
		os.Exit(1)
	}
	groupBStart, err := strconv.Atoi(groupBPairs[0])
	if err != nil {
		os.Exit(1)
	}
	groupBEnd, err := strconv.Atoi(groupBPairs[1])
	if err != nil {
		os.Exit(1)
	}
	if (groupAStart <= groupBStart &&
		groupAEnd >= groupBEnd) ||
		(groupBStart <= groupAStart &&
			groupBEnd >= groupAEnd) {
		return true
	} else {
		return false
	}
}

func CountOverlappingAssignments(assignments []string) int {
	curGroup := make([]string, 0)
	overlaps := 0
	for _, assignment := range assignments {
		//evaluate the groups to look for overlapping assignments
		curGroup = strings.Split(assignment, ",")
		if isOverlapping(curGroup[0], curGroup[1]) == true {
			overlaps++
		}
		curGroup = make([]string, 0)
	}
	return overlaps
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
	overlaps := CountOverlappingAssignments(fileLines)
	fmt.Printf("Count of overlaps: %v\n", overlaps)

}
