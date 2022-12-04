package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScoreShape(shape string) int {
	if shape == "X" {
		return 1
	}
	if shape == "Y" {
		return 2
	}
	if shape == "Z" {
		return 3
	}
	panic("didn't recognize the shape")
}

func ScoreOutcome(shapes []string) int {
	them := shapes[0]
	me := shapes[1]
	if them == "A" {
		if me == "X" {
			return 3
		}
		if me == "Y" {
			return 6
		}
		if me == "Z" {
			return 0
		}
	}
	if them == "B" {
		if me == "X" {
			return 0
		}
		if me == "Y" {
			return 3
		}
		if me == "Z" {
			return 6
		}
	}
	if them == "C" {
		if me == "X" {
			return 6
		}
		if me == "Y" {
			return 0
		}
		if me == "Z" {
			return 3
		}
	}
	panic("couldn't recognize one shape or another")
}

func NewScoreOutcome(outcome string) int {
	if outcome == "X" {
		return 0
	}
	if outcome == "Y" {
		return 3
	}
	if outcome == "Z" {
		return 6
	}
	panic("didn't recognize the outcome string")
}

func NewScoreShape(shapes []string) int {
	them := shapes[0]
	outcome := shapes[1]
	if them == "A" {
		if outcome == "X" {
			return 3
		}
		if outcome == "Y" {
			return 1
		}
		if outcome == "Z" {
			return 2
		}
	}
	if them == "B" {
		if outcome == "X" {
			return 1
		}
		if outcome == "Y" {
			return 2
		}
		if outcome == "Z" {
			return 3
		}
	}
	if them == "C" {
		if outcome == "X" {
			return 2
		}
		if outcome == "Y" {
			return 3
		}
		if outcome == "Z" {
			return 1
		}
	}
	panic("didn't recognize either the shape or the outcome")
}

func NewScoreRounds(rounds []string) int {
	totalScore := 0
	for _, moves := range rounds {
		moveArr := strings.Split(moves, " ")
		outcomeScore := NewScoreOutcome(moveArr[1])
		shapeScore := NewScoreShape(moveArr)
		totalScore += shapeScore + outcomeScore
	}
	return totalScore
}

func ScoreRounds(rounds []string) int {
	totalScore := 0
	for _, moves := range rounds {
		moveArr := strings.Split(moves, " ")
		shapeScore := ScoreShape(moveArr[1])
		outcomeScore := ScoreOutcome(moveArr)
		totalScore += shapeScore + outcomeScore
	}
	return totalScore
}

func main() {
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

	totalScore := ScoreRounds(fileLines)
	fmt.Printf("The total score is %v\n", totalScore)

	newTotalScore := NewScoreRounds(fileLines)
	fmt.Printf("The new total score is %v\n", newTotalScore)

}
