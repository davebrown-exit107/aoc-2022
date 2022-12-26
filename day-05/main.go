package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Simple container for moves to keep magic values out of the functions
type Move struct {
	qty  int
	from int
	to   int
}

// Parse the input file
func parseInput(input []string) (stacks [][]string, moves []Move) {
	unparsedStacks := make([]string, 0)
	dividingLine := 0
	unparsedMoves := make([]string, 0)

	for i, line := range input {
		// first identify the first set of lines that are stacks only
		if line != "" {
			unparsedStacks = append(unparsedStacks, line)
		} else {
			dividingLine = i
			break
		}
	}

	// then move to capture the stacks
	for i, line := range input {
		// skip the stacks part (and divider) and capture all moves
		if i > dividingLine {
			unparsedMoves = append(unparsedMoves, line)
		}
	}
	// parse the stacks and then the moves
	stacks = parseStacks(unparsedStacks)
	moves = parseMoves(unparsedMoves)
	return stacks, moves
}

// Parse out the stacks and return them as a two-dimensional array The general strategy
// is to parse the input backwards (bottom to top) so that we can easily identify how
// many stacks we have. Create a top level array for each one and then proceed to find
// crates and append them to the correct stack..
// Example input: [          [D]     [N] [C]     [Z] [M] [P]  1   2   3 ]
func parseStacks(input []string) (stacks [][]string) {
	for i := len(input) - 1; i >= 0; i-- {
		if i == len(input)-1 {
			// header row
			for _, colValue := range strings.Split(input[i], " ") {
				// this is where we'll build the []string for each stack and drop them into a larger array for the return
				if strings.TrimSpace(colValue) != "" {
					stacks = append(stacks, []string{})
				}
			}
		} else {
			for j := 0; j < len(input[i]); j++ {
				//fmt.Printf("Row: %2d/%2d Col: %2d/%2d Value: %q", i, len(input), j, len(input[i]), input[i][j])
				colValue := string(input[i][j])
				if (j-1)%4 == 0 || j == 1 {
					if strings.TrimSpace(colValue) == "" {
						continue
					} else {
						stackNum := (j - 1) / 4
						stacks[stackNum] = append(stacks[stackNum], colValue)
					}
				} else {
					continue
				}
			}
		}
	}
	return
}

// parseMoves: Parse out the components of the move to make craneOperator simpler to implement
func parseMoves(input []string) []Move {
	moves := make([]Move, 0)
	for _, line := range input {
		moveParts := strings.Split(line, " ")
		qty, err := strconv.Atoi(moveParts[1])
		if err != nil {
			os.Exit(1)
		}
		from, err := strconv.Atoi(moveParts[3])
		if err != nil {
			os.Exit(1)
		}
		to, err := strconv.Atoi(moveParts[5])
		if err != nil {
			os.Exit(1)
		}
		curMove := Move{qty: qty, from: from, to: to}
		moves = append(moves, curMove)
	}
	return moves
}

// this will actually manipulate the stacks according to the listed moves
func craneOperator(stacks [][]string, moves []Move) [][]string {
	return stacks
}

// just checking the tops of each stack and return it as a string
func checkTops(stacks [][]string) string {
	tops := make([]string, 0)
	for _, curStack := range stacks {
		tops = append(tops, curStack[len(curStack)-1])
	}
	return strings.Join(tops, "")
}

// the main function the first part of the challenge goes through
// first we parse the input file and separate the stacks from the moves
// then we step through the moves and apply them to the stacks
// then we check the tops and return them
func DayFivePartOne(input []string) string {
	var stacks [][]string
	var moves []Move
	stacks, moves = parseInput(input)
	finalStacks := craneOperator(stacks, moves)
	tops := checkTops(finalStacks)
	return tops
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

/*********************************************************
* I'm abandoning this particular implementation of the   *
* linked list as I'm really struggling with it and just  *
* want to get moving again. I'm not deleting this as I'm *
* not sure if I want to come back and try again. No      *
* promises
**********************************************************/
/*
type Crate struct {
	contents string
	next     *Crate
}

type Stack struct {
	head   *Crate
	length int
}

func (s *Stack) printAll() {
	fmt.Printf("stack length: %v\n", s.length)
	if s.length <= 0 {
		fmt.Println("No members")
	}
	curCrate := s.head
	for i := 0; i < s.length; i++ {
		fmt.Printf("%d :: crate: %v\n", i, curCrate.contents)
		curCrate = curCrate.next
	}
}

func (s *Stack) pop() *Crate {
	if s.length >= 1 {
		retVal := s.head
		s.head = s.head.next
		return retVal
	} else {
		fmt.Println("no members to pop off of stack")
		return nil
	}
}

func (s *Stack) add(c *Crate) {
	c.next = s.head
	s.head = c
	s.length++
}
*/
