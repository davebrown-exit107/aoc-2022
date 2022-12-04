package main

import (
	"bufio"
	"fmt"
	"os"
)

func ScoreRounds(rounds []string) int {
	return 0
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

	fmt.Printf("The total score is %v", totalScore)

}
