package etherscan

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"github.com/nanmu42/etherscan-api"
	"testing"
)

func TestAllNormalTxByAddress(t *testing.T) {
	type args struct {
		client   *etherscan.Client
		endBlock *int
		address  common.Address
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "0x93072e593a233Cef979F37F898fBECfb2a09218e",
			args: args{
				client:   etherscan.New(etherscan.Mainnet, ""),
				endBlock: nil,
				address:  common.HexToAddress("0x93072e593a233Cef979F37F898fBECfb2a09218e"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTxs, err := AllNormalTxByAddress(tt.args.client, tt.args.endBlock, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllNormalTxByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			log.Infof("got all normal tx by address %s, count = %d", tt.args.address, len(gotTxs))
		})
	}
}

func TestAllNormalTxFromAddress(t *testing.T) {
	type args struct {
		client   *etherscan.Client
		endBlock *int
		address  common.Address
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "0x93072e593a233Cef979F37F898fBECfb2a09218e",
			args: args{
				client:   etherscan.New(etherscan.Mainnet, ""),
				endBlock: nil,
				address:  common.HexToAddress("0x93072e593a233Cef979F37F898fBECfb2a09218e"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTxs, err := AllNormalTxFromAddress(tt.args.client, tt.args.endBlock, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllNormalTxFromAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			log.Infof("got all normal tx from address %s, count = %d", tt.args.address, len(gotTxs))
		})
	}
}

func TestAllNormalTxToAddress(t *testing.T) {
	type args struct {
		client   *etherscan.Client
		endBlock *int
		address  common.Address
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "0x93072e593a233Cef979F37F898fBECfb2a09218e",
			args: args{
				client:   etherscan.New(etherscan.Mainnet, ""),
				endBlock: nil,
				address:  common.HexToAddress("0x93072e593a233Cef979F37F898fBECfb2a09218e"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTxs, err := AllNormalTxToAddress(tt.args.client, tt.args.endBlock, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllNormalTxToAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			log.Infof("got all normal tx to address %s, count = %d", tt.args.address, len(gotTxs))
		})
	}
}
