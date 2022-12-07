package main

import (
	"bufio"
	"fmt"
	"os"
)

// i'm going to get adventerous and try to implement the stacks as linked lists
type crate struct {
	contents string
	next     *crate
}

type stack struct {
	head   *crate
	length int
}

func (s *stack) pop() *crate {
	if s.length >= 1 {
		retVal := s.head
		s.head = s.head.next
		return retVal
	} else {
		fmt.Println("no members to pop off of stack")
		return nil
	}
}

func (s *stack) add(c *crate) {
	c.next = s.head
	s.head = c
	s.length++
}

func parseInput(input []string) (stacks []stack, moves []string) {
	// parse the input file
	return []stack{}, []string{}
}

func parseStacks(input []string) []stack {
	// this will parse out the stacks from the input into the stacks struct
	return []stack{}
}

func craneOperator(stacks []stack, moves []string) []stack {
	// this will actually manipulate the stacks according to the listed moves
	return []stack{}
}

func checkTops(stacks []stack) string {
	// just checking the tops of each stack and return it as a string
	return ""
}

func DayFivePartOne(input []string) string {
	// the main function the first part of the challenge goes through
	// first we parse the input file and separate the stacks from the moves
	// then we step through the moves and apply them to the stacks
	// then we check the tops and return them
	return ""
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
	stackTops := DayFivePartOne(fileLines)
	fmt.Printf("Final tops of stacks: %v\n", stackTops)
	//
	//	overlaps := CountOverlappingAssignments(fileLines)
	//	fmt.Printf("Count of overlaps: %v\n", overlaps)
}
