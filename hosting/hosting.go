package hosting

import (
	"context"
	"time"
)

type Repository struct {
	Owner         string    `json:"owner"`
	Name          string    `json:"name"`
	DefaultBranch string    `json:"default_branch"`
	CloneURL      string    `json:"clone_url"`
	GitURL        string    `json:"git_url"`
	CreatedAt     time.Time `json:"created_at"`
	PushedAt      time.Time `json:"pushed_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Branch struct {
	Name   string  `json:"name"`
	Commit *Commit `json:"commit,omitempty"`
}

type File struct {
	Type     string  `json:"type"`
	Content  *string `json:"content,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	Size     *int    `json:"size,omitempty"`
	Name     string  `json:"name"`
	Path     string  `json:"path"`
}

type CommitAuthor struct {
	Date  time.Time `json:"date"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
type Commit struct {
	SHA       string       `json:"sha"`
	Author    CommitAuthor `json:"author"`
	Committer CommitAuthor `json:"committer"`
	Message   string       `json:"message"`
	Date      time.Time    `json:"date"`
}

type GetCommitsOpts struct {
	Ref *string `json:"ref"`
}

type CreateBranchOpts struct {
	// Ref from where the branch will be created.
	Ref *string `json:"ref"`

	// SHA from where the branch will be created.
	SHA *string `json:"sha"`

	// Branch name to create
	Branch *string `json:"branch"`
}

type Hosting interface {
	GetRepositories(ctx context.Context, owner string) ([]Repository, error)
	GetRepository(ctx context.Context, owner string, repo string) (*Repository, error)
	GetBranches(ctx context.Context, repo *Repository) ([]Branch, error)
	CreateBranch(ctx context.Context, repo *Repository, opts *CreateBranchOpts) (*Branch, error)
	GetCommits(ctx context.Context, repo *Repository, opts *GetCommitsOpts) ([]Commit, error)
	GetFiles(ctx context.Context, repo *Repository, path string) (*File, []File, error)
	GetRawFile(ctx context.Context, repo *Repository, path string) ([]byte, error)
}
