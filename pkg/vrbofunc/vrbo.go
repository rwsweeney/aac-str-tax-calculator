package vrbofunc

import (
	"encoding/csv"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/rwsweeney/aac-str-tax-calculator/pkg/utils"
)

// Calculates the GrossEarnings, number of nights, and Anne Arrundel Occupancy Tax for VRBO
func CalculateVRBO(file string) (taxData utils.TaxData) {
	// Process VRBO data
	vrboFile, err := os.Open(file)
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
	log.Debug("VRBO: CSV loaded")

	_, aaTaxColumn := utils.GetColumn("Vrbo's Taxes | Taxes sent to Vrbo", vrboRecords)
	log.Debug("VRBO: aaTaxColumn acquired: ", aaTaxColumn)

	_, jurisdictionColumn := utils.GetColumn("Jurisdiction name", vrboRecords)
	log.Debug("VRBO: jurdistictionColumn acquired: ", jurisdictionColumn)

	aaOccupancyTax := CalculateGrossTaxes(aaTaxColumn, jurisdictionColumn, vrboRecords)
	if error != nil {
		log.Fatal(error)
	}

	_, nightsColumn := utils.GetColumn("Nights", vrboRecords)
	log.Debug("VRBO - nightsColumn acquired: ", nightsColumn)

	vrboNights := CalculateTotalNights(nightsColumn, jurisdictionColumn, vrboRecords)
	if error != nil {
		log.Fatal(error)
	}

	vrboTaxData := utils.TaxData{
		// Since we know how much tax we paid we're reversing it to figure out the grossEarnings since the CSV doesn't have it.
		GrossEarnings: aaOccupancyTax / float64(0.08), // This rounds floating point value up slightly.
		Nights:        vrboNights,
		Aatax:         aaOccupancyTax,
	}
	log.Debug("VRBO - vrboTaxData computed (grossEarnings, Nights, Aatax): ", vrboTaxData)

	return vrboTaxData
}

/*
	Calculates the amount of tax paid to Anne Arrundel County for VRBO. This function takes in an int for the location

of the tax column and the jurisdiction column along with the CSV records in order to return the Gross Taxes to AA.
*/
func CalculateGrossTaxes(columnTax, columnJurisdiction int, records [][]string) float64 {

	var grossTaxes float64

	for x := range records {
		if x != 0 {
			//fmt.Println(records[x][columnJurisdiction])
			if records[x][columnJurisdiction] != "ANNE ARUNDEL" {
				continue
			}
			singlePayment, error := strconv.ParseFloat(records[x][columnTax], 64)
			grossTaxes += singlePayment
			// fmt.Printf("CalculateGrossEarnings loop #%d\n\tvalue: %f singleField: %f\n\n", x, grossTaxes, singlePayment)
			if error != nil {
				log.Fatal(error)
			}

		}
	}

	log.Debug("VRBO - grossTaxes computed: ", grossTaxes)

	return grossTaxes
}

// Calculates the amount of tax paid to Anne Arrundel County for VRBO
func CalculateTotalNights(columnNight, columnJurisdiction int, records [][]string) int {

	var totalNights int

	for x := range records {
		if x != 0 {
			if records[x][columnJurisdiction] != "ANNE ARUNDEL" {
				continue
			}
			singleNight, error := strconv.Atoi(records[x][columnNight])
			//fmt.Println(singleNight)
			totalNights += singleNight
			if error != nil {
				log.Fatal(error)
			}

		}
	}
	log.Debug("VRBO: totalNights computed: ", totalNights)

	return totalNights
}
