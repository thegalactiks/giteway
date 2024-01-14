package github

import (
	"context"

	"github.com/google/go-github/v58/github"
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

func mapCommit(c *github.RepositoryCommit) *hosting.Commit {
	commit := hosting.Commit{
		SHA:       c.GetSHA(),
		Author:    *mapCommitAuthor(c.GetCommit().GetAuthor()),
		Committer: *mapCommitAuthor(c.GetCommit().GetCommitter()),
		Message:   c.GetCommit().GetMessage(),
		Date:      c.GetCommit().GetCommitter().GetDate().Time,
	}

	return &commit
}

func (h *HostingGithub) GetCommits(ctx context.Context, repo *hosting.Repository) ([]hosting.Commit, error) {
	githubCommits, _, err := h.client.Repositories.ListCommits(ctx, repo.Owner, repo.Name, &github.CommitsListOptions{})
	if err != nil {
		return nil, err
	}

	var commits []hosting.Commit
	for _, c := range githubCommits {
		hostingCommit := mapCommit(c)
		commits = append(commits, *hostingCommit)
	}

	return commits, nil
}
