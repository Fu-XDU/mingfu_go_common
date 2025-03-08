package flags

import (
	"fmt"
	"github.com/Fu-XDU/mingfu_go_common/auth"
	"github.com/urfave/cli/v2"
)

var JwtSalt string

var (
	jwtSaltFlag = cli.StringFlag{
		Name:        "jwt_salt",
		Usage:       "JWT salt, a random string which at least 32 bytes",
		EnvVars:     []string{"JWT_SALT"},
		Required:    true,
		Destination: &JwtSalt,
		Action: func(ctx *cli.Context, salt string) (err error) {
			if len(salt) < 32 {
				err = fmt.Errorf("length of salt should be at least 32 bytes, but got %v", len(salt))
			}
			auth.SetJwtSalt(salt)
			return
		},
	}
)

var AuthFlags = []cli.Flag{
	&jwtSaltFlag,
}
