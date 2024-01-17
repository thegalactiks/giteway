package gitlab

import (
	"fmt"

	"github.com/thegalactiks/giteway/hosting"
	"github.com/xanzy/go-gitlab"
)

type HostingGitlab struct {
	client *gitlab.Client
}

var _ hosting.Hosting = (*HostingGitlab)(nil)

func New(token string) (*HostingGitlab, error) {
	client, err := gitlab.NewOAuthClient(token)
	if err != nil {
		return nil, err
	}

	return &HostingGitlab{
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
