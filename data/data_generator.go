package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
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

func chooseCrop() string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(20)

	crops := []string{
		"Wheat",
		"Rice",
		"Maize",
		"Barley",
		"Oats",
		"Rye",
		"Sorghum",
		"Millet",
		"Soybean",
		"Lentil",
		"Chickpea",
		"Potato",
		"Sweet",
		"Cassava",
		"Yam",
		"Tomato",
		"Onion",
		"Pepper",
		"Coffee",
		"Peanut",
	}
	return crops[num]
}

func constructData(path string, dataNumber int32) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.New("File creation impossible")
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	_, err = writer.WriteString("id,date,crop,hectares,yield_tons,rainfall_mm,temperature_avg\n")
	if err != nil {
		return err
	}
	for i := range dataNumber {
		harvi := Harvest {
			Date: dateGenerator(),
			crop: chooseCrop(),
			hectares: valueGenerator(0.1, 5000.0),
			yields_tons: valueGenerator(0.0, 5000.0),
			rainfall_mm: int32(rand.Intn(3000)),
			temperature_avg: int32(valueGenerator(-10.0, 40)),
		}
		line := fmt.Sprintf("%d,%s,%s,%.2f,%.2f,%d,%d\n", i,harvi.Date.Format("2006-01-02"), harvi.crop, harvi.hectares, harvi.yields_tons, harvi.rainfall_mm, harvi.temperature_avg)
		_,err = writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
		if err != nil{
		return errors.New("Toutes les données n'ont pas pu être écrites")
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Au moins deux argumebts")
		return
	}
	path := os.Args[1]
	totalLine, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Passez un vrai nombre")
		return
	}
	err = constructData(path, int32(totalLine))
	if err != nil {
		fmt.Println(err)
	}
}