package auth

import (
	"testing"
)

func TestPasswordHash(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
	}{
		{
			name:     "Test 1",
			password: "password",
			want:     "9f44fd7f261862d05f846585c4b254f778211b2ab35b830f841ecaba4d3f22da",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := PasswordHash(tt.password, "fd80260a706445a7450835998fa3e5e2f171ac560feccd863377b474d68ff470"); got != tt.want {
				t.Errorf("PasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
