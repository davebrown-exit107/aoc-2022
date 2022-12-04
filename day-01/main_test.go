package main

import "testing"

func TestTopThreeCalories(t *testing.T) {
	sample_input := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	got := TopThreeCalories(sample_input)
	expected := 45000

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}

}

func TestTopCalories(t *testing.T) {
	sample_input := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	got := TopCalories(sample_input)
	expected := 24000

	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}

}
