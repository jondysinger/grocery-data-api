package kclient

import (
	"strconv"
	"strings"
)

// Counts the number of digits in a string
func countDigits(numberString string) (digits int) {
	for _, char := range strings.Split(numberString, "") {
		if _, err := strconv.Atoi(char); err == nil {
			digits++
		}
	}
	return digits
}
