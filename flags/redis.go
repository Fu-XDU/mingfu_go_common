package flags

import "github.com/urfave/cli/v2"

var (
	RedisHost     string
	RedisPort     int
	RedisDb       int
	RedisUser     string
	RedisPasswd   string
	RedisPoolSize int
)

var (
	redisHostFlag = cli.StringFlag{
		Name:        "RedisHost",
		Usage:       "RedisHost",
		Value:       "127.0.0.1",
		EnvVars:     []string{"REDIS_HOST"},
		Destination: &RedisHost,
	}

	redisPortFlag = cli.IntFlag{
		Name:        "RedisPort",
		Usage:       "RedisPort",
		Value:       6379,
		EnvVars:     []string{"REDIS_PORT"},
		Destination: &RedisPort,
	}

	redisDbFlag = cli.IntFlag{
		Name:        "RedisDb",
		Usage:       "RedisDb",
		Value:       0,
		EnvVars:     []string{"REDIS_DB"},
		Destination: &RedisDb,
	}

	redisUserFlag = cli.StringFlag{
		Name:        "RedisUser",
		Usage:       "Database username for Redis",
		Value:       "",
		EnvVars:     []string{"REDIS_USER"},
		Destination: &RedisUser,
	}

	redisPasswdFlag = cli.StringFlag{
		Name:        "RedisPasswd",
		Usage:       "Database password for Redis",
                Required:    true,
		EnvVars:     []string{"REDIS_PASSWD"},
		Destination: &RedisPasswd,
	}

	redisPoolSizeFlag = cli.IntFlag{
		Name:        "RedisPoolSize",
		Usage:       "RedisPoolSize",
		Value:       2,
		EnvVars:     []string{"REDIS_POOL_SIZE"},
		Destination: &RedisPoolSize,
	}
)

var RedisFlags = []cli.Flag{
	&redisHostFlag,
	&redisPortFlag,
	&redisDbFlag,
	&redisUserFlag,
	&redisPasswdFlag,
	&redisPoolSizeFlag,
}
