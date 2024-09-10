package github

import (
	"context"

	"github.com/google/go-github/v64/github"
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
		UpdatedAt:     r.GetUpdatedAt().Time,
	}

	return &repo
}

func mapRepositories(owner string, rs []*github.Repository) []hosting.Repository {
	var repos []hosting.Repository
	for _, r := range rs {
		hostingRepo := mapRepository(owner, r)
		repos = append(repos, *hostingRepo)
	}

	return repos
}

func (h *GithubService) GetRepositories(ctx context.Context, owner string) ([]hosting.Repository, error) {
	org, _, err := h.client.Organizations.Get(ctx, owner)
	if err == nil && org.GetLogin() == owner {
		githubRepos, _, err := h.client.Repositories.ListByOrg(ctx, owner, &github.RepositoryListByOrgOptions{
			Sort: "updated",
		})
		if err != nil {
			return nil, err
		}

		return mapRepositories(owner, githubRepos), nil
	}

	if h.user != nil && h.user.GetLogin() == owner {
		githubRepos, _, err := h.client.Repositories.ListByAuthenticatedUser(ctx, &github.RepositoryListByAuthenticatedUserOptions{
			Sort: "updated",
		})
		if err != nil {
			return nil, err
		}

		return mapRepositories(owner, githubRepos), nil
	}

	githubRepos, _, err := h.client.Repositories.ListByUser(ctx, owner, &github.RepositoryListByUserOptions{
		Sort: "updated",
	})
	if err != nil {
		return nil, err
	}

	return mapRepositories(owner, githubRepos), nil
}

func (h *GithubService) GetRepository(ctx context.Context, owner string, repo string) (*hosting.Repository, error) {
	githubRepo, _, err := h.client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, err
	}

	hostingRepo := mapRepository(owner, githubRepo)
	return hostingRepo, nil
}
