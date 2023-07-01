# Anne Arrundel County Short Term Tax Calculator

The purpose of this repo is primarily for me to play around with Golang. I also started an Airbnb and have grown to loathe the rqeuirement to fill out this tax form every. single. month. Both AirBNB and VRBO automatically remit taxes to Anne Arrundel county, but we still must fill it out every month regardless because paperwork. Hooray.

## How to use:
1. Visit https://www.airbnb.com/users/transaction_history#gross-earnings and grab the earnings for the month. Save the PDF to aac-str-tax-calculator/financials folder
2. Get the payout summary from VRBO for the month: https://www.vrbo.com/p/opr/payoutsummary
3. Change main.go to point to the new csv
4. Open the CSVs and save as Text CSV format