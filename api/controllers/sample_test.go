package controller

import "testing"

func TestSample(t *testing.T) {
	result := 1 + 1
	expected := 2

	if result != expected {
		t.Errorf("1 + 1 = %d; want %d", result, expected)
	}
}
