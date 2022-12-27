package main

import (
	"testing"
)

func TestDaySixPartOne(t *testing.T) {
	sample_input := map[int]string{
		7:  "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		5:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
		6:  "nppdvjthqldpwncqszvftbrmjlhg",
		10: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		11: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}

	for want, input := range sample_input {
		got := DaySixPartOne(input)

		if want != got {
			t.Errorf("expected %v, got %v.", want, got)
		}
	}
}

func TestDaySixPartTwo(t *testing.T) {
	sample_input := map[int]string{
		19: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		23: "bvwbjplbgvbhsrlpgdmjqwftvncz",
		//23:  "nppdvjthqldpwncqszvftbrmjlhg", // Removed this one as I don't want to deal with the map/key collision
		29: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		26: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}

	for want, input := range sample_input {
		got := DaySixPartTwo(input)

		if want != got {
			t.Errorf("expected %v, got %v.", want, got)
		}
	}
}
