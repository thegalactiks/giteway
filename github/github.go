package github

import (
	"github.com/google/go-github/v59/github"
	"github.com/thegalactiks/giteway/hosting"
)

type GithubService struct {
	client   *github.Client
	hasToken bool
}

var _ hosting.GitHostingService = (*GithubService)(nil)

func NewGithubService(token *string) (*GithubService, error) {
	hasToken := false
	client := github.NewClient(nil)
	if token != nil {
		client = client.WithAuthToken(*token)
		hasToken = true
	}

	return &GithubService{
		client:   client,
		hasToken: hasToken,
	}, nil
}
