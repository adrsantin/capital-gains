package services

import (
	"bytes"
	"errors"
	"testing"

	"github.com/adrsantin/taxapp/internal/entities"
	"github.com/stretchr/testify/assert"
)

func Test_ReadOperationsJSONFromReader(t *testing.T) {

	expectedRes := [][]entities.StockMarketOperation{
		{
			{
				Operation: "buy",
				UnitCost:  10,
				Quantity:  10000,
			},
			{
				Operation: "buy",
				UnitCost:  25,
				Quantity:  5000,
			},
			{
				Operation: "sell",
				UnitCost:  15,
				Quantity:  10000,
			},
		},
	}

	var stdin bytes.Buffer

	stdin.Write([]byte(`[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"buy", "unit-cost":25.00, "quantity": 5000},{"operation":"sell", "unit-cost":15.00, "quantity": 10000}]`))

	res, err := ReadOperationsJSONFromReader(&stdin)
	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)

	res, err = ReadOperationsJSONFromReader(&errorReader{data: "error"})
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func Test_ProcessOperations(t *testing.T) {
	type args struct {
		allOperations [][]entities.StockMarketOperation
	}
	tests := []struct {
		name string
		args args
		want [][]entities.Tax
	}{
		{
			name: "Should process operations wilhe multiple buy and sell operations",
			args: args{
				allOperations: [][]entities.StockMarketOperation{
					{
						{
							Operation: "buy",
							UnitCost:  10,
							Quantity:  10000,
						},
						{
							Operation: "sell",
							UnitCost:  2,
							Quantity:  5000,
						},
						{
							Operation: "sell",
							UnitCost:  20,
							Quantity:  2000,
						},
						{
							Operation: "sell",
							UnitCost:  20,
							Quantity:  2000,
						},
						{
							Operation: "sell",
							UnitCost:  25,
							Quantity:  1000,
						},
						{
							Operation: "buy",
							UnitCost:  20,
							Quantity:  10000,
						},
						{
							Operation: "sell",
							UnitCost:  15,
							Quantity:  5000,
						},
						{
							Operation: "sell",
							UnitCost:  30,
							Quantity:  4350,
						},
						{
							Operation: "sell",
							UnitCost:  30,
							Quantity:  650,
						},
					},
				},
			},
			want: [][]entities.Tax{
				{
					{
						Tax: 0,
					},
					{
						Tax: 0,
					},
					{
						Tax: 0,
					},
					{
						Tax: 0,
					},
					{
						Tax: 3000,
					},
					{
						Tax: 0,
					},
					{
						Tax: 0,
					},
					{
						Tax: 3700,
					},
					{
						Tax: 0,
					},
				},
			},
		},
		{
			name: "Should process 2 sets of operations",
			args: args{
				allOperations: [][]entities.StockMarketOperation{
					{
						{
							Operation: "buy",
							UnitCost:  10,
							Quantity:  100,
						},
						{
							Operation: "sell",
							UnitCost:  15,
							Quantity:  50,
						},
						{
							Operation: "sell",
							UnitCost:  15,
							Quantity:  50,
						},
					},
					{
						{
							Operation: "buy",
							UnitCost:  10,
							Quantity:  10000,
						},
						{
							Operation: "sell",
							UnitCost:  20,
							Quantity:  5000,
						},
						{
							Operation: "sell",
							UnitCost:  5,
							Quantity:  5000,
						},
					},
				},
			},
			want: [][]entities.Tax{
				{
					{
						Tax: 0,
					},
					{
						Tax: 0,
					},
					{
						Tax: 0,
					},
				},
				{
					{
						Tax: 0,
					},
					{
						Tax: 10000,
					},
					{
						Tax: 0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ProcessOperations(tt.args.allOperations)
			assert.Equal(t, tt.want, res)
		})
	}
}

func Test_ProcessOperation(t *testing.T) {
	type args struct {
		op     entities.StockMarketOperation
		status entities.State
	}
	type want struct {
		res    entities.Tax
		status entities.State
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Should process buy operation",
			args: args{
				op: entities.StockMarketOperation{
					Operation: "buy",
					UnitCost:  10,
					Quantity:  10,
				},
				status: entities.State{
					CurrentAverage: 0,
					CurrentStocks:  0,
					CurrentDebts:   0,
				},
			},
			want: want{
				res: entities.Tax{
					Tax: 0,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  10,
					CurrentDebts:   0,
				},
			},
		},
		{
			name: "Should process sell operation",
			args: args{
				op: entities.StockMarketOperation{
					Operation: "sell",
					UnitCost:  20,
					Quantity:  10,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  10,
					CurrentDebts:   0,
				},
			},
			want: want{
				res: entities.Tax{
					Tax: 0,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  0,
					CurrentDebts:   0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, status := processOperation(tt.args.op, tt.args.status)
			assert.Equal(t, tt.want.res, res)
			assert.Equal(t, tt.want.status, status)
		})
	}
}

func Test_ProcessBuyOperation(t *testing.T) {
	type args struct {
		status entities.State
		op     entities.StockMarketOperation
	}
	tests := []struct {
		name string
		args args
		want entities.State
	}{
		{
			name: "Should process buy operation",
			args: args{
				status: entities.State{
					CurrentAverage: 0,
					CurrentStocks:  0,
					CurrentDebts:   0,
				},
				op: entities.StockMarketOperation{
					Operation: "buy",
					UnitCost:  10,
					Quantity:  10,
				},
			},
			want: entities.State{
				CurrentAverage: 10,
				CurrentStocks:  10,
				CurrentDebts:   0,
			},
		},
		{
			name: "Should process buy operation recalculating average",
			args: args{
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  10,
					CurrentDebts:   0,
				},
				op: entities.StockMarketOperation{
					Operation: "buy",
					UnitCost:  20,
					Quantity:  10,
				},
			},
			want: entities.State{
				CurrentAverage: 15,
				CurrentStocks:  20,
				CurrentDebts:   0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := processBuyOperation(tt.args.status, tt.args.op)
			assert.Equal(t, tt.want, res)
		})
	}
}

func Test_ProcessSellOperation(t *testing.T) {
	type args struct {
		op     entities.StockMarketOperation
		res    entities.Tax
		status entities.State
	}
	type want struct {
		res    entities.Tax
		status entities.State
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Should process sell operation generating loss",
			args: args{
				op: entities.StockMarketOperation{
					Operation: "sell",
					UnitCost:  2,
					Quantity:  5000,
				},
				res: entities.Tax{
					Tax: 0,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  10000,
					CurrentDebts:   0,
				},
			},
			want: want{
				res: entities.Tax{
					Tax: 0,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  5000,
					CurrentDebts:   -40000,
				},
			},
		},
		{
			name: "Should process sell operation generating profit that does not pay losses",
			args: args{
				op: entities.StockMarketOperation{
					Operation: "sell",
					UnitCost:  20,
					Quantity:  2000,
				},
				res: entities.Tax{
					Tax: 0,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  5000,
					CurrentDebts:   -40000,
				},
			},
			want: want{
				res: entities.Tax{
					Tax: 0,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  3000,
					CurrentDebts:   -20000,
				},
			},
		},
		{
			name: "Should process sell operation generating profit that pays losses",
			args: args{
				op: entities.StockMarketOperation{
					Operation: "sell",
					UnitCost:  20,
					Quantity:  2000,
				},
				res: entities.Tax{
					Tax: 0,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  3000,
					CurrentDebts:   -20000,
				},
			},
			want: want{
				res: entities.Tax{
					Tax: 0,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  1000,
					CurrentDebts:   0,
				},
			},
		},
		{
			name: "Should process sell operation generating profit and generates tax",
			args: args{
				op: entities.StockMarketOperation{
					Operation: "sell",
					UnitCost:  25,
					Quantity:  1000,
				},
				res: entities.Tax{
					Tax: 0,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  1000,
					CurrentDebts:   0,
				},
			},
			want: want{
				res: entities.Tax{
					Tax: 3000,
				},
				status: entities.State{
					CurrentAverage: 10,
					CurrentStocks:  0,
					CurrentDebts:   0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, status := processSellOperation(tt.args.op, tt.args.res, tt.args.status)
			assert.Equal(t, tt.want.res, res)
			assert.Equal(t, tt.want.status, status)
		})
	}
}

type errorReader struct {
	data string
}

func (er *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("error")
}
