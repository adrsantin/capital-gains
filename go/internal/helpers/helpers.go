package helpers

import (
	"fmt"

	"github.com/adrsantin/taxapp/internal/entities"
)

func CalculateAverage(
	currentAverage float64,
	currentStocks float64,
	stocksBought float64,
	price float64,
) float64 {
	return ((currentStocks * currentAverage) + (stocksBought * price)) / (currentStocks + stocksBought)
}

func TaxesToPrint(allTaxes [][]entities.Tax) string {
	res := ""
	for _, taxes := range allTaxes {
		line := "["
		for _, tax := range taxes {
			line += fmt.Sprint(tax.ToString())
			line += ","
		}
		line = line[:len(line)-1]
		line += "]\n"
		res += line
	}
	return res
}
