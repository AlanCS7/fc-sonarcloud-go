package main

import "testing"

func TestTimes(t *testing.T) {
	result := times(2, 3)

	if result != 6 {
		t.Errorf("times(2, 3) = %d; want 6", result)
	}
}
