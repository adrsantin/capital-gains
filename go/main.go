package main

import (
	"fmt"
	"os"

	"github.com/adrsantin/taxapp/internal/helpers"
	"github.com/adrsantin/taxapp/internal/services"
)

func main() {
	operations, err := services.ReadOperationsJSONFromReader(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	taxes := services.ProcessOperations(operations)
	fmt.Print(helpers.TaxesToPrint(taxes))
}
