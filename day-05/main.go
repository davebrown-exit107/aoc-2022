package main

import (
	"bufio"
	"os"
)

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
	//	fullyOverlaps := CountFullyOverlappingAssignments(fileLines)
	//	fmt.Printf("Count of full overlaps: %v\n", fullyOverlaps)
	//
	//	overlaps := CountOverlappingAssignments(fileLines)
	//	fmt.Printf("Count of overlaps: %v\n", overlaps)
}
