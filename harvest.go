package main

import (
	"time"
)

type Harvest struct {
	Date time.Time
	crop string
	hectares float64
	yields_tons float64
	rainfall_mm int32
	temperature_avg int32
}


