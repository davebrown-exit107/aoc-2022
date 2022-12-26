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

/*
func TestDaySixPartTwo(t *testing.T) {
	sample_input := map[int]string{
		7:  "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		5:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
		6:  "nppdvjthqldpwncqszvftbrmjlhg",
		10: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		11: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}

	for want, input := range sample_input {
		got := DaySixPartTwo(input)

		if want != got {
			t.Errorf("expected %v, got %v.", want, got)
		}
	}
}
*/
