package main

import "testing"

func TestScoreRounds(t *testing.T) {
	sample_input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}

	want := 157
	got := SumDuplicates(sample_input)

	if want != got {
		t.Errorf("expected %d, got %d.", want, got)
	}
}

func TestSumBadges(t *testing.T) {
	sample_input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}

	want := 70
	got := SumBadges(sample_input)

	if want != got {
		t.Errorf("expected %d, got %d.", want, got)
	}
}
