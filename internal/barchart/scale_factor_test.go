package barchart

import (
	"testing"
)

func TestCalculateScaleFactorForOneBarWithDesiredValueGreaterThanCount(t *testing.T) {
	bar := NewBar("something useful", 10)
	bars := make([]Bar, 1)
	bars = append(bars, bar)

	scaleFactor := CalculateScaleFactor(bars, 100.0)

	if scaleFactor != 1.0 {
		t.Fatalf("scaleFactor should be 1.0 but was %f", scaleFactor)
	}
}

func TestCalculateScaleFactorForOneBarWithDesiredValueEqualToCount(t *testing.T) {
	bar := NewBar("something useful", 10)
	bars := make([]Bar, 1)
	bars = append(bars, bar)

	scaleFactor := CalculateScaleFactor(bars, 10.0)

	if scaleFactor != 1.0 {
		t.Fatalf("scaleFactor should be 1.0 but was %f", scaleFactor)
	}
}

func TestCalculateScaleFactorForOneBarWithDesiredValueLowerThanCount(t *testing.T) {
	bar := NewBar("something useful", 100)
	bars := make([]Bar, 1)
	bars = append(bars, bar)

	scaleFactor := CalculateScaleFactor(bars, 10.0)

	if scaleFactor != 0.0625 {
		t.Fatalf("scaleFactor should be 0.0625 but was %f", scaleFactor)
	}
}

func TestCalculateScaleFactorForMultipleBarsWithDesiredValueGreaterThanHighestCount(t *testing.T) {
	bars := make([]Bar, 3)
	bars = append(bars, NewBar("1st Jan", 100))
	bars = append(bars, NewBar("2nd Jan", 200))
	bars = append(bars, NewBar("3rd Jan", 300))

	scaleFactor := CalculateScaleFactor(bars, 1000.0)

	if scaleFactor != 1.0 {
		t.Fatalf("scaleFactor should be 1.0 but was %f", scaleFactor)
	}
}

func TestCalculateScaleFactorForMultipleBarsWithDesiredValueEqualToHighestCount(t *testing.T) {
	bars := make([]Bar, 3)
	bars = append(bars, NewBar("1st Jan", 100))
	bars = append(bars, NewBar("2nd Jan", 200))
	bars = append(bars, NewBar("3rd Jan", 300))

	scaleFactor := CalculateScaleFactor(bars, 300.0)

	if scaleFactor != 1.0 {
		t.Fatalf("scaleFactor should be 1.0 but was %f", scaleFactor)
	}
}

func TestCalculateScaleFactorForMultipleBarsWithDesiredValueLessThanHighestCount(t *testing.T) {
	bars := make([]Bar, 3)
	bars = append(bars, NewBar("1st Jan", 100))
	bars = append(bars, NewBar("2nd Jan", 200))
	bars = append(bars, NewBar("3rd Jan", 300))

	scaleFactor := CalculateScaleFactor(bars, 10.0)

	if scaleFactor != 0.03125 {
		t.Fatalf("scaleFactor should be 1.0 but was %f", scaleFactor)
	}
}
