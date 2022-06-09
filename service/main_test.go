package main

import "testing"

func TestAdd(t *testing.T) {
	someNumber := 2
	expectedResult := result()

	if someNumber+2 != expectedResult {
		t.Errorf("Failed expectation: %d + 2 != %d", someNumber, expectedResult)
	}
	if someNumber+3 != expectedResult+1 {
		t.Errorf("Failed expectation: %d + 3 != %d", someNumber, expectedResult+1)
	}
}

func result() int {
	return 4
}
