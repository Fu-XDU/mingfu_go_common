package datetime

import (
	"testing"
	"time"
)

func TestBirthdateToYearMonthString(t *testing.T) {
	type args struct {
		birthdate *time.Time
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := BirthdateToYearMonthString(tt.args.birthdate); gotOutput != tt.wantOutput {
				t.Errorf("BirthdateToYearMonthString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestStartEarlierThanEnd(t *testing.T) {
	type args struct {
		startDateStr string
		endDateStr   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				startDateStr: "2020-01",
				endDateStr:   "2020-01",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				startDateStr: "2020-01",
				endDateStr:   "2022-01",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				startDateStr: "2020-01",
				endDateStr:   "2019-12",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartEarlierThanEnd(tt.args.startDateStr, tt.args.endDateStr); got != tt.want {
				t.Errorf("StartEarlierThanEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearMonthStringToDatetime(t *testing.T) {
	type args struct {
		dateStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				dateStr: "2020-01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBirthdate := YearMonthStringToDatetime(tt.args.dateStr); gotBirthdate == nil {
				t.Errorf("YearMonthStringToDatetime() should not return nil")
			}
		})
	}
}
