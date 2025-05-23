package somecode

import (
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Basic cases
		{"racecar", true},
		{"hello", false},
		{"A man, a plan, a canal: Panama", true},
		{"Not a palindrome", false},

		// Corner cases
		{"", true},                           // Empty string
		{"a", true},                          // Single character
		{".,!?", true},                       // Only non-alphabetic characters
		{"RaCeCaR", true},                    // Mixed case
		{"  ", true},                         // Only spaces
		{strings.Repeat("a", 100000), true},  // Very long palindrome
		{strings.Repeat("ab", 50000), false}, // Very long non-palindrome
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
