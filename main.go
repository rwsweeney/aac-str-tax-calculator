// TO DO
// Create a function to show the work behind the calculations.
// Airbnb: Should run a test calculation of (Gross Earnings * .06) + (Gross Earnings * .07) = occupancy tax value.

package main

import (
	"fmt"

	"github.com/rwsweeney/aac-str-tax-calculator/pkg/airbnbfunc"
	"github.com/rwsweeney/aac-str-tax-calculator/pkg/vrbofunc"
)

func main() {

	airbnbTaxData := airbnbfunc.CalculateAirbnb()
	vrboTaxData := vrbofunc.CalculateVRBO()

	fmt.Println("Box #1: ", airbnbTaxData.Nights+vrboTaxData.Nights)
	fmt.Println("Box #5: ", airbnbTaxData.GrossEarnings+vrboTaxData.GrossEarnings)
	fmt.Println("Box #7: ", airbnbTaxData.GrossEarnings+vrboTaxData.GrossEarnings)
	fmt.Println("Box #9: ", airbnbTaxData.GrossEarnings+vrboTaxData.GrossEarnings)
	fmt.Println("Box #10: ", airbnbTaxData.Aatax+vrboTaxData.Aatax)
	fmt.Println("Box #13: ", airbnbTaxData.Aatax+vrboTaxData.Aatax)
	fmt.Println("Box #14: ", airbnbTaxData.Aatax)
	fmt.Println("Box #15: ", vrboTaxData.Aatax)

}
