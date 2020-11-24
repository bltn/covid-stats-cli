package coviddata

import "time"

type data struct {
	date   time.Time
	cases  int
	deaths int
}
