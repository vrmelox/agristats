package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func main() {
	file := flag.String("file", "", "The data file path")
	crop := flag.String("crop", "all", "The crop you want to analyse specifly")
	year := flag.String("year", "all", "The year in this format YYYY-MM-DD")
	report := flag.String("report", "text", "The output format json|text")
	if len(os.Args) < 8 {
		fmt.Println("Arguments insuffisants. Merci de faire --help pour voir les arguments")
		os.Exit(84)
	}

	flag.Parse()

	File := *file
	Crop := *crop
	Year := *year
	Report := *report

	harvis := retrieveData(File)
	cropHarvest := cropFilter(harvis, Crop)
	i, err := strconv.Atoi(Year)
	if err != nil {
		log.Fatal("Donnez une année")
	}
	yearHarvest := yearFilter(cropHarvest, i)
	endStats := orderCrops(setCrops(yearHarvest))
	if Report == "text" {
		s := strings.Repeat("=", 10)
		fmt.Printf("%s\nAGRISTATS ANALYSIS REPORT\n%s\n", s, s)
		peri := fmt.Sprintf("Period: %s-01-01 to %s-12-31", Year, Year)
		fmt.Println(peri)
		totalHarvest := fmt.Sprintf("Total Harvests: %d", len(endStats))
		fmt.Println(totalHarvest)
		if (Crop == "all") {
			fmt.Println("CROPS by YIELD:")
			for i, cro := range endStats {
				se := fmt.Sprintf("%d.%s        : %.2f tons/hectare (avg)", i, cro.Cro, cro.Tons/cro.Hectares)
				fmt.Println(se)
			}
			return
		}
		se := fmt.Sprintf("%d.%s        : %.2f tons/hectare (avg)", i, endStats[0].Cro, endStats[0].Tons/endStats[0].Hectares)
		fmt.Println(se)
	}
}