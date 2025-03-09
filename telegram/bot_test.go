package telegram

import (
	"testing"
	"time"
)

func TestBot_SendMessage(t *testing.T) {
	type fields struct {
		token string
	}
	type args struct {
		msg    string
		chatId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				token: "",
			},
			args: args{
				msg:    time.Now().Format(time.DateTime),
				chatId: "",
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
			_, err := b.SendMessage(tt.args.msg, tt.args.chatId, Text)
			if (err != nil) != tt.wantErr {
				t.Fatalf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
