package telegram

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseTime(t *testing.T) {
	y, m, d := time.Now().In(LocalZone).Date()
	startDay := time.Date(y, m, d, 0, 0, 0, 0, LocalZone)
	type args struct {
		payload string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			args:    args{payload: "9:25"},
			wantErr: false,
			want:    startDay.Add(9*time.Hour + 25*time.Minute),
		},
		{
			args:    args{payload: "9h25"},
			wantErr: false,
			want:    startDay.Add(9*time.Hour + 25*time.Minute),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTime(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
