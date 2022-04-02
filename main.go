// TO DO
// Airbnb: Need to capture the # of nights, Gross Earnings, and Occupancy Tax
// Airbnb: Need to sum the nights, gross earnings, and OT.
// Airbnb: Should run a test calculation of (Gross Earnings * .06) + (Gross Earnings * .07) = occupancy tax value.

// All VRBO stuff

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type airbnbData struct {
	grossEarnings float64
	nights        int
	aatax         float64
}

func main() {

	file, err := os.Open("airbnb_tax_return.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	airbnb := airbnbData{
		grossEarnings: calculateGrossEarningsAirbnb(records),
	}

	fmt.Println("\t\t\t Airbnb Gross Earnings:")
	for x := range records {
		if x != 0 {
			fmt.Println(records[x][5], "\t", records[x][14])
		}
	}
	fmt.Println("\t\t\t=======\n\t\t\t", airbnb.grossEarnings)
}

// This function calculates the gross earnings for AirBNB for the month
func calculateGrossEarningsAirbnb(records [][]string) float64 {

	var grossEarnings float64

	for x := range records {
		if x != 0 {
			singleEarning, error := strconv.ParseFloat(records[x][14], 8)
			grossEarnings += singleEarning
			if error != nil {
				fmt.Println(error)
			}

		}
	}
	return grossEarnings
}

// This function iterates over the CSV file to find the Gross Earnings row and column position.
func get_earnings_column(records [][]string) (row, column int) {
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
