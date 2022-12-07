package main

import "testing"

func TestScoreRounds(t *testing.T) {
	sample_input := []string{"A Y", "B X", "C Z"}

	want := 15
	got := ScoreRounds(sample_input)

	if want != got {
		t.Errorf("expected %d, got %d.", want, got)
	}
}

func TestNewScoreRounds(t *testing.T) {
	sample_input := []string{"A Y", "B X", "C Z"}

	want := 12
	got := NewScoreRounds(sample_input)

	if want != got {
		t.Errorf("expected %d, got %d.", want, got)
	}
}
