package main

import (
	"math"
)

type Seat struct {
    code string
}

// Parse the first 7 symbols of code to determine
// seat row
func (s Seat) Row() int {
	var diff float64
	var min float64 = 0
	var max float64 = 127

	for i := 0; i < 7; i++ {
		diff = max - min

		if (s.code[i:(i + 1)] == "F") {
			max = max - math.Floor(diff / 2)

			continue
		}

		min = min + math.Ceil(diff / 2)
	}

	return int(min)
}

// Parse the last 3 symbols of code to determine
// seat col
func (s Seat) Column() int {
	var diff float64
	var min float64 = 0
	var max float64 = 7

	for i := 7; i < 10; i++ {
		diff = max - min

		if (s.code[i:(i + 1)] == "L") {
			max = max - math.Floor(diff / 2)

			continue
		}

		min = min + math.Ceil(diff / 2)
	}

	return int(min)
}

// Calculate seat ID
func (s Seat) Id() int {
	return (s.Row() * 8) + s.Column();
}