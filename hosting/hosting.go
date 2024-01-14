package hosting

import "context"

type Repository struct{}
type Branch struct{}
type File struct{}
type Commit struct{}

type Hosting interface {
	GetRepositories(ctx *context.Context) ([]Repository, error)
	GetRepository(ctx *context.Context, url string) (*Repository, error)
	GetBranches(ctx *context.Context, repo *Repository) (*Branch, error)
	GetFiles(ctx *context.Context, repo *Repository, branch *Branch) ([]File, error)
	GetCommits(ctx *context.Context, repo *Repository, branch *Branch) ([]Commit, error)
}
