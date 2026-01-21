package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Harvest struct {
	Date time.Time
	crop string
	yields_tons string
	rainfall_mm string
	temperature_avg string
}

func dateGenerator() time.Time {
	rand.Seed(time.Now().UnixNano())

	start := time.Date(1958, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	end := time.Date(2026, 1, 21, 0, 0, 0, 0, time.UTC).Unix()
	randomDate := time.Unix(rand.Int63n(end-start) + start, 0).UTC()
	return randomDate
}

func valueGenerator(start, end float64) float64 {
	rand.Seed(time.Now().UnixNano())

	randomFloat := start + rand.Float64()*(end-start)
	rounded := math.Round(randomFloat * 100) / 100
	return rounded
}

