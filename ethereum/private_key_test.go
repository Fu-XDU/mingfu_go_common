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

func TestPrivateKeyToPublicKey(t *testing.T) {
	type args struct {
		keyHex string
	}
	tests := []struct {
		name             string
		args             args
		wantPublicKeyHex string
		wantErr          bool
	}{
		{
			name: "TestPrivateKeyToPublicKey",
			args: args{
				keyHex: "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
			},
			wantPublicKeyHex: "0x048318535b54105d4a7aae60c08fc45f9687181b4fdfc625bd1a753fa7397fed753547f11ca8696646f2f3acb08e31016afac23e630c5d11f59f61fef57b0d2aa5",
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := HexToPrivateKey(tt.args.keyHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotPublicKey := PrivateKeyToPublicKey(key)
			if gotPublicKeyHex := PublicKeyToHex(&gotPublicKey, true); gotPublicKeyHex != tt.wantPublicKeyHex {
				t.Errorf("PrivateKeyToPublicKey() = %v, want %v", gotPublicKeyHex, tt.wantPublicKeyHex)
			}
		})
	}
}
