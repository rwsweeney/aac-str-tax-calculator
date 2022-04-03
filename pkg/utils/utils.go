package utils

import "fmt"

type AirbnbData struct {
	GrossEarnings float64
	Nights        int
	Aatax         float64
}

func ShowOutput(records [][]string, airbnbTaxData AirbnbData) {

	fmt.Println("\t\t\t Airbnb GE: \t # Nights: \t AATax:")
	for x := range records {
		if x != 0 {
			fmt.Println(records[x][5], "\t", records[x][14], "\t", records[x][4])
		}
	}
	fmt.Println("\t\t\t=======\t\t=======\t\t=======\n\t\t\t", airbnbTaxData.GrossEarnings, "\t\t", airbnbTaxData.Nights, "\t\t", airbnbTaxData.Aatax)

}
