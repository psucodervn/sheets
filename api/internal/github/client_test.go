package github

import (
	"context"
	"testing"
)

func TestNewClientFromEnv(t *testing.T) {
	cli := NewClientFromEnv()
	if cli == nil {
		t.Error("NewClientFromEnv returns nil")
		return
	}
	orgs, _, err := cli.GithubClient().Issues.ListByOrg(context.Background(), "verichains", nil)
	if err != nil {
		t.Fatalf("list orgs: %v", err)
	}
	t.Logf("orgs: %v", orgs)
	repos, _, err := cli.GithubClient().Repositories.List(context.Background(), "psucodervn", nil)
	for _, repo := range repos {
		t.Logf("name = %s, visiblity = %s", repo.GetName(), repo.GetLanguage())
	}
}
