package influxdb

import (
	"github.com/Fu-XDU/mingfu_go_common/flags"
	"github.com/InfluxCommunity/influxdb3-go/v2/influxdb3"
)

func NewClientConfigFromFlags() *influxdb3.ClientConfig {
	return &influxdb3.ClientConfig{
		Host:       flags.InfluxdbHost,
		Token:      flags.InfluxdbToken,
		AuthScheme: flags.InfluxdbAuthScheme,
		Database:   flags.InfluxdbDB,
	}
}

func Connect(config *influxdb3.ClientConfig) (client *influxdb3.Client, err error) {
	return influxdb3.New(*config)
}
