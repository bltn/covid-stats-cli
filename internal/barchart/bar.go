package barchart

import "strings"

type Bar struct {
	label string
	count int
}

func (b Bar) Label() string {
	return b.label
}

func (b Bar) Count() int {
	return b.count
}

func NewBar(label string, count int) Bar {
	maxSize := 5

	// trim the label if it's > 5 chars long
	if len(label) > maxSize {
		return Bar{label[:maxSize], count}
	}

	// otherwise add padding so the label is 5 chars long
	leftPadding, rightPadding := getPadding(label, maxSize)
	return Bar{strings.Repeat(" ", leftPadding) + label + strings.Repeat(" ", rightPadding), count}
}

func getPadding(label string, maxSize int) (left int, right int) {
	totalPadding := maxSize - len(label)
	left, right = totalPadding / 2, totalPadding / 2
	if totalPadding % 2 != 0 {
		right = right + 1
	}
	return left, right
}

