package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func cropFilter(harvis []Harvest, crop string) []Harvest {
	if crop == "all" {
		return harvis
	}
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
	end := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC).Unix()

	min := (time.Unix(start, 0).UTC())
	max := (time.Unix(end, 0).UTC())

	for _, h := range harvis {
		if !h.Date.Before(min) && !h.Date.After(max) {
			harvs = append(harvs, h)
		}
	}
	return harvs
}

func setCrops(harvis []Harvest) map[string]Stats {
	cropMaps := make(map[string]Stats)

	for _, cro := range harvis {
		stat := Stats{}
		crostats, ok := cropMaps[cro.crop]
		if !ok {
			stat.Cro = cro.crop
			stat.Year = cro.Date.Year()
			stat.Harvestotal = 1
			stat.Hectares = cro.hectares
			stat.Max = cro.yields_tons
			stat.Tons = cro.yields_tons
			stat.Rainfall = cro.rainfall_mm
			stat.Min = cro.yields_tons
			stat.Ratio = stat.Tons / stat.Hectares
			cropMaps[cro.crop] = stat
		} else {
			crostats.Harvestotal++
			crostats.Hectares += cro.hectares
			if crostats.Max < cro.yields_tons {
				crostats.Min = crostats.Max
				crostats.Max = cro.yields_tons
			} else {
				crostats.Min = cro.yields_tons
			}
			crostats.Tons += cro.yields_tons
			crostats.Rainfall += cro.rainfall_mm
			crostats.Ratio = crostats.Tons / crostats.Hectares
			cropMaps[cro.crop] = crostats
		}
	}
	return cropMaps
}
func sortHarvis(harvis []Stats) []Stats {
    sort.Slice(harvis, func(i, j int) bool {
        return harvis[i].Ratio > harvis[j].Ratio
    })
    return harvis
}

func orderCrops(harvis map[string]Stats) []Stats {
	croStats := make([]Stats, 0)
	for _, cro := range harvis {
		croStats = append(croStats, cro)
	}
	sortHarvis(croStats)
	fmt.Println(croStats[0])
	return croStats
}


