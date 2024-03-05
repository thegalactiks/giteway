package gitlab

import (
	"fmt"
	"net/http"

	"github.com/thegalactiks/giteway/hosting"
	"github.com/xanzy/go-gitlab"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type GitlabService struct {
	client *gitlab.Client
}

var _ hosting.GitHostingService = (*GitlabService)(nil)

func NewGitlabService(token string) (*GitlabService, error) {
	httpClient := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	client, err := gitlab.NewOAuthClient(token, gitlab.WithHTTPClient(httpClient))
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
