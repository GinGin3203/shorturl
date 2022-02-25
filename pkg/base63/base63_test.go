package base63

import (
	"errors"
	"fmt"
	"testing"
)

func TestEncoder_Encode(t *testing.T) {
	testTable := []struct {
		expected  string
		num       int
		minLength int
	}{
		{
			expected: "APDLA",
			num: base*base*base*15 + // P
				base*base*3 + // D
				base*11 + // L
				0, // A
			minLength: 5,
		},
		{
			expected: "AAPDLA",
			num: base*base*base*15 + // P
				base*base*3 + // D
				base*11 + // L
				0, // A
			minLength: 6,
		},
		{
			expected: "Ky1_",
			num: base*base*base*10 + // K
				base*base*50 + // y
				base*53 + // 1
				base - 1, // _
			minLength: 4,
		},
	}
	for num, tt := range testTable {
		t.Run(fmt.Sprintf("(test_%d)", num), func(t *testing.T) {
			expected := tt.expected
			got := Encode(tt.num, tt.minLength)
			t.Logf("got: %s", got)
			if got != expected {
				t.Fatalf("expected %s, but got %s", expected, got)
			}
		})
	}
}

func TestEncoder_Decode(t *testing.T) {
	testTable := []struct {
		s           string
		expected    int
		errExpected error
	}{
		{
			s: "APDLA",
			expected: base*base*base*15 + // P
				base*base*3 + // D
				base*11 + // L
				0, // A
			errExpected: nil,
		},
		{
			s: "AAPDLA",
			expected: base*base*base*15 + // P
				base*base*3 + // D
				base*11 + // L
				0, // A
			errExpected: nil,
		},
		{
			s: "Ky1_",
			expected: base*base*base*10 + // K
				base*base*50 + // y
				base*53 + // 1
				base - 1, // _
			errExpected: nil,
		},
		{
			s:           "Hello, 世界",
			expected:    0,
			errExpected: ErrInvalidCharacters,
		},
	}
	for num, tt := range testTable {
		t.Run(fmt.Sprintf("(test_%d)", num), func(t *testing.T) {
			got, err := Decode(tt.s)
			t.Logf("got: %d", got)
			if err != nil && !errors.Is(err, tt.errExpected) {
				t.Logf("expected err: %v, but got: %v", tt.errExpected, err)
			}
			if got != tt.expected {
				t.Logf("expected: %d, but got: %d", tt.expected, got)
			}
		})
	}
}
