package kclient

import (
	"testing"
)

func TestCountDigits(t *testing.T) {
	testCases := []struct {
		name          string
		digitString   string
		expectedCount int
	}{
		{"empty string", "", 0},
		{"all digits", "654321", 6},
		{"mixed string", "12ab34de", 4},
		{"odd characters", " %1B^-?", 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var actualCount = countDigits(tc.digitString)
			if tc.expectedCount != actualCount {
				t.Errorf("expected count %d but actual was %d", tc.expectedCount, actualCount)
			}
		})
	}
}
