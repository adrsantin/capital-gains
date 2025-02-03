package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/adrsantin/taxapp/internal/entities"
	"github.com/adrsantin/taxapp/internal/helpers"
)

func ReadOperationsJSONFromReader(stdin io.Reader) ([][]entities.StockMarketOperation, error) {
	scanner := bufio.NewScanner(stdin)

	var allOperations [][]entities.StockMarketOperation

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var op []entities.StockMarketOperation
		err := json.Unmarshal([]byte(line), &op)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
		}
		allOperations = append(allOperations, op)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
		return nil, err
	}
	return allOperations, nil
}

func ProcessOperations(allOperations [][]entities.StockMarketOperation) [][]entities.Tax {
	allTaxes := make([][]entities.Tax, 0)
	for _, ops := range allOperations {

		status := entities.State{
			CurrentAverage: 0,
			CurrentStocks:  0,
			CurrentDebts:   0,
		}

		taxes := make([]entities.Tax, 0)
		for _, op := range ops {
			var tax entities.Tax
			tax, status = processOperation(op, status)
			taxes = append(taxes, tax)
		}
		allTaxes = append(allTaxes, taxes)
	}
	return allTaxes
}

func processOperation(op entities.StockMarketOperation, status entities.State) (entities.Tax, entities.State) {
	res := entities.Tax{
		Tax: 0,
	}
	switch op.Operation {
	case "buy":
		{
			status = processBuyOperation(status, op)
		}
	case "sell":
		{
			res, status = processSellOperation(op, res, status)
		}
	}
	return res, status
}

func processBuyOperation(status entities.State, op entities.StockMarketOperation) entities.State {
	status.CurrentAverage = helpers.CalculateAverage(status.CurrentAverage, float64(status.CurrentStocks), float64(op.Quantity), op.UnitCost)
	status.CurrentStocks += op.Quantity
	return status
}

func processSellOperation(op entities.StockMarketOperation, res entities.Tax, status entities.State) (entities.Tax, entities.State) {
	amount := float64(op.Quantity) * (op.UnitCost - status.CurrentAverage)
	fullOperationAmount := float64(op.Quantity) * op.UnitCost
	status.CurrentStocks -= op.Quantity
	if amount > 0 { // profit
		remainder := amount + status.CurrentDebts
		if remainder > 0 {
			status.CurrentDebts = 0
			amount = remainder
			if fullOperationAmount > 20000 {
				res.Tax = amount * 0.2
			}
		} else {
			status.CurrentDebts = remainder
			amount = 0
		}
	} else { // loss
		status.CurrentDebts += amount
	}
	return res, status
}
