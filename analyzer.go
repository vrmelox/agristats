package main

import (
	"strings"
	"time"
)

func cropFilter(harvis []Harvest, crop string) []Harvest {
	harvs := make([]Harvest, 0)

	for _, harv := range harvis {
		if strings.EqualFold(harv.crop, crop) {
			harvs = append(harvs, harv)
		}
	}
	return harvs
}

func yearFilter(harvis []Harvest, year int) []Harvest {
	harvs := make([]Harvest, 0)

	start := time.Date(year, 01, 01, 0, 0, 0, 0, time.UTC).Unix()
	end := time.Date(year, 12, 31, 23, 59, 59,0, time.UTC).Unix()

	min := (time.Unix(start, 0).UTC())
	max := (time.Unix(end, 0).UTC())

	for _, h := range harvis {
		if !h.Date.Before(min) && !h.Date.After(max) {
			harvs = append(harvs, h)
		}
	}
	return harvs
}

func setCrops(harvis []Harvest, ) map[string]struct{} {
	
}


