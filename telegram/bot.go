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

// BotMessageType 定义消息类型枚举
type BotMessageType string

const (
	Text     BotMessageType = "Text"
	Markdown BotMessageType = "Markdown"
	HTML     BotMessageType = "HTML"
)

// SendMessageOptions 发送消息的可选参数
type SendMessageOptions struct {
	DisablePreview           bool        // 关闭链接预览
	DisableNotification      bool        // 静默消息，不触发通知
	ProtectContent           bool        // 保护消息内容，防止转发
	ReplyToMessageID         int         // 回复的消息 ID
	AllowSendingWithoutReply bool        // 允许在没有有效回复时发送
	ReplyMarkup              interface{} // 自定义键盘
}

// SendTextMessage 发送普通文本消息
func (b *Bot) SendTextMessage(msg string, chatId string, opts ...SendMessageOptions) (resp *SendMessageResponse, err error) {
	return b.SendMessage(msg, chatId, Text, opts...)
}

// SendMarkdownMessage 发送 Markdown 格式的消息
func (b *Bot) SendMarkdownMessage(msg string, chatId string, opts ...SendMessageOptions) (resp *SendMessageResponse, err error) {
	return b.SendMessage(msg, chatId, Markdown, opts...)
}

// SendHtmlMessage 发送 HTML 格式的消息
func (b *Bot) SendHtmlMessage(msg string, chatId string, opts ...SendMessageOptions) (resp *SendMessageResponse, err error) {
	return b.SendMessage(msg, chatId, HTML, opts...)
}

// SendMessage 发送消息的底层实现
func (b *Bot) SendMessage(msg string, chatId string, parseMode BotMessageType, opts ...SendMessageOptions) (resp *SendMessageResponse, err error) {
	// 构造消息数据
	data := map[string]interface{}{
		"chat_id": chatId,
		"text":    msg,
	}

	if parseMode != Text {
		data["parse_mode"] = parseMode
	}

	// 解析可选参数
	var options SendMessageOptions
	if len(opts) > 0 {
		options = opts[0]
		data["disable_web_page_preview"] = options.DisablePreview
		data["disable_notification"] = options.DisableNotification
		data["protect_content"] = options.ProtectContent
		data["reply_to_message_id"] = options.ReplyToMessageID
		data["allow_sending_without_reply"] = options.AllowSendingWithoutReply
		data["reply_markup"] = options.ReplyMarkup
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
