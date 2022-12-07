package main

import "testing"

func TestCountFullyOverlappingAssignments(t *testing.T) {
	sample_input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8"}

	want := 2
	got := CountFullyOverlappingAssignments(sample_input)

	if want != got {
		t.Errorf("expected %d, got %d.", want, got)
	}
}

func TestCountOverlappingAssignments(t *testing.T) {
	sample_input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8"}

	want := 4
	got := CountOverlappingAssignments(sample_input)

	if want != got {
		t.Errorf("expected %d, got %d.", want, got)
	}
}
