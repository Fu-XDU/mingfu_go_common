package flags

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strconv"
)

var (
	ServerPort string
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
)

var GinFlags = []cli.Flag{
	&portFlag,
}
