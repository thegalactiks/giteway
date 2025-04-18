package github

import (
	"context"

	"github.com/google/go-github/v68/github"
	"github.com/thegalactiks/giteway/hosting"
)

func mapCommitAuthor(a *github.CommitAuthor) *hosting.CommitAuthor {
	author := hosting.CommitAuthor{
		Date:  a.GetDate().Time,
		Name:  a.GetName(),
		Email: a.GetEmail(),
	}

	return &author
}

func mapCommit(c *github.Commit) *hosting.Commit {
	commit := hosting.Commit{
		SHA: c.GetSHA(),
		Tree: &hosting.CommitTree{
			SHA: c.GetTree().GetSHA(),
		},
		Author:    mapCommitAuthor(c.GetAuthor()),
		Committer: mapCommitAuthor(c.GetCommitter()),
		Message:   c.Message,
		Date:      &c.GetCommitter().Date.Time,
	}

	return &commit
}

func (h *GithubService) GetCommits(ctx context.Context, repo *hosting.Repository, opts *hosting.GetCommitsOpts) ([]hosting.Commit, error) {
	githubCommits, _, err := h.client.Repositories.ListCommits(ctx, repo.Owner, repo.Name, &github.CommitsListOptions{
		SHA: *opts.Ref,
	})
	if err != nil {
		return nil, err
	}

	var commits []hosting.Commit
	for _, c := range githubCommits {
		hostingCommit := mapCommit(c.GetCommit())
		commits = append(commits, *hostingCommit)
	}

	return commits, nil
}

func (h *GithubService) GetCommit(ctx context.Context, repo *hosting.Repository, opts *hosting.GetCommitsOpts) (*hosting.Commit, error) {
	githubCommit, _, err := h.client.Repositories.GetCommit(ctx, repo.Owner, repo.Name, *opts.Ref, &github.ListOptions{})
	if err != nil {
		return nil, err
	}

	return mapCommit(githubCommit.GetCommit()), nil
}
