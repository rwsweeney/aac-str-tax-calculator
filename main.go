// TO DO
// Airbnb: Need to capture the # of nights, Gross Earnings, and Occupancy Tax
// Airbnb: Need to sum the nights, gross earnings, and OT.
// Airbnb: Should run a test calculation of (Gross Earnings * .06) + (Gross Earnings * .07) = occupancy tax value.

// All VRBO stuff

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/rwsweeney/aac-str-tax-calculator/pkg/airbnbfunc"
	"github.com/rwsweeney/aac-str-tax-calculator/pkg/utils"
)

func main() {

	airbnbFile, err := os.Open("airbnb_tax_return.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer airbnbFile.Close()

	airbnbReader := csv.NewReader(airbnbFile)
	airbnbRecords, err := airbnbReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	vrboFile, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer vrboFile.Close()

	vrboReader := csv.NewReader(vrboFile)
	vrboReader.FieldsPerRecord = -1
	vrboRecords, error := vrboReader.ReadAll()
	if error != nil {
		log.Fatal(error)
	}

	airbnbTaxData := utils.TaxData{
		GrossEarnings: airbnbfunc.CalculateGrossEarnings(airbnbRecords),
		Nights:        airbnbfunc.CalculateTotalNights(airbnbRecords),
		Aatax:         airbnbfunc.CalculateAATax(airbnbfunc.CalculateGrossEarnings(airbnbRecords)), // This runs CalculateGrossEarnings twice which is gross.
	}

	aaOccupancyTax, error := strconv.ParseFloat(vrboRecords[1][11], 64)
	if error != nil {
		log.Fatal(error)
	}

	vrboNights, error := strconv.Atoi(vrboRecords[1][5])
	if error != nil {
		log.Fatal(error)
	}

	vrboTaxData := utils.TaxData{
		GrossEarnings: aaOccupancyTax / float64(0.07), // This rounds floating point value up slightly.
		Nights:        vrboNights,
		Aatax:         aaOccupancyTax,
	}

	fmt.Println("Box #1: ", airbnbTaxData.Nights+vrboTaxData.Nights)
	fmt.Println("Box #5: ", airbnbTaxData.GrossEarnings+vrboTaxData.GrossEarnings)
	fmt.Println("Box #7: ", airbnbTaxData.GrossEarnings+vrboTaxData.GrossEarnings)
	fmt.Println("Box #9: ", airbnbTaxData.GrossEarnings+vrboTaxData.GrossEarnings)
	fmt.Println("Box #10: ", airbnbTaxData.Aatax+vrboTaxData.Aatax)
	fmt.Println("Box #13: ", airbnbTaxData.Aatax+vrboTaxData.Aatax)
	fmt.Println("Box #14: ", airbnbTaxData.Aatax)
	fmt.Println("Box #15: ", vrboTaxData.Aatax)
	//fmt.Println(airbnbTaxData)
	//utils.ShowOutput(airbnbRecords, airbnbTaxData)

}
