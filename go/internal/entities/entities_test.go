package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tax_ToString(t *testing.T) {
	type args struct {
		tax *Tax
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				tax: &Tax{
					Tax: 10,
				},
			},
			want: `{"tax":10.00}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.args.tax.ToString()
			assert.Equal(t, tt.want, res)
		})
	}
}
