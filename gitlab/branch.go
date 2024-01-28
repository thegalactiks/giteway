package gitlab

import (
	"context"
	"fmt"

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

func (h *GitlabService) GetBranches(ctx context.Context, repo *hosting.Repository) ([]hosting.Branch, error) {
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

func (h *GitlabService) CreateBranch(ctx context.Context, repo *hosting.Repository, opts *hosting.CreateBranchOpts) (*hosting.Branch, error) {
	pid := createPid(repo)

	var ref *string
	if opts.Ref != nil {
		ref = opts.Ref
	} else if opts.SHA != nil {
		ref = opts.SHA
	}

	if ref == nil {
		gitlabRepo, _, err := h.client.Projects.GetProject(pid, &gitlab.GetProjectOptions{})
		if err != nil {
			return nil, err
		}

		defaultBranchRef := fmt.Sprintf("heads/%v", gitlabRepo.DefaultBranch)
		ref = &defaultBranchRef
	}

	gitlabBranch, _, err := h.client.Branches.CreateBranch(pid, &gitlab.CreateBranchOptions{Ref: ref, Branch: opts.Branch})
	if err != nil {
		return nil, err
	}

	return mapBranch(gitlabBranch), nil
}

func (h *GitlabService) DeleteBranch(ctx context.Context, repo *hosting.Repository, branch *hosting.Branch) error {
	_, err := h.client.Branches.DeleteBranch(createPid(repo), branch.Name)

	return err
}
