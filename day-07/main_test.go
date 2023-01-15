package main

import (
	"testing"
)

func TestDaySevenPartOne(t *testing.T) {
	sample_input := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	got := DaySevenPartOne(sample_input)
	want := 95437

	if want != got {
		t.Errorf("expected %v, got %v.", want, got)
	}
}

/*
func TestDaySevenPartTwo(t *testing.T) {
	sample_input := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}

	got := DaySevenPartOne(sample_input)

	if want != got {
		t.Errorf("expected %v, got %v.", want, got)
	}
}
*/
