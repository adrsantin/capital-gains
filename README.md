# Capital Gains

This Go program reads JSON representing stock operation from stdin and calculates the amount of taxes for each operation. 

## JSON Input Format

The JSON input should be provided line by line, ending with a blank line. The input should consist of arrays of operations, where each line is an array of operations, following the format:

```json
[{"operation": "buy", "unit-cost": 10.00, "quantity": 10000},{"operation": "sell", "unit-cost": 20.00, "quantity": 5000}]
[{"operation": "buy", "unit-cost": 15.00, "quantity": 30000},{"operation": "sell", "unit-cost": 20.00, "quantity": 7000}]
```

### JSON Fields

- `operation` (string): The type of operation, either "buy" or "sell".
- `unit-cost` (float64): The cost per unit of the stock.
- `quantity` (int): The number of units involved in the operation.


## JSON Output format

The output consists of arrays of JSON objects describing the tax to be collected for each operation. The number of arrays and the number of objects will be the same as the number of arrays and operations from the input:

```json
[{"tax":0.00}, {"tax":10000.00}]
[{"tax":0.00}, {"tax":0.00}]
```

## How to compile

To compile the program, you need to have Go installed on your machine. Follow these steps:

- Copy all the source code to a directory.
- Open a terminal and navigate to the directory containing the source code.
- Run the following command to compile the program:

```bash
go build -o capital_gains main.go
```

This will create an executable file named `capital_gains`

## How to run

To run the program, use the following command:

```bash
./capital_gains
```

The program will wait for the JSON input from the stdin. You can provide the input directly:

```bash
echo '[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":20.00, "quantity": 5000}]' | ./stock_tax_calculator
```

Or use a file:

```bash
./stock_tax_calculator < input.txt
```

## Running unit tests

To run the unit tests for the program, use the following command:

```bash
go test ./...
```

This command will run all the tests in the current directory and its subdirectories. Test files are identifiable by the `_test.go` suffix.

## Examples

Examples of inputs are provided on the `example_inputs` directory.