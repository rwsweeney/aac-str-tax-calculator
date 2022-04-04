package utils

import "fmt"

type TaxData struct {
	GrossEarnings float64
	Nights        int
	Aatax         float64
}

// This function iterates over the CSV file to find the Gross Earnings row and column position.
func GetColumn(name string, records [][]string) (row, column int) {
	for x, y := range records {
		for a := 0; a < len(y); a++ {
			if records[x][a] == name {
				row := x
				column := a
				return row, column
			}
		}
	}
	return row, column
}

func ShowOutput(records [][]string, airbnbTaxData TaxData) {

	fmt.Println("\t\t\t Airbnb GE: \t # Nights: \t AATax:")
	for x := range records {
		if x != 0 {
			fmt.Println(records[x][5], "\t", records[x][14], "\t", records[x][4])
		}
	}
	fmt.Println("\t\t\t=======\t\t=======\t\t=======\n\t\t\t", airbnbTaxData.GrossEarnings, "\t\t", airbnbTaxData.Nights, "\t\t", airbnbTaxData.Aatax)

}
