package main

import "testing"

func TestTimes(t *testing.T) {
	result := times(2, 4)

	if result != 8 {
		t.Errorf("times(2, 4) = %d; want 8", result)
	}
}
