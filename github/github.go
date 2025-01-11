package github

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v68/github"
	"github.com/thegalactiks/giteway/hosting"
	"github.com/thegalactiks/giteway/internal/config"
	internalhttp "github.com/thegalactiks/giteway/internal/http"
)

type GithubService struct {
	cfg    *config.GithubConfig
	client *github.Client
	user   *github.User
}

var _ hosting.GitHostingService = (*GithubService)(nil)

func NewGithubService(cfg *config.Config) *GithubService {
	client := github.NewClient(internalhttp.NewHttpClient(nil))
	return &GithubService{
		cfg:    &cfg.GithubConfig,
		client: client,
	}
}

func (s *GithubService) IsKnownInstallation(owner string) bool {
	_, ok := s.cfg.Installations[owner]
	return ok
}

func (s *GithubService) WithInstallationOwner(owner string) (*GithubService, error) {
	installation, ok := s.cfg.Installations[owner]
	if !ok {
		return nil, fmt.Errorf("unknown installation owner: %s", owner)
	}

	return s.WithInstallation(installation.ID)
}

func (s *GithubService) WithInstallation(installationID int64) (*GithubService, error) {
	// load private key file from path
	privateKey, err := os.ReadFile(s.cfg.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	var httpTransport http.RoundTripper
	itr, err := ghinstallation.New(http.DefaultTransport, s.cfg.AppID, installationID, privateKey)
	if err != nil {
		return nil, err
	}
	httpTransport = itr

	client := github.NewClient(internalhttp.NewHttpClient(&httpTransport))
	return &GithubService{client: client}, nil
}

func (s *GithubService) WithAuthToken(ctx context.Context, token string) (*GithubService, error) {
	client := s.client.WithAuthToken(token)
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	return &GithubService{
		client: client,
		user:   user,
	}, nil
}
