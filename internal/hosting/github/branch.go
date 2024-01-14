package github

import (
	"context"

	"github.com/google/go-github/v58/github"
	"github.com/thegalactiks/giteway/hosting"
)

func mapBranch(b *github.Branch) *hosting.Branch {
	branch := hosting.Branch{
		Name:   b.GetName(),
		Commit: nil,
	}

	return &branch
}

func (h *HostingGithub) GetBranches(ctx context.Context, repo *hosting.Repository) ([]hosting.Branch, error) {
	githubBranches, _, err := h.client.Repositories.ListBranches(ctx, repo.Owner, repo.Name, &github.BranchListOptions{})
	if err != nil {
		return nil, err
	}

	var branches []hosting.Branch
	for _, b := range githubBranches {
		hostingBranch := mapBranch(b)
		branches = append(branches, *hostingBranch)
	}

	return branches, nil
}
