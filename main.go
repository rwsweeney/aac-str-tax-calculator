// TO DO
// Create a function to show the work behind the calculations.
// Airbnb: Should run a test calculation of (Gross Earnings * .06) + (Gross Earnings * .07) = occupancy tax value.
// Let the app take in a debug flag

package main

import (
	"fmt"

	"github.com/rwsweeney/aac-str-tax-calculator/pkg/airbnbfunc"
	"github.com/rwsweeney/aac-str-tax-calculator/pkg/vrbofunc"
)

func main() {

	airbnbTaxData := airbnbfunc.CalculateAirbnb("airbnb_april.csv")
	vrboTaxData := vrbofunc.CalculateVRBO("vrbo_april.csv")

	fmt.Println("Nights rented on Airbnb: ", airbnbTaxData.Nights)
	fmt.Println("Nights rented on VRBO: ", vrboTaxData.Nights)

	fmt.Println("Box #1: ", airbnbTaxData.Nights+vrboTaxData.Nights)
	fmt.Println("Box #5: ", airbnbTaxData.GrossEarnings+vrboTaxData.GrossEarnings)
	fmt.Println("Box #7: ", airbnbTaxData.GrossEarnings+vrboTaxData.GrossEarnings)
	fmt.Println("Box #9: ", airbnbTaxData.GrossEarnings+vrboTaxData.GrossEarnings)
	fmt.Println("Box #10: ", airbnbTaxData.Aatax+vrboTaxData.Aatax)
	fmt.Println("Box #13: ", airbnbTaxData.Aatax+vrboTaxData.Aatax)
	fmt.Println("Box #14: ", airbnbTaxData.Aatax)
	fmt.Println("Box #15: ", vrboTaxData.Aatax)

}
