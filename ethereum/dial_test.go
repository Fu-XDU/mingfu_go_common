package ethereum

import (
	"testing"
)

func TestDial(t *testing.T) {
	type args struct {
		rawurl string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Dial",
			args:    args{rawurl: "http://127.0.0.1:8545"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEthClient, err := Dial(tt.args.rawurl)
			if err == nil {
				defer gotEthClient.Close()
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Dial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDialWithAuth(t *testing.T) {
	type args struct {
		rawurl string
		auth   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Dial",
			args:    args{rawurl: "http://127.0.0.1:8545", auth: "auth"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEthClient, err := DialWithAuth(tt.args.rawurl, tt.args.auth)
			if err == nil {
				defer gotEthClient.Close()
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("DialWithAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
