package main

import (
	"testing"
)

/*
func TestDayFivePartOne(t *testing.T) {
	sample_input := []string{
		"	   [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2"}

	want := "CMZ"
	got := DayFivePartOne(sample_input)

	if want != got {
		t.Errorf("expected %v, got %v.", want, got)
	}
}
*/

func assertMovesEqual(t testing.TB, want, got []Move) {
	t.Helper()
	if len(want) != len(got) {
		t.Errorf("length incorrect, expected %d got %d", len(want), len(got))
	}
	for i := 0; i < len(want); i++ {
		if want[i].qty != got[i].qty {
			t.Errorf("qty incorrect: want %d got %d", want[i].qty, got[i].qty)
		}
		if want[i].from != got[i].from {
			t.Errorf("from incorrect: want %d got %d", want[i].from, got[i].from)
		}
		if want[i].to != got[i].to {
			t.Errorf("to incorrect: want %d got %d", want[i].to, got[i].to)
		}
	}
}

func TestParseMoves(t *testing.T) {
	sample_input := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2"}

	// I'm not sure how I'll do this other than just generating it manually and copy/pasting here
	move1 := Move{qty: 1, from: 2, to: 1}
	move2 := Move{qty: 3, from: 1, to: 3}
	move3 := Move{qty: 2, from: 2, to: 1}
	move4 := Move{qty: 1, from: 1, to: 2}
	want := make([]Move, 0)
	want = append(want, move1)
	want = append(want, move2)
	want = append(want, move3)
	want = append(want, move4)

	got := parseMoves(sample_input)

	assertMovesEqual(t, want, got)
}

/*
func TestParseStacks(t *testing.T) {
	sample_input := []string{
		"	   [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2"}

	// I'm not sure how I'll do this other than just generating it manually and copy/pasting here
	var stack1, stack2, stack3 stack
	var crate1, crate2, crate3, crate4, crate5, crate6 crate
	crate1.contents = "Z"
	crate2.contents = "N"
	crate3.contents = "M"
	crate4.contents = "C"
	crate5.contents = "D"
	crate6.contents = "P"
	stack1.add(&crate1)
	stack1.add(&crate2)
	stack2.add(&crate3)
	stack2.add(&crate4)
	stack2.add(&crate5)
	stack3.add(&crate6)
	want := []stack{stack1, stack2, stack3}
	got := parseStacks(sample_input)
	fmt.Println(want, got)
	fmt.Println(want[0].head.contents)

	if reflect.DeepEqual(want, got) {
		t.Errorf("expected %v, got %v.", want, got)
	}
}
*/

//func TestCountOverlappingAssignments(t *testing.T) {
//	sample_input := []string{
//		"2-4,6-8",
//		"2-3,4-5",
//		"5-7,7-9",
//		"2-8,3-7",
//		"6-6,4-6",
//		"2-6,4-8"}
//
//	want := 4
//	got := CountOverlappingAssignments(sample_input)
//
//	if want != got {
//		t.Errorf("expected %d, got %d.", want, got)
//	}
//}
