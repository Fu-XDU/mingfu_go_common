package flags

import "github.com/urfave/cli/v2"

var (
	KafkaHost   string
	KafkaPort   int
	KafkaUser   string
	KafkaPasswd string
)

var (
	kafkaHostFlag = cli.StringFlag{
		Name:        "KafkaHost",
		Usage:       "Kafka host",
		Value:       "kafka.common.svc.cluster.local",
		EnvVars:     []string{"KAFKA_HOST"},
		Destination: &KafkaHost,
	}

	kafkaPortFlag = cli.IntFlag{
		Name:        "KafkaPort",
		Usage:       "Kafka port",
		Value:       9092,
		EnvVars:     []string{"KAFKA_PORT"},
		Destination: &KafkaPort,
	}

	kafkaUserFlag = cli.StringFlag{
		Name:        "KafkaUser",
		Usage:       "Kafka user",
		Value:       "root",
		EnvVars:     []string{"KAFKA_USER"},
		Destination: &KafkaUser,
	}

	kafkaPasswdFlag = cli.StringFlag{
		Name:        "KafkaPasswd",
		Usage:       "Kafka password",
		Value:       "d26363056c64e99ae255a949b83c30e8",
		EnvVars:     []string{"KAFKA_PASSWD"},
		Destination: &KafkaPasswd,
	}
)

var KafkaFlags = []cli.Flag{
	&kafkaHostFlag,
	&kafkaPortFlag,
	&kafkaUserFlag,
	&kafkaPasswdFlag,
}
