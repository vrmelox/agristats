package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

// type Harvest struct {
// 	Date time.Time
// 	crop string
// 	hectares float64
// 	yields_tons float64
// 	rainfall_mm int32
// 	temperature_avg int32
// }

func harvParser(record []string) Harvest {
	harv := Harvest{}
	date, err := time.Parse("2006-01-02", record[0])
	if err == nil {
		harv.Date = date
	}
	harv.crop = record[1]
	hect, err := strconv.ParseFloat(record[2], 64)
	if err == nil {
		harv.hectares = hect
	}
	yield, err := strconv.ParseFloat(record[3], 64)
	if err == nil {
		harv.yields_tons = yield
	}
	rainfa, err := strconv.Atoi(record[4])
	if err == nil {
		harv.rainfall_mm = int32(rainfa)
	}
	tempera, err := strconv.Atoi(record[5])
	if err == nil {
		harv.temperature_avg = int32(tempera)
	}
	return harv
}

func retrieveData(path string) []Harvest {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	harvis := make([]Harvest, 0)
	i := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		if i > 0 {
			harv := harvParser(record)
			harvis = append(harvis, harv)
		}
		i++
	}
	return harvis
}