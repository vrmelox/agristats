package main

import (
	"flag"
	"fmt"
)
func main() {
	file := flag.String("file", "", "The data file path")
	crop := flag.String("crop", "all", "The crop you want to analyse specifly")
	year := flag.String("year", "all", "The year in this format YYYY-MM-DD")
	report := flag.String("report", "text", "The output format json|text")

	flag.Parse()

	File := *file
	Crop := *crop
	Year := *year
	Report := *report

	fmt.Println(Crop)
	fmt.Println(Year)
	fmt.Println(Report)
	harvs := retrieveData(File)
	fmt.Println(harvs)
}