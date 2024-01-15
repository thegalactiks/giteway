package github

import (
	"github.com/google/go-github/v58/github"
	"github.com/thegalactiks/giteway/hosting"
)

type HostingGithub struct {
	client *github.Client
}

var _ hosting.Hosting = (*HostingGithub)(nil)

func New(token *string) (*HostingGithub, error) {
	return &HostingGithub{
		client: github.NewClient(nil),
	}, nil
}
