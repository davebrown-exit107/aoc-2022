package main

import (
	"bufio"
	"fmt"
	"os"
)

// Solver for Day 06 Pt 01
func DaySixPartOne(input string) (numChars int) {
	return
}

/*
// Solver for Day 06 Pt 02
func DaySixPartTwo(input []string) string {
	return
}
*/

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
	numChars := DaySixPartOne(fileLines[0])
	fmt.Printf("Characters needed for processing: %v\n", numChars)

	/*
		newStackTops := DaySixPartTwo(fileLines)
		fmt.Printf("New tops of stacks: %v\n", newStackTops)
	*/
}
