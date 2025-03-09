package wecom

import (
	"fmt"
	"github.com/Fu-XDU/mingfu_go_common/network"
)

type Bot struct {
	key            string
	sendMessageUrl string
}

func NewBot(key string) *Bot {
	return &Bot{
		key:            key,
		sendMessageUrl: fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", key),
	}
}

// SendMessage 发送消息的底层实现
func (b *Bot) SendMessage(msg string) (err error) {
	// 请求数据
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": msg,
		},
	}

	// 发送请求
	_, err = network.Post(b.sendMessageUrl, payload)
	if err != nil {
		return
	}

	return
}
