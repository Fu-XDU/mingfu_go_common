package flags

import "github.com/urfave/cli/v2"

var (
	BotToken  string
	BotChatID string
)

var (
	tgBotTokenFlag = cli.StringFlag{
		Name:        "tg.bot_token",
		Usage:       "Telegram bot token",
		Required:    false,
		EnvVars:     []string{"TG_BOT_TOKEN"},
		Destination: &BotToken,
	}

	tgBotChatIDFlag = cli.StringFlag{
		Name:        "tg.bot_chat_id",
		Usage:       "Telegram bot chat ID",
		Required:    false,
		EnvVars:     []string{"TG_BOT_CHAT_ID"},
		Destination: &BotChatID,
	}
)

var TelegramFlags = []cli.Flag{
	&tgBotTokenFlag,
	&tgBotChatIDFlag,
}
