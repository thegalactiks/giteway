package gitlab

import (
	"context"

	"github.com/thegalactiks/giteway/hosting"
	"github.com/xanzy/go-gitlab"
)

func mapBranch(b *gitlab.Branch) *hosting.Branch {
	branch := hosting.Branch{
		Name:   b.Name,
		Commit: mapCommit(b.Commit),
	}

	return &branch
}

func (h *HostingGitlab) GetBranches(ctx context.Context, repo *hosting.Repository) ([]hosting.Branch, error) {
	gitlabBranches, _, err := h.client.Branches.ListBranches(createPid(repo), &gitlab.ListBranchesOptions{})
	if err != nil {
		return nil, err
	}

	var branches []hosting.Branch
	for _, b := range gitlabBranches {
		hostingBranch := mapBranch(b)
		branches = append(branches, *hostingBranch)
	}

	return branches, nil
}
