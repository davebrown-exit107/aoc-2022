package main

import "testing"

func TestCalorieCounter(t *testing.T) {
	sample_input := []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"}

	got_elf, got_cals := CalorieCalculator(sample_input)
	expect_elf, expect_cals := 4, 24000

	if got_elf != expect_elf {
		t.Errorf("expected '%d' but got '%d'", expect_elf, got_elf)
	}
	if got_cals != expect_cals {
		t.Errorf("expected '%d' but got '%d'", expect_cals, got_cals)
	}

}
