package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/Fu-XDU/mingfu_go_common/network"
)

type Bot struct {
	token          string
	sendMessageUrl string
}

type SendMessageResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID        int64  `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

func (b *Bot) SendMessage(msg string, chatId string) (resp *SendMessageResponse, err error) {
	// 构造消息数据
	data := map[string]interface{}{
		"chat_id": chatId,
		"text":    msg,
	}

	body, err := network.Post(b.sendMessageUrl, data)
	if err != nil {
		return
	}

	resp = new(SendMessageResponse)
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}

	return
}

func NewBot(token string) *Bot {
	return &Bot{
		token:          token,
		sendMessageUrl: fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token),
	}
}
