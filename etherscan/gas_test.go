package etherscan

import (
	"github.com/nanmu42/etherscan-api"
	"math/big"
	"reflect"
	"testing"
)

func TestSumGasFees(t *testing.T) {
	type args struct {
		txs []etherscan.NormalTx
	}
	tests := []struct {
		name    string
		args    args
		wantSum *big.Int
	}{
		{
			name: "SumGasFees",
			args: args{
				txs: []etherscan.NormalTx{
					{
						GasUsed:  1,
						GasPrice: (*etherscan.BigInt)(big.NewInt(1)),
					},
					{
						GasUsed:  2,
						GasPrice: (*etherscan.BigInt)(big.NewInt(2)),
					},
				},
			},
			wantSum: big.NewInt(int64(5)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := SumGasFees(tt.args.txs); !reflect.DeepEqual(gotSum, tt.wantSum) {
				t.Errorf("SumGasFees() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
