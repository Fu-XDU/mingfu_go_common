package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"reflect"
	"testing"
)

func TestHexToPublicKey(t *testing.T) {
	type args struct {
		keyHex string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestHexToPublicKey",
			args: args{
				keyHex: "0x048318535b54105d4a7aae60c08fc45f9687181b4fdfc625bd1a753fa7397fed753547f11ca8696646f2f3acb08e31016afac23e630c5d11f59f61fef57b0d2aa5",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToPublicKey(tt.args.keyHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHex := PublicKeyToHex(got, true); gotHex != tt.args.keyHex {
				t.Errorf("HexToPublicKey() got = %v, want %v", gotHex, tt.args.keyHex)
			}
		})
	}
}

func TestPublicKeyToAddress(t *testing.T) {
	type args struct {
		keyHex string
	}
	tests := []struct {
		name    string
		args    args
		want    common.Address
		wantErr bool
	}{
		{
			name: "TestPublicKeyToAddress",
			args: args{
				keyHex: "0x048318535b54105d4a7aae60c08fc45f9687181b4fdfc625bd1a753fa7397fed753547f11ca8696646f2f3acb08e31016afac23e630c5d11f59f61fef57b0d2aa5",
			},
			want:    common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			publicKey, err := HexToPublicKey(tt.args.keyHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got := PublicKeyToAddress(*publicKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublicKeyToAddress() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}

func TestPublicKeyToHex(t *testing.T) {
	type args struct {
		privateKeyHex string
		withPrefix    bool
	}
	tests := []struct {
		name    string
		args    args
		wantH   string
		wantErr bool
	}{
		{
			name: "TestPublicKeyToHex",
			args: args{
				privateKeyHex: "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
				withPrefix:    true,
			},
			wantH:   "0x048318535b54105d4a7aae60c08fc45f9687181b4fdfc625bd1a753fa7397fed753547f11ca8696646f2f3acb08e31016afac23e630c5d11f59f61fef57b0d2aa5",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := HexToPrivateKey(tt.args.privateKeyHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotH := PublicKeyToHex(&key.PublicKey, tt.args.withPrefix); gotH != tt.wantH {
				t.Errorf("PublicKeyToHex() = %v, want %v", gotH, tt.wantH)
			}
		})
	}
}
