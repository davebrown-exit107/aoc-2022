package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func TopThreeCalories(calorieList []string) int {
	totalElves := 0
	curCals := 0
	var totalCals []int
	for _, cals := range calorieList {
		if cals == "" {
			totalCals = append(totalCals, curCals)
			totalElves++
			curCals = 0
		} else {
			int_cals, err := strconv.Atoi(cals)
			if err != nil {
				panic(err)
			}
			curCals += int_cals
		}
	}
	totalCals = append(totalCals, curCals)
	sort.Ints(totalCals)
	return totalCals[totalElves] + totalCals[totalElves-1] + totalCals[totalElves-2]
}

func TopCalories(calorieList []string) int {
	curCals := 0
	topCals := 0
	for _, cals := range calorieList {
		if cals == "" {
			// Check for top status
			if curCals > topCals {
				topCals = curCals
			}
			// Next elf up
			curCals = 0
		} else {
			int_cals, err := strconv.Atoi(cals)
			if err != nil {
				panic(err)
			}
			curCals += int_cals
		}
	}
	if curCals > topCals {
		topCals = curCals
	}
	return topCals
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

	topCals := TopCalories(fileLines)
	topThreeCals := TopThreeCalories(fileLines)

	fmt.Printf("The top elf has a calorie count of %v\n", topCals)
	fmt.Printf("The top three elves have a calorie count of %v\n", topThreeCals)

}
