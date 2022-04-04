package vrbofunc

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/rwsweeney/aac-str-tax-calculator/pkg/utils"
)

// Calculates the GrossEarnings, number of nights, and Anne Arrundel Occupancy Tax for VRBO
func CalculateVRBO() (taxData utils.TaxData) {
	// Process VRBO data
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

	_, aaTaxColumn := utils.GetColumn("Vrbo's Taxes | Taxes sent to Vrbo", vrboRecords)
	aaOccupancyTax, error := strconv.ParseFloat(vrboRecords[1][aaTaxColumn], 64)
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

	return vrboTaxData
}
