package point

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestRestService_UserPoints(t *testing.T) {
	svc := newRestServiceFromEnv()
	ctx := context.Background()
	ups, err := svc.UserPointsInMonth(ctx, 3, 2020)
	if err != nil {
		t.Fatalf("UserPointsInMonth failed: %v", err)
	}
	t.Logf("UserPointsInMonth result: %#v", ups)
}

func TestRestService_WorkingIssues(t *testing.T) {
	svc := newRestServiceFromEnv()
	ctx := context.Background()
	ups, err := svc.WorkingIssues(ctx, time.Now().Add(-14*24*time.Hour), time.Now())
	if err != nil {
		t.Fatalf("WorkingIssues failed: %v", err)
	}
	for _, u := range ups {
		t.Log(u.DisplayName)
		for _, i := range u.Issues {
			t.Logf("%v %v %v %v", i.Created.String(), i.Updated.String(), i.Resolved, i.Status)
		}
	}
}

func newRestServiceFromEnv() Service {
	username := os.Getenv("JIRA_USERNAME")
	password := os.Getenv("JIRA_PASSWORD")
	host := os.Getenv("JIRA_HOST")
	return NewRestService(username, password, host)
}

func Test_getFloat64(t *testing.T) {
	getAddress := func(v float64) *float64 {
		return &v
	}
	type args struct {
		v *float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "not nil", args: args{v: getAddress(5.6)}, want: 5.6},
		{name: "nil", args: args{v: nil}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFloat64(tt.args.v); got != tt.want {
				t.Errorf("getFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTimeBound(t *testing.T) {
	type args struct {
		month int
		year  int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{args: args{year: 2020, month: 3}, want: "2020-03-01", want1: "2020-04-01"},
		{args: args{year: 2020, month: 1}, want: "2020-01-01", want1: "2020-02-01"},
		{args: args{year: 2020, month: 12}, want: "2020-12-01", want1: "2021-01-01"},
		{args: args{year: 2020, month: 4}, want: "2020-04-01", want1: "2020-05-01"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getTimeBoundByMonth(tt.args.month, tt.args.year)
			if got != tt.want {
				t.Errorf("getTimeBoundByMonth() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getTimeBoundByMonth() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseUpdatedTime(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{args: args{s: "2020-03-17T15:51:27.884+0000"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseUpdatedTime(tt.args.s); got.Year() == 0 {
				t.Errorf("parseUpdatedTime() return zero")
			} else {
				t.Logf("%v", got)
			}
		})
	}
}
