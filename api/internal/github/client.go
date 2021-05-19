package github

import (
	"context"
	"os"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

type Client struct {
	cli *github.Client
}

func (c Client) GithubClient() *github.Client {
	return c.cli
}

func NewClientFromEnv() *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	cli := github.NewClient(tc)
	return &Client{
		cli: cli,
	}
}
