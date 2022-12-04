package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CalorieCalculator(calorie_list []string) (int, int) {
	cur_elf := 1
	cur_elf_cals := 0
	top_elf := 0
	top_elf_cals := 0
	for _, cals := range calorie_list {
		if cals == "" {
			// Output elf status
			fmt.Printf("%+v: %+v\n", cur_elf, cur_elf_cals)
			// Check for top status
			if cur_elf_cals > top_elf_cals {
				top_elf_cals = cur_elf_cals
				top_elf = cur_elf
			}
			// Next elf up
			cur_elf++
			cur_elf_cals = 0
		} else {
			int_cals, err := strconv.Atoi(cals)
			if err != nil {
				panic(err)
			}
			cur_elf_cals += int_cals
		}
	}
	return top_elf, top_elf_cals
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

	topElf, topCals := CalorieCalculator(fileLines)

	fmt.Printf("The top elf (%v) has a calorie count of %v\n", topElf, topCals)

}
