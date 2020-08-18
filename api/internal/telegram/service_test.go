package telegram

import (
	"testing"
	"time"
)

func Test_toDateStr(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{t: time.Date(2020, 8, 19, 1, 2, 3, 4, time.UTC)},
			want: "2020/08/19",
		},
		{
			args: args{t: time.Date(2020, 8, 19, 1, 2, 3, 4, time.FixedZone("UTC+7", 7*60*60))},
			want: "2020/08/19",
		},
		{
			args: args{t: time.Date(2020, 8, 19, 19, 2, 3, 4, time.UTC)},
			want: "2020/08/20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toUTCDateStr(tt.args.t); got != tt.want {
				t.Errorf("toUTCDateStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOnTime(t *testing.T) {
	type args struct {
		at time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{at: time.Date(2020, 8, 19, 1, 2, 3, 4, time.UTC)},
			want: true,
		},
		{
			args: args{at: time.Date(2020, 8, 19, 1, 2, 3, 4, time.FixedZone("UTC+7", 7*60*60))},
			want: true,
		},
		{
			args: args{at: time.Date(2020, 8, 19, 19, 2, 3, 4, time.UTC)},
			want: true,
		},
		{
			args: args{at: time.Date(2020, 8, 19, 3, 2, 3, 4, time.UTC)},
			want: false,
		},
		{
			args: args{at: time.Date(2020, 8, 19, 9, 30, 3, 4, time.FixedZone("UTC+7", 7*60*60))},
			want: true,
		},
		{
			args: args{at: time.Date(2020, 8, 19, 9, 31, 3, 4, time.FixedZone("UTC+7", 7*60*60))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOnTime(tt.args.at); got != tt.want {
				t.Errorf("isOnTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
