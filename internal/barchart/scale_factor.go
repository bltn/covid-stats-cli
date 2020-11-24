package barchart

func CalculateScaleFactor(bars []Bar, desiredValue float64) float64 {
	highestValue := 0.0
	for _, bar := range bars {
		if float64(bar.Count()) > highestValue {
			highestValue = float64(bar.Count())
		}
	}

	scaleFactor := 1.0

	for highestValue * scaleFactor > desiredValue {
		scaleFactor = scaleFactor / 2.0
	}

	return scaleFactor
}
