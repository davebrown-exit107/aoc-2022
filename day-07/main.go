package main

import (
	"bufio"
	"fmt"
	"os"
)

// Solver for Day 07 Pt 01
func DaySevenPartOne(input []string) (numChars int) {
	return
}

/*
// Solver for Day 07 Pt 02
func DaySevenPartTwo(input []string) (numChars int) {
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
	directorySize := DaySevenPartOne(fileLines)
	fmt.Printf("Sum of directory sizes: %v\n", directorySize)

	/*
		numChars = DaySevenPartTwo(fileLines)
		fmt.Printf("Characters til first message: %v\n", numChars)
	*/
}
