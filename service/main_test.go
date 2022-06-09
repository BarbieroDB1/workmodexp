package main

import "testing"

func TestAdd(t *testing.T) {
	someNumber := 2
	expectedResult := result()

	if someNumber+2 != expectedResult {
		t.Errorf("Failed expectation: %d + 2 != %d", someNumber, expectedResult)
	}
}

func result() int {
	return 4
}
