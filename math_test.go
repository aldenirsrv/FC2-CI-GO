package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(15, 15)

	if total != 30 {
		t.Errorf("Sum result is not valid: Result %d. Expected: %d", total, 30)
	}
}
