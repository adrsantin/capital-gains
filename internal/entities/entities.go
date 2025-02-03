package entities

import "fmt"

type StockMarketOperation struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

type Tax struct {
	Tax float64 `json:"tax"`
}

type State struct {
	CurrentAverage float64
	CurrentStocks  int
	CurrentDebts   float64
}

func (t *Tax) ToString() string {
	return fmt.Sprintf(`{"tax":%.2f}`, t.Tax)
}
