package github

import (
	"context"

	"github.com/google/go-github/v58/github"
	"github.com/thegalactiks/giteway/hosting"
)

func mapRepository(owner string, r *github.Repository) *hosting.Repository {
	repo := hosting.Repository{
		Owner:         owner,
		Name:          r.GetName(),
		DefaultBranch: r.GetDefaultBranch(),
		CloneURL:      r.GetCloneURL(),
		GitURL:        r.GetGitURL(),
		CreatedAt:     r.GetCreatedAt().Time,
		PushedAt:      r.GetPushedAt().Time,
		UpdatedAt:     r.GetUpdatedAt().Time,
	}

	return &repo
}

func (h *HostingGithub) GetRepositories(ctx context.Context, owner string) ([]hosting.Repository, error) {
	githubRepos, _, err := h.client.Repositories.ListByUser(ctx, owner, &github.RepositoryListByUserOptions{})
	if err != nil {
		return nil, err
	}

	var repos []hosting.Repository
	for _, r := range githubRepos {
		hostingRepo := mapRepository(owner, r)
		repos = append(repos, *hostingRepo)
	}

	return repos, nil
}

func (h *HostingGithub) GetRepository(ctx context.Context, owner string, repo string) (*hosting.Repository, error) {
	githubRepo, _, err := h.client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, err
	}

	hostingRepo := mapRepository(owner, githubRepo)
	return hostingRepo, nil
}
