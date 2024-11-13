package context

import (
	"testing"
)

func TestNewContextWithUuid(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{uuid: "5ce1c3e64a9a80b602ad172b2f3aa4c7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtx := NewContextWithUuid(tt.args.uuid)
			uuid := GetUuidFromContext(gotCtx)
			if uuid != tt.args.uuid {
				t.Errorf("GetUuidFromContext() = %v, want %v", uuid, tt.args.uuid)
			}
		})
	}
}
