package airbnbfunc

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/rwsweeney/aac-str-tax-calculator/pkg/utils"
)

// Calculates the GrossEarnings, number of nights, and Anne Arrundel Occupancy Tax for Airbnb
func CalculateAirbnb(file string) (airbnbTax utils.TaxData) {
	// Process Airbnb data
	airbnbFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer airbnbFile.Close()

	airbnbReader := csv.NewReader(airbnbFile)
	airbnbRecords, err := airbnbReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	_, grossEarningsColumn := utils.GetColumn("Gross Earnings", airbnbRecords)
	_, nightsColumn := utils.GetColumn("Nights", airbnbRecords)

	airbnbTaxData := utils.TaxData{
		GrossEarnings: CalculateGrossEarnings(grossEarningsColumn, airbnbRecords),
		Nights:        CalculateTotalNights(nightsColumn, airbnbRecords),
		Aatax:         CalculateAATax(CalculateGrossEarnings(grossEarningsColumn, airbnbRecords)), // This runs CalculateGrossEarnings twice which is gross.
	}
	return airbnbTaxData
}

// This function calculates the gross earnings for AirBNB for the month
func CalculateGrossEarnings(column int, records [][]string) float64 {

	var grossEarnings float64

	for x := range records {
		if x != 0 {
			singleEarning, error := strconv.ParseFloat(records[x][column], 64)
			grossEarnings += singleEarning
			//fmt.Printf("CalculateGrossEarnings loop #%d\n\tvalue: %f singleField: %f\n\n", x, grossEarnings, singleEarning)
			if error != nil {
				fmt.Println(error)
			}

		}
	}
	//fmt.Println(grossEarnings)
	return grossEarnings
}

// Calculates the occupancy tax for AirBNB to Anne Arrundel County
func CalculateAATax(grossEarnings float64) float64 {

	aaTax := grossEarnings * 0.07

	return aaTax
}

func CalculateTotalNights(column int, records [][]string) int {

	var totalNights int

	for x := range records {
		if x != 0 {
			singleNight, error := strconv.Atoi(records[x][column])
			//fmt.Println(singleNight)
			totalNights += singleNight
			if error != nil {
				fmt.Println(error)
			}

		}
	}
	return totalNights
}
