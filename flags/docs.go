package flags

import (
	"github.com/urfave/cli/v2"
)

var (
	ShowDocs bool
)

var (
	DocsFlag = cli.BoolFlag{
		Name:        "docs",
		Usage:       "Show docs",
		EnvVars:     []string{"SHOW_DOCS"},
		Value:       false,
		Destination: &ShowDocs,
	}
)
