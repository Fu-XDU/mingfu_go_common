package flags

import "github.com/urfave/cli/v2"

var (
	BotKey string
)

var (
	wecomBotKeyFlag = cli.StringFlag{
		Name:        "wecom.bot_key",
		Usage:       "WeCom bot key",
		Required:    false,
		EnvVars:     []string{"WECOM_BOT_KEY"},
		Destination: &BotKey,
	}
)

var WecomFlags = []cli.Flag{
	&wecomBotKeyFlag,
}
