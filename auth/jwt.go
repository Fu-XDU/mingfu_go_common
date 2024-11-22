package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"github.com/labstack/gommon/random"
	"strings"
	"time"
)

const BearerPrefix = "Bearer "

var invalidJwtClaims = errors.New("invalid jwt claims")
var jwtSalt []byte

// SetJwtSalt Can only be set once
func SetJwtSalt(salt string) bool {
	if len(jwtSalt) != 0 {
		return false
	}

	log.Info("SetJwtSalt Success")

	jwtSalt = []byte(salt)
	return true
}

type JwtClaims struct {
	Uuid string    `json:"uuid"`
	Nbf  time.Time `json:"nbf"`
	Iat  time.Time `json:"iat"` // For record only, not for verification
	Exp  time.Time `json:"exp"`
	Rand string    `json:"rand"`
}

func (c *JwtClaims) Valid() error {
	now := time.Now()
	if c.Nbf.Before(now) && c.Exp.After(now) {
		return nil
	}
	return invalidJwtClaims
}

type JwtToken struct {
	Token *jwt.Token
}

func NewJwt(uuid string, exp time.Duration) *JwtToken {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtClaims{
		Uuid: uuid,
		Nbf:  now,
		Iat:  now,
		Exp:  now.Add(exp),
		Rand: random.New().String(31),
	})

	return &JwtToken{
		Token: token,
	}
}

func ParseJwt(tokenStr string) (jwtToken *JwtToken, err error) {
	if strings.HasPrefix(tokenStr, BearerPrefix) {
		tokenStr = tokenStr[len(BearerPrefix):]
	}

	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSalt, nil
	})

	if err != nil {
		return
	}

	jwtToken = &JwtToken{Token: token}
	return
}

func (j *JwtToken) Claims() *JwtClaims {
	return j.Token.Claims.(*JwtClaims)
}

func (j *JwtToken) String() string {
	tokenStr, _ := j.Token.SignedString(jwtSalt)
	return tokenStr
}

func (j *JwtToken) StringWithBearerPrefix() string {
	return BearerPrefix + j.String()
}

func (j *JwtToken) Valid() bool {
	err := j.Claims().Valid()
	return err == nil
}

func VerifyJwt(tokenStr string) (uuid string, valid bool) {
	jwtToken, err := ParseJwt(tokenStr)
	if err != nil || !jwtToken.Valid() {
		return
	}

	uuid = jwtToken.Claims().Uuid
	valid = true
	return
}
