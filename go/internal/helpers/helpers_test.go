package helpers

import (
	"testing"

	"github.com/adrsantin/taxapp/internal/entities"
	"github.com/stretchr/testify/assert"
)

func Test_CalculateAverage(t *testing.T) {
	type args struct {
		currentAverage float64
		currentStocks  float64
		stocksBought   float64
		price          float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				currentAverage: 10,
				currentStocks:  10,
				stocksBought:   10,
				price:          10,
			},
			want: 10,
		},
		{
			name: "Test 2",
			args: args{
				currentAverage: 10,
				currentStocks:  10,
				stocksBought:   10,
				price:          20,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := CalculateAverage(tt.args.currentAverage, tt.args.currentStocks, tt.args.stocksBought, tt.args.price)
			assert.Equal(t, tt.want, res)
		})
	}
}

func Test_TaxesToPrint(t *testing.T) {
	type args struct {
		allTaxes [][]entities.Tax
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				allTaxes: [][]entities.Tax{
					{
						{
							Tax: 10,
						},
						{
							Tax: 20,
						},
					},
					{
						{
							Tax: 30,
						},
						{
							Tax: 40,
						},
					},
				},
			},
			want: `[{"tax":10.00},{"tax":20.00}]
[{"tax":30.00},{"tax":40.00}]
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := TaxesToPrint(tt.args.allTaxes)
			assert.Equal(t, tt.want, res)
		})
	}
}
