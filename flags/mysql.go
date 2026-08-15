package flags

import "github.com/urfave/cli/v2"

var (
	MysqlHost   string
	MysqlPort   int
	MysqlUser   string
	MysqlPasswd string
	MysqlDB     string
)

var (
	mysqlHostFlag = cli.StringFlag{
		Name:        "MysqlHost",
		Usage:       "Mysql host",
		Value:       "127.0.0.1",
		EnvVars:     []string{"MYSQL_HOST"},
		Destination: &MysqlHost,
	}

	mysqlPortFlag = cli.IntFlag{
		Name:        "MysqlPort",
		Usage:       "Mysql port",
		Value:       3306,
		EnvVars:     []string{"MYSQL_PORT"},
		Destination: &MysqlPort,
	}

	mysqlUserFlag = cli.StringFlag{
		Name:        "MysqlUser",
		Usage:       "Mysql user",
		Value:       "root",
		EnvVars:     []string{"MYSQL_USER"},
		Destination: &MysqlUser,
	}

	mysqlPasswdFlag = cli.StringFlag{
		Name:        "MysqlPasswd",
		Usage:       "Mysql password",
		Required:    true,
		EnvVars:     []string{"MYSQL_PASSWD"},
		Destination: &MysqlPasswd,
	}

	mysqlDBFlag = cli.StringFlag{
		Name:        "MysqlDB",
		Usage:       "Mysql database",
                Required:    true,
		EnvVars:     []string{"MYSQL_DB"},
		Destination: &MysqlDB,
	}
)

var MysqlFlags = []cli.Flag{
	&mysqlHostFlag,
	&mysqlPortFlag,
	&mysqlUserFlag,
	&mysqlPasswdFlag,
	&mysqlDBFlag,
}
