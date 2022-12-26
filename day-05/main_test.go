package main

import (
	"testing"
)

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

func TestCraneOperator(t *testing.T) {
	// Generate some test data here
	move1 := Move{qty: 1, from: 2, to: 1}
	move2 := Move{qty: 3, from: 1, to: 3}
	move3 := Move{qty: 2, from: 2, to: 1}
	move4 := Move{qty: 1, from: 1, to: 2}
	sampleMoves := make([]Move, 0)
	sampleMoves = append(sampleMoves, move1)
	sampleMoves = append(sampleMoves, move2)
	sampleMoves = append(sampleMoves, move3)
	sampleMoves = append(sampleMoves, move4)

	var s1, s2, s3 []string
	var c1, c2, c3, c4, c5, c6 string
	c1 = "Z"
	c2 = "N"
	c3 = "M"
	c4 = "C"
	c5 = "D"
	c6 = "P"
	s1 = append(s1, c1)
	s1 = append(s1, c2)
	s2 = append(s2, c3)
	s2 = append(s2, c4)
	s2 = append(s2, c5)
	s3 = append(s3, c6)
	sampleStacks := [][]string{s1, s2, s3}

	var ws1, ws2, ws3 []string
	ws1 = append(ws1, c4)
	ws2 = append(ws2, c3)
	ws3 = append(ws3, c6)
	ws3 = append(ws3, c5)
	ws3 = append(ws3, c2)
	ws3 = append(ws3, c1)
	want := [][]string{ws1, ws2, ws3}

	got := craneOperator(sampleStacks, sampleMoves)
	assertStacksEqual(t, want, got)
}

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

func assertStacksEqual(t *testing.T, want, got [][]string) {
	t.Helper()
	if len(want) != len(got) {
		t.Errorf("different number of stacks: want %d got %d", len(want), len(got))
	}
	for i := 0; i < len(want); i++ {
		if len(want[i]) != len(got[i]) {
			t.Errorf("different number of crates: want %d got %d", len(want[i]), len(got[i]))
		}
		for j := 0; j < len(want[i]); j++ {
			if want[i][j] != got[i][j] {
				t.Errorf("incorrect crate contents: want %q got %q", want[i][j], got[i][j])
			}
		}
	}
}

func TestParseStacks(t *testing.T) {
	sample_input := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 "}

	var stack1, stack2, stack3 []string
	var crate1, crate2, crate3, crate4, crate5, crate6 string
	crate1 = "Z"
	crate2 = "N"
	crate3 = "M"
	crate4 = "C"
	crate5 = "D"
	crate6 = "P"
	stack1 = append(stack1, crate1)
	stack1 = append(stack1, crate2)
	stack2 = append(stack2, crate3)
	stack2 = append(stack2, crate4)
	stack2 = append(stack2, crate5)
	stack3 = append(stack3, crate6)
	want := [][]string{stack1, stack2, stack3}

	got := parseStacks(sample_input)

	assertStacksEqual(t, want, got)
}

//func TestDayFivePartTwo(t *testing.T) {
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
