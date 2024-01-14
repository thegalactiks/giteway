package github

import (
	"context"

	"github.com/google/go-github/v58/github"
	"github.com/thegalactiks/giteway/hosting"
)

type HostingGithub struct {
	client *github.Client
}

var _ hosting.Hosting = (*HostingGithub)(nil)

func New() *HostingGithub {
	return &HostingGithub{
		client: github.NewClient(nil),
	}
}

func (*HostingGithub) GetFiles(ctx context.Context, repo *hosting.Repository, branch *hosting.Branch) ([]hosting.File, error) {
	return []hosting.File{}, nil
}

func (*HostingGithub) GetCommits(ctx context.Context, repo *hosting.Repository, branch *hosting.Branch) ([]hosting.Commit, error) {
	return []hosting.Commit{}, nil
}
