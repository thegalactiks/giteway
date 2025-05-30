package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v68/github"
	"github.com/thegalactiks/giteway/hosting"
)

func mapBranch(b *github.Branch) *hosting.Branch {
	branch := hosting.Branch{
		Name: b.GetName(),
	}
	c := b.Commit
	if c != nil {
		branch.Commit = &hosting.Commit{
			SHA: c.GetSHA(),
		}
	}

	return &branch
}

func mapBranchRef(r *github.Reference) *hosting.Branch {
	branch := hosting.Branch{
		Name: r.GetRef(),
		Commit: &hosting.Commit{
			SHA: r.GetObject().GetSHA(),
		},
	}

	return &branch
}

func getBranchRefFromName(name string) string {
	return fmt.Sprintf("heads/%v", strings.TrimPrefix(name, "heads/"))
}

func (h *GithubService) GetBranches(ctx context.Context, repo *hosting.Repository) ([]hosting.Branch, error) {
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

func (h *GithubService) CreateBranch(ctx context.Context, repo *hosting.Repository, opts *hosting.CreateBranchOpts) (*hosting.Branch, error) {
	var githubRef = github.Reference{}
	if opts.SHA != nil {
		githubRef.Object = &github.GitObject{
			SHA: opts.SHA,
		}
	} else {
		var ref string
		if opts.Ref != nil {
			ref = *opts.Ref
		} else {
			githubRepo, _, err := h.client.Repositories.Get(ctx, repo.Owner, repo.Name)
			if err != nil {
				return nil, err
			}

			ref = getBranchRefFromName(githubRepo.GetDefaultBranch())
		}

		commit, _, err := h.client.Git.GetRef(ctx, repo.Owner, repo.Name, ref)
		if err != nil {
			return nil, err
		}

		githubRef.Object = commit.Object
	}

	branchRef := getBranchRefFromName(*opts.Branch)
	githubRef.Ref = &branchRef

	githubBranchRef, _, err := h.client.Git.CreateRef(ctx, repo.Owner, repo.Name, &githubRef)
	if err != nil {
		return nil, err
	}

	return mapBranchRef(githubBranchRef), nil
}

func (h *GithubService) DeleteBranch(ctx context.Context, repo *hosting.Repository, branch *hosting.Branch) error {
	branchRef := getBranchRefFromName(branch.Name)
	_, err := h.client.Git.DeleteRef(ctx, repo.Owner, repo.Name, branchRef)

	return err
}
