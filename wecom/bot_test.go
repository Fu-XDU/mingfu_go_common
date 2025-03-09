package wecom

import (
	"testing"
	"time"
)

func TestBot_SendMessage(t *testing.T) {
	key := ""
	type fields struct {
		key string
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "DateTime",
			fields: fields{
				key: key,
			},
			args: args{
				msg: time.Now().Format(time.DateTime),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.fields.key) == 0 {
				t.Fatal("tt.fields.key can't be empty")
			}

			b := NewBot(tt.fields.key)
			if err := b.SendMessage(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
