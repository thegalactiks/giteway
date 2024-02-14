package gitlab

import (
	"fmt"

	"github.com/thegalactiks/giteway/hosting"
	"github.com/xanzy/go-gitlab"
)

type GitlabService struct {
	client *gitlab.Client
}

var _ hosting.GitHostingService = (*GitlabService)(nil)

func NewGitlabService(token string) (*GitlabService, error) {
	client, err := gitlab.NewOAuthClient(token)
	if err != nil {
		return nil, err
	}

	return &GitlabService{
		client: client,
	}, nil
}

func createPid(repo *hosting.Repository) string {
	return fmt.Sprintf("%s/%s", repo.Owner, repo.Name)
}

func newTrue() *bool {
	b := true
	return &b
}

func newFalse() *bool {
	b := false
	return &b
}
