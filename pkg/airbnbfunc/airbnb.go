package airbnbfunc

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

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
	log.Debug("Airbnb: CSV loaded")

	_, grossEarningsColumn := utils.GetColumn("Gross Earnings", airbnbRecords)
	log.Debug("Airbnb: grossEarnings Column acquired: ", grossEarningsColumn)

	_, nightsColumn := utils.GetColumn("Nights", airbnbRecords)
	log.Debug("Airbnb: nights Column acquired: ", nightsColumn)

	airbnbTaxData := utils.TaxData{
		GrossEarnings: CalculateGrossEarnings(grossEarningsColumn, airbnbRecords),
		Nights:        CalculateTotalNights(nightsColumn, airbnbRecords),
		Aatax:         CalculateAATax(CalculateGrossEarnings(grossEarningsColumn, airbnbRecords)), // This runs CalculateGrossEarnings twice which is gross.
	}

	log.Debug("Airbnb: Tax Data computed:", airbnbTaxData)
	return airbnbTaxData
}

// This function calculates the gross earnings for AirBNB for the month
func CalculateGrossEarnings(column int, records [][]string) float64 {

	var grossEarnings float64

	for x := range records {
		if x != 0 {
			singleEarning, error := strconv.ParseFloat(records[x][column], 64)
			grossEarnings += singleEarning
			if error != nil {
				fmt.Println(error)
			}

		}
	}

	log.Debug("Airbnb: grossEarnings computed: ", grossEarnings)

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

	log.Debug("Airbnb: totalNights computed: ", totalNights)

	return totalNights
}
