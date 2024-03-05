package github

import (
	"net/http"

	"github.com/google/go-github/v60/github"
	"github.com/thegalactiks/giteway/hosting"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type GithubService struct {
	client   *github.Client
	hasToken bool
}

var _ hosting.GitHostingService = (*GithubService)(nil)

func NewGithubService(token *string) (*GithubService, error) {
	hasToken := false
	httpClient := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	client := github.NewClient(httpClient)
	if token != nil {
		client = client.WithAuthToken(*token)
		hasToken = true
	}

	return &GithubService{
		client:   client,
		hasToken: hasToken,
	}, nil
}
