package flags

import "github.com/urfave/cli/v2"

var (
	InfluxdbHost       string
	InfluxdbToken      string
	InfluxdbAuthScheme string
	InfluxdbDB         string
)

var (
	influxdbHostFlag = cli.StringFlag{
		Name:        "InfluxdbHost",
		Usage:       "Influxdb host",
		Value:       "127.0.0.1",
		EnvVars:     []string{"INFLUXDB_HOST"},
		Destination: &InfluxdbHost,
	}

	influxdbTokenFlag = cli.StringFlag{
		Name:        "InfluxdbToken",
		Usage:       "Influxdb token",
		Required:    true,
		EnvVars:     []string{"INFLUXDB_USER"},
		Destination: &InfluxdbToken,
	}

	influxdbAuthSchemeFlag = cli.StringFlag{
		Name:        "InfluxdbAuthScheme",
		Usage:       "Influxdb auth scheme",
		Value:       "Bearer",
		EnvVars:     []string{"INFLUXDB_AUTH_SCHEME"},
		Destination: &InfluxdbAuthScheme,
	}

	influxdbDBFlag = cli.StringFlag{
		Name:        "InfluxdbDB",
		Usage:       "Influxdb database",
		EnvVars:     []string{"INFLUXDB_DB"},
		Destination: &InfluxdbDB,
	}
)

var InfluxdbFlags = []cli.Flag{
	&influxdbHostFlag,
	&influxdbTokenFlag,
	&influxdbAuthSchemeFlag,
	&influxdbDBFlag,
}
