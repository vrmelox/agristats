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

type Stats struct {
	Cro string
	Year int
	Max float64
	Min float64
	Harvestotal int64
	Tons float64
	Hectares float64
	Rainfall int32
	Moy float64
	Ratio float64
}

type ByRatio []Stats

func (a ByRatio) Len() int {
	return len(a)
}

func (a ByRatio) Less(i, j int) bool {
	return a[i].Ratio < a[j].Ratio
}

func (a ByRatio) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}



