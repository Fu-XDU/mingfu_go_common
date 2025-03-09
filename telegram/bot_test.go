package telegram

import (
	"testing"
	"time"
)

func TestBot_SendMessage(t *testing.T) {
	token := ""
	chatId := ""
	type fields struct {
		token string
	}
	type args struct {
		msg     string
		chatId  string
		msgType BotMessageType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Text Message",
			fields: fields{
				token: token,
			},
			args: args{
				msg:     time.Now().Format(time.DateTime),
				chatId:  chatId,
				msgType: Text,
			},
			wantErr: false,
		},
		{
			name: "Markdown Message",
			fields: fields{
				token: token,
			},
			args: args{
				msg:     time.Now().Format(time.DateTime),
				chatId:  chatId,
				msgType: Markdown,
			},
			wantErr: false,
		},
		{
			name: "HTML Message",
			fields: fields{
				token: token,
			},
			args: args{
				msg:     time.Now().Format(time.DateTime),
				chatId:  chatId,
				msgType: HTML,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.fields.token) == 0 {
				t.Fatal("tt.fields.token can't be empty")
			}

			if len(tt.args.chatId) == 0 {
				t.Fatal("tt.args.chatId can't be empty")
			}

			b := NewBot(tt.fields.token)
			_, err := b.SendMessage(tt.args.msg, tt.args.chatId, tt.args.msgType, SendMessageOptions{
				DisablePreview:           false,
				DisableNotification:      false,
				ProtectContent:           false,
				ReplyToMessageID:         0,
				AllowSendingWithoutReply: false,
			})
			if (err != nil) != tt.wantErr {
				t.Fatalf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
