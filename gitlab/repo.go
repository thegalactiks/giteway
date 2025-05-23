package gitlab

import (
	"context"

	"github.com/thegalactiks/giteway/hosting"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

func mapProject(owner string, p *gitlab.Project) *hosting.Repository {
	repo := hosting.Repository{
		Owner:         owner,
		Name:          p.Name,
		DefaultBranch: p.DefaultBranch,
		CloneURL:      p.SSHURLToRepo,
		GitURL:        p.HTTPURLToRepo,
		UpdatedAt:     *p.LastActivityAt,
		CreatedAt:     *p.CreatedAt,
	}

	return &repo
}

func (h *GitlabService) GetRepositories(ctx context.Context, owner string) ([]hosting.Repository, error) {
	gitlabProjects, _, err := h.client.Groups.ListGroupProjects(owner, &gitlab.ListGroupProjectsOptions{})
	if err != nil {
		return nil, err
	}

	var repos []hosting.Repository
	for _, p := range gitlabProjects {
		hostingRepo := mapProject(owner, p)
		repos = append(repos, *hostingRepo)
	}

	return repos, nil
}

func (h *GitlabService) GetRepository(ctx context.Context, owner string, repo string) (*hosting.Repository, error) {
	gitlabProject, _, err := h.client.Projects.GetProject(createPid(&hosting.Repository{
		Owner: owner,
		Name:  repo,
	}), &gitlab.GetProjectOptions{})
	if err != nil {
		return nil, err
	}

	hostingRepo := mapProject(owner, gitlabProject)
	return hostingRepo, nil
}
