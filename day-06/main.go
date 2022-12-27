package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Solver for Day 06 Pt 01
func DaySixPartOne(input string) (numChars int) {
	numChars = findMarker(input, 14)
	for i := 0; i < len(input)-4; i++ {
		if string(input[i]) != string(input[i+1]) &&
			string(input[i]) != string(input[i+2]) &&
			string(input[i]) != string(input[i+3]) &&
			string(input[i+1]) != string(input[i+2]) &&
			string(input[i+1]) != string(input[i+3]) &&
			string(input[i+2]) != string(input[i+3]) {
			numChars = i + 4
			break
		}
	}
	return
}

// Looks for duplicate charaters in a slice of a string (the size defined by window)
// and returns the position of the start of the window where all charaters are unique
func findMarker(input string, window int) (pos int) {
	for i := 0; i < len(input); i++ {
		curWindow := input[i : i+window]
		curUnique := 0
		for _, curChar := range curWindow {
			if strings.Count(curWindow, string(curChar)) > 1 {
				// if there's a dupe, break out of this window and continue on to the next
				break
			} else {
				// else, this is a unique character. note it and move on
				curUnique++
			}
		}
		if curUnique == window {
			// if all of the characters in the window are unique, we've found the marker
			// return the position plus the window (as the marker is not part of the signal)
			pos = i + window
			break
		}
	}
	return
}

// Solver for Day 06 Pt 02
func DaySixPartTwo(input string) (numChars int) {
	// Same as before, but 14 chars instead of 4...
	// Probably need to break this out into a function that can take a string and variable number of chars
	numChars = findMarker(input, 14)
	return
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
	numChars := DaySixPartOne(fileLines[0])
	fmt.Printf("Characters needed for processing: %v\n", numChars)

	numChars = DaySixPartTwo(fileLines[0])
	fmt.Printf("Characters til first message: %v\n", numChars)
}
