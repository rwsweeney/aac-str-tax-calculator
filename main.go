// TO DO
// Airbnb: Need to capture the # of nights, Gross Earnings, and Occupancy Tax
// Airbnb: Need to sum the nights, gross earnings, and OT.
// Airbnb: Should run a test calculation of (Gross Earnings * .06) + (Gross Earnings * .07) = occupancy tax value.

// All VRBO stuff

package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/rwsweeney/aac-str-tax-calculator/pkg/airbnbfunc"
	"github.com/rwsweeney/aac-str-tax-calculator/pkg/utils"
)

func main() {

	file, err := os.Open("airbnb_tax_return.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	airbnbTaxData := utils.AirbnbData{
		GrossEarnings: airbnbfunc.CalculateGrossEarnings(records),
		Nights:        airbnbfunc.CalculateTotalNights(records),
		Aatax:         airbnbfunc.CalculateAATax(airbnbfunc.CalculateGrossEarnings(records)), // This runs CalculateGrossEarnings twice which is gross.
	}

	utils.ShowOutput(records, airbnbTaxData)
}
