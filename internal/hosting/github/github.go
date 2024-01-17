package github

import (
	"github.com/google/go-github/v58/github"
	"github.com/thegalactiks/giteway/hosting"
)

type HostingGithub struct {
	client   *github.Client
	hasToken bool
}

var _ hosting.Hosting = (*HostingGithub)(nil)

func New(token *string) (*HostingGithub, error) {
	hasToken := false
	client := github.NewClient(nil)
	if token != nil {
		client = client.WithAuthToken(*token)
		hasToken = true
	}

	return &HostingGithub{
		client:   client,
		hasToken: hasToken,
	}, nil
}
