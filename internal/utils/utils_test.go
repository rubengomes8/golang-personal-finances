package utils

import (
	"reflect"
	"testing"
	"time"
)

var (
	firstFebruary2020ZeroHoursUTCTime = time.Date(2020, time.Month(2), 1, 0, 0, 0, 0, time.UTC)
	firstFebruary2020String           = "2020-02-01"
)

func Test_timeToStringDate(t *testing.T) {
	type args struct {
		t time.Time
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Successful",
			args: args{
				t: firstFebruary2020ZeroHoursUTCTime,
			},
			want: firstFebruary2020String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToStringDate(tt.args.t); got != tt.want {
				t.Errorf("timeToStringDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dateStringToTime(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				date: firstFebruary2020String,
			},
			want:    firstFebruary2020ZeroHoursUTCTime,
			wantErr: false,
		},
		{
			name: "ErrorWrongDateLayout",
			args: args{
				date: "2020-Feb-01",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateStringToTime(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("dateStringToTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dateStringToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
