package gitlab

import (
	"fmt"

	"github.com/thegalactiks/giteway/hosting"
	"github.com/thegalactiks/giteway/internal/http"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

type GitlabService struct {
	client *gitlab.Client
}

var _ hosting.GitHostingService = (*GitlabService)(nil)

func NewGitlabService(token string) (*GitlabService, error) {
	client, err := gitlab.NewOAuthClient(token, gitlab.WithHTTPClient(http.NewHttpClient(nil)))
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

func newFalse() *bool {
	b := false
	return &b
}
