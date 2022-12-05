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

//
//func TestNewScoreRounds(t *testing.T) {
//	sample_input := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"}
//
//	want := 12
//	got := NewScoreRounds(sample_input)
//
//	if want != got {
//		t.Errorf("expected %d, got %d.", want, got)
//	}
//}
