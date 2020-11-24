package barchart

import (
	"testing"
)

func TestNewBarTrimsLabelIfLongerThanFiveChars(t *testing.T) {
	bar := NewBar("123456", 500)

	if bar.Label() != "12345" {
		t.Fatalf("Label '%s' should have been trimmed to 5 characters", bar.Label())
	}
}

func TestNewBarAddsPaddingToLabelsWithLessThanFourChars(t *testing.T) {
	bar := NewBar("1234", 500)
	if bar.Label() != "1234 " {
		t.Fatalf("Label '%s' should have been whitespace-padded to five chars", bar.Label())
	}

	bar = NewBar("123", 500)
	if bar.Label() != " 123 " {
		t.Fatalf("Label '%s' should have been whitespace-padded to five chars", bar.Label())
	}

	bar = NewBar("1", 500)
	if bar.Label() != "  1  " {
		t.Fatalf("Label '%s' should have been whitespace-padded to five chars", bar.Label())
	}

	bar = NewBar("", 500)
	if bar.Label() != "     " {
		t.Fatalf("Label '%s' should have been whitespace-padded to five chars", bar.Label())
	}
}