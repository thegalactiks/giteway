package hosting

import (
	"context"
)

type GitHostingService interface {
	GetRepositories(ctx context.Context, owner string) ([]Repository, error)
	GetRepository(ctx context.Context, owner string, repo string) (*Repository, error)
	GetBranches(ctx context.Context, repo *Repository) ([]Branch, error)
	CreateBranch(ctx context.Context, repo *Repository, opts *CreateBranchOpts) (*Branch, error)
	DeleteBranch(ctx context.Context, repo *Repository, branch *Branch) error
	GetCommits(ctx context.Context, repo *Repository, opts *GetCommitsOpts) ([]Commit, error)
	GetFiles(ctx context.Context, repo *Repository, path string) (*File, []File, error)
	GetRawFile(ctx context.Context, repo *Repository, path string, opts *GetFileOpts) ([]byte, error)
	CreateFile(ctx context.Context, repo *Repository, file *File, opts *CreateFileOpts) (*File, *Commit, error)
	UpdateFile(ctx context.Context, repo *Repository, file *File, opts *UpdateFileOpts) (*File, *Commit, error)
	DeleteFile(ctx context.Context, repo *Repository, path string, opts *DeleteFileOpts) (*Commit, error)
}
