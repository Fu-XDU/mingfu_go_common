package ethereum

import (
	"reflect"
	"testing"
)

func TestNewRandomPrivateKey(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "TestNewRandomPrivateKey",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRandomPrivateKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRandomPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHexToPrivateKey(t *testing.T) {
	type args struct {
		keyHex string
	}
	tests := []struct {
		name        string
		args        args
		wantAddress string
		wantErr     bool
	}{
		{
			name:        "TestHexToPrivateKey",
			args:        args{keyHex: "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"},
			wantAddress: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToPrivateKey(tt.args.keyHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			address := PrivateKeyToAddress(got)

			if address.Hex() != tt.wantAddress {
				t.Errorf("HexToPrivateKey() got = %v, want %v", address, tt.wantAddress)
			}
		})
	}
}

func TestPrivateKeyToAddress(t *testing.T) {
	tests := []struct {
		name        string
		keyHex      string
		wantAddress string
		wantErr     bool
	}{
		{
			name:        "TestPrivateKeyToAddress",
			keyHex:      "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
			wantAddress: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := HexToPrivateKey(tt.keyHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got := PrivateKeyToAddress(key); !reflect.DeepEqual(got.String(), tt.wantAddress) {
				t.Errorf("PrivateKeyToAddress() = %v, want %v", got.String(), tt.wantAddress)
			}
		})
	}
}

func TestPrivateKeyToHex(t *testing.T) {
	tests := []struct {
		name    string
		keyHex  string
		wantErr bool
	}{
		{
			name:    "TestPrivateKeyToHex",
			keyHex:  "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := HexToPrivateKey(tt.keyHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotHex := PrivateKeyToHex(key, true); gotHex != tt.keyHex {
				t.Errorf("PrivateKeyToHex() = %v, want %v", gotHex, tt.keyHex)
			}
		})
	}
}
