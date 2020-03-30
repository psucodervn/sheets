package point

import (
	"context"
	"os"
	"testing"
)

func TestRestService_UserPoints(t *testing.T) {
	svc := newRestServiceFromEnv()
	ctx := context.Background()
	ups, err := svc.UserPoints(ctx, 3, 2020)
	if err != nil {
		t.Fatalf("UserPoints failed: %v", err)
	}
	t.Logf("UserPoints result: %#v", ups)
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
