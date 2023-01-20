package grpc

import (
	"reflect"
	"testing"
	"time"
)

var (
	firstFebruary2020ZeroHoursUTCTime = time.Date(2020, time.Month(2), 1, 0, 0, 0, 0, time.UTC)
	firstFebruary2020Unix             = int64(1580515200)
)

func Test_timeToUnix(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Success",
			args: args{
				date: firstFebruary2020ZeroHoursUTCTime,
			},
			want: firstFebruary2020Unix,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeToUnix(tt.args.date); got != tt.want {
				t.Errorf("timeToUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unixToTime(t *testing.T) {
	type args struct {
		unix int64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "Success",
			args: args{
				unix: firstFebruary2020Unix,
			},
			want: firstFebruary2020ZeroHoursUTCTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unixToTime(tt.args.unix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unixToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
