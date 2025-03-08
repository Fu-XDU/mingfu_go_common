package flags

import "github.com/urfave/cli/v2"

var (
	InternalApiToken string
)

var (
	InternalApiTokenFlag = cli.StringFlag{
		Name:        "internal.token",
		Usage:       "Internal API token",
		EnvVars:     []string{"INTERNAL_API_TOKEN"},
		Required:    true,
		Destination: &InternalApiToken,
	}
)
