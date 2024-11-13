package flags

import (
	"fmt"
	"github.com/Fu-XDU/mingfu_go_common/file"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
	"strconv"
)

var (
	ServerPort     string
	SslCertPath    string
	SslKeyPath     string
	TrustedProxies []string
)

var (
	portFlag = cli.StringFlag{
		Name:        "port",
		Aliases:     []string{"p"},
		Usage:       "Server port",
		Value:       "1423",
		EnvVars:     []string{"SERVER_PORT"},
		Destination: &ServerPort,
		Action: func(ctx *cli.Context, portStr string) (err error) {
			port, err := strconv.Atoi(portStr)
			if err != nil || port <= 0 || port >= 1<<16 {
				err = fmt.Errorf("flag port value %v out of range[0-65535].", port)
			}
			return
		},
	}

	sslCertFlag = cli.StringFlag{
		Name:        "ssl_cert, c",
		Aliases:     []string{"c"},
		Usage:       "SSL cert path",
		Value:       "",
		EnvVars:     []string{"SSL_CERT"},
		Destination: &SslCertPath,
		Action: func(ctx *cli.Context, path string) (err error) {
			if len(path) != 0 {
				exist, _ := file.FileExists(path)
				if !exist {
					err = fmt.Errorf("SSL cert %v is not exist.", path)
				}
			}
			return
		},
	}

	sslKeyFlag = cli.StringFlag{
		Name:        "ssl_key",
		Aliases:     []string{"k"},
		Usage:       "SSL key path",
		Value:       "",
		EnvVars:     []string{"SSL_KEY"},
		Destination: &SslKeyPath,
		Action: func(ctx *cli.Context, path string) (err error) {
			if len(path) != 0 {
				exist, _ := file.FileExists(path)
				if !exist {
					err = fmt.Errorf("SSL key %v is not exist.", path)
				}
			}
			return
		},
	}

	trustedProxiesFlag = cli.StringSliceFlag{
		Name:    "trusted_proxies",
		Usage:   "Trusted proxies, example: --trusted_proxies 127.0.0.1,172.0.0.1",
		EnvVars: []string{"TRUSTED_PROXIES"},
		Value:   cli.NewStringSlice(),
		Action: func(ctx *cli.Context, trustedProxies []string) (err error) {
			if len(trustedProxies) != 0 {
				log.Infof("Trust proxies: %v.", trustedProxies)
				TrustedProxies = trustedProxies
			}
			return
		},
	}
)

var GinFlags = []cli.Flag{
	&portFlag,
	&sslCertFlag,
	&sslKeyFlag,
	&trustedProxiesFlag,
}
