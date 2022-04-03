package airbnbfunc

import (
	"fmt"
	"strconv"
)

// This function calculates the gross earnings for AirBNB for the month
func CalculateGrossEarnings(records [][]string) float64 {

	var grossEarnings float64

	for x := range records {
		if x != 0 {
			singleEarning, error := strconv.ParseFloat(records[x][14], 64)
			grossEarnings += singleEarning
			if error != nil {
				fmt.Println(error)
			}

		}
	}
	return grossEarnings
}

// Calculates the occupancy tax for AirBNB to Anne Arrundel County
func CalculateAATax(grossEarnings float64) float64 {

	aaTax := grossEarnings * 0.07

	return aaTax
}

func CalculateTotalNights(records [][]string) int {

	var totalNights int

	for x := range records {
		if x != 0 {
			singleNight, error := strconv.Atoi(records[x][4])
			totalNights += singleNight
			if error != nil {
				fmt.Println(error)
			}

		}
	}
	return totalNights
}

// This function iterates over the CSV file to find the Gross Earnings row and column position.
func Get_earnings_column(records [][]string) (row, column int) {
	for x, y := range records {
		for a := 0; a < len(y); a++ {
			if records[x][a] == "Gross Earnings" {
				row := x
				column := a
				return row, column
			}
		}
	}
	return row, column
}
