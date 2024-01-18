package gitlab

import (
	"context"

	"github.com/thegalactiks/giteway/hosting"
	"github.com/xanzy/go-gitlab"
)

func mapCommit(c *gitlab.Commit) *hosting.Commit {
	commit := hosting.Commit{
		SHA: c.ID,
		Author: hosting.CommitAuthor{
			Date:  *c.AuthoredDate,
			Name:  c.AuthorName,
			Email: c.AuthorEmail,
		},
		Committer: hosting.CommitAuthor{
			Date:  *c.CommittedDate,
			Name:  c.CommitterName,
			Email: c.CommitterEmail,
		},
		Message: c.Message,
		Date:    *c.CommittedDate,
	}

	return &commit
}

func (h *HostingGitlab) GetCommits(ctx context.Context, repo *hosting.Repository, opts *hosting.GetCommitsOpts) ([]hosting.Commit, error) {
	h.client.Projects.GetProject(ctx, &gitlab.GetProjectOptions{})
	gitlabCommits, _, err := h.client.Commits.ListCommits(createPid(repo), &gitlab.ListCommitsOptions{
		RefName: opts.Ref,
	})
	if err != nil {
		return nil, err
	}

	var commits []hosting.Commit
	for _, c := range gitlabCommits {
		hostingCommit := mapCommit(c)
		commits = append(commits, *hostingCommit)
	}

	return commits, nil
}
