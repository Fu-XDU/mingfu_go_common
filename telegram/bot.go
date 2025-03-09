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

// SendTextMessage 发送普通文本消息
func (b *Bot) SendTextMessage(msg string, chatId string) (resp *SendMessageResponse, err error) {
	return b.sendMessage(msg, chatId, "")
}

// SendMarkdownMessage 发送 Markdown 格式的消息
func (b *Bot) SendMarkdownMessage(msg string, chatId string) (resp *SendMessageResponse, err error) {
	return b.sendMessage(msg, chatId, "Markdown")
}

// SendMarkdownV2Message 发送 MarkdownV2 格式的消息
func (b *Bot) SendMarkdownV2Message(msg string, chatId string) (resp *SendMessageResponse, err error) {
	return b.sendMessage(msg, chatId, "MarkdownV2")
}

// SendHtmlMessage 发送 HTML 格式的消息
func (b *Bot) SendHtmlMessage(msg string, chatId string) (resp *SendMessageResponse, err error) {
	return b.sendMessage(msg, chatId, "HTML")
}

// sendMessage 发送消息的底层实现
func (b *Bot) sendMessage(msg string, chatId string, parseMode string) (resp *SendMessageResponse, err error) {
	// 构造消息数据
	data := map[string]interface{}{
		"chat_id": chatId,
		"text":    msg,
	}
	if len(parseMode) != 0 {
		data["parse_mode"] = parseMode
	}

	// 发送请求
	body, err := network.Post(b.sendMessageUrl, data)
	if err != nil {
		return
	}

	// 解析响应
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
