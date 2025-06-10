package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestTransferEthTo(t *testing.T) {
	type args struct {
		rpcRawUrl string
		senderHex string
		receipt   common.Address
		valueWei  *big.Int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestTransferEthTo",
			args: args{
				rpcRawUrl: "http://127.0.0.1:8545",
				senderHex: "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
				receipt:   common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"),
				valueWei:  big.NewInt(0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := Dial(tt.args.rpcRawUrl)
			if err != nil {
				t.Errorf("Dial() error = %v", err)
				return
			}

			sender, err := HexToPrivateKey(tt.args.senderHex)
			if err != nil {
				t.Errorf("HexToPrivateKey() error = %v", err)
				return
			}

			_, _, err = TransferEthTo(client, sender, &tt.args.receipt, tt.args.valueWei)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransferEthTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
