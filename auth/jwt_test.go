package auth

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"os"
	"testing"
	"time"
)

var jwtToken *JwtToken

func TestMain(m *testing.M) {
	SetJwtSalt("0420bef29f65c4f1fb833231a9b1287ffb7e3376dae4d27796b22c6f4a1a14f8")
	jwtToken = NewJwt("1234567890", time.Hour)
	os.Exit(m.Run())
}

func TestJwtClaims_Valid(t *testing.T) {
	type fields JwtClaims
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid case, wrong uuid and iat",
			fields: fields{
				Uuid: uuid.New().String(),
				Iat:  jwtToken.Claims().Iat.Add(-24 * time.Hour),
				Nbf:  jwtToken.Claims().Nbf,
				Exp:  jwtToken.Claims().Exp,
				Rand: random.New().String(31),
			},
			wantErr: false,
		},
		{
			name: "Invalid case, wrong nbf",
			fields: fields{
				Uuid: jwtToken.Claims().Uuid,
				Iat:  jwtToken.Claims().Iat,
				Nbf:  jwtToken.Claims().Nbf.Add(10 * time.Minute),
				Exp:  jwtToken.Claims().Exp,
				Rand: jwtToken.Claims().Rand,
			},
			wantErr: true,
		},
		{
			name: "Invalid case, wrong exp",
			fields: fields{
				Uuid: jwtToken.Claims().Uuid,
				Iat:  jwtToken.Claims().Iat,
				Nbf:  jwtToken.Claims().Nbf,
				Exp:  jwtToken.Claims().Exp.Add(-24 * time.Hour),
				Rand: jwtToken.Claims().Rand,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &JwtClaims{
				Uuid: tt.fields.Uuid,
				Iat:  tt.fields.Iat,
				Nbf:  tt.fields.Nbf,
				Exp:  tt.fields.Exp,
				Rand: tt.fields.Rand,
			}
			if err := c.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewJwt(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{uuid: "1234567890"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := NewJwt(tt.args.uuid, time.Hour)
			tokenParsed, err := ParseJwt(token.String())
			if err != nil {
				panic(err)
			}
			if !tokenParsed.Valid() {
				panic("invalid token")
			}
			if tokenParsed.String() != token.String() {
				panic("token string mismatch")
			}
			if tokenParsed.Claims().Uuid != token.Claims().Uuid {
				panic("token claims uuid mismatch")
			}
		})
	}
}

func TestParseJwt(t *testing.T) {
	type args struct {
		tokenStr string
	}
	tests := []struct {
		name            string
		args            args
		wantJwtValid    bool
		wantJwtTokenStr string
		wantErr         bool
	}{
		{
			name: "Valid case",
			args: args{
				tokenStr: jwtToken.String(),
			},
			wantJwtTokenStr: jwtToken.String(),
			wantErr:         false,
		},
		{
			name: "Wrong case",
			args: args{
				tokenStr: random.New().String(uint8(len(jwtToken.String()))),
			},
			wantJwtValid:    false,
			wantJwtTokenStr: jwtToken.String(),
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotJwtToken, err := ParseJwt(tt.args.tokenStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJwt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if gotJwtToken.String() != tt.wantJwtTokenStr {
				t.Errorf("ParseJwt() gotJwtToken = %v, want %v", gotJwtToken, tt.wantJwtTokenStr)
			}
		})
	}
}
