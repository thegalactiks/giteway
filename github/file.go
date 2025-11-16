package github

import (
	"context"
	"errors"
	"strings"

	"github.com/google/go-github/v79/github"
	"github.com/thegalactiks/giteway/hosting"
)

func mapFile(c *github.RepositoryContent) *hosting.File {
	var size *int
	if c.GetType() != "dir" {
		size = c.Size
	}

	var encoding hosting.Encoding
	switch c.GetEncoding() {
	case "base64":
		encoding = hosting.Base64Encoding
	case "none":
		encoding = hosting.NoneEncoding
	}

	file := hosting.File{
		ID:       c.GetSHA(),
		Type:     c.GetType(),
		Encoding: &encoding,
		Content:  c.Content,
		Size:     size,
		Name:     c.GetName(),
		Path:     c.GetPath(),
	}

	return &file
}

func formatPath(path string) string {
	return strings.TrimLeft(path, "/")
}

func (h *GithubService) GetFiles(ctx context.Context, repo *hosting.Repository, path string) (*hosting.File, []hosting.File, error) {
	fileContent, dirContent, _, err := h.client.Repositories.GetContents(ctx, repo.Owner, repo.Name, path, &github.RepositoryContentGetOptions{})
	if err != nil {
		return nil, nil, err
	}

	if fileContent != nil {
		return mapFile(fileContent), nil, err
	}

	var files []hosting.File
	for _, f := range dirContent {
		hostingFile := mapFile(f)
		files = append(files, *hostingFile)
	}

	return nil, files, nil
}

func (h *GithubService) GetRawFile(ctx context.Context, repo *hosting.Repository, path string, opts *hosting.GetFileOpts) ([]byte, error) {
	fileContent, _, _, err := h.client.Repositories.GetContents(ctx, repo.Owner, repo.Name, path, &github.RepositoryContentGetOptions{})
	if err != nil {
		return nil, err
	}

	if fileContent == nil {
		return nil, errors.New("the path should be a valid file path")
	}

	c, err := fileContent.GetContent()
	if err != nil {
		return nil, err
	}

	return []byte(c), nil
}

func (h *GithubService) CreateFile(ctx context.Context, repo *hosting.Repository, file *hosting.File, opts *hosting.CreateFileOpts) (*hosting.File, *hosting.Commit, error) {
	branch := opts.Branch
	if opts.Ref != nil {
		branch = opts.Ref
	}

	githubContentResponse, _, err := h.client.Repositories.CreateFile(ctx, repo.Owner, repo.Name, formatPath(file.Path), &github.RepositoryContentFileOptions{
		Branch:  branch,
		Content: []byte(*file.Content),
		Message: opts.Message,
	})
	if err != nil {
		return nil, nil, err
	}

	return mapFile(githubContentResponse.GetContent()), mapCommit(&githubContentResponse.Commit), nil
}

func (h *GithubService) getFileSHA(ctx context.Context, repo *hosting.Repository, path string) (*string, error) {
	file, _, err := h.GetFiles(ctx, repo, path)
	if err != nil {
		return nil, err
	}

	return &file.ID, nil
}

func (h *GithubService) UpdateFile(ctx context.Context, repo *hosting.Repository, file *hosting.File, opts *hosting.UpdateFileOpts) (*hosting.File, *hosting.Commit, error) {
	path := formatPath(file.Path)

	branch := opts.Branch
	if opts.Ref != nil {
		branch = opts.Ref
	}

	sha, err := h.getFileSHA(ctx, repo, path)
	if err != nil {
		return nil, nil, err
	}

	githubContentResponse, _, err := h.client.Repositories.UpdateFile(ctx, repo.Owner, repo.Name, path, &github.RepositoryContentFileOptions{
		Branch:  branch,
		SHA:     sha,
		Content: []byte(*file.Content),
		Message: opts.Message,
	})
	if err != nil {
		return nil, nil, err
	}

	return mapFile(githubContentResponse.GetContent()), mapCommit(&githubContentResponse.Commit), nil
}

func (h *GithubService) DeleteFile(ctx context.Context, repo *hosting.Repository, path string, opts *hosting.DeleteFileOpts) (*hosting.Commit, error) {
	path = formatPath(path)

	branch := opts.Branch
	if opts.Ref != nil {
		branch = opts.Ref
	}

	sha, err := h.getFileSHA(ctx, repo, path)
	if err != nil {
		return nil, err
	}

	githubContentResponse, _, err := h.client.Repositories.DeleteFile(ctx, repo.Owner, repo.Name, formatPath(path), &github.RepositoryContentFileOptions{
		Branch:  branch,
		SHA:     sha,
		Message: opts.Message,
	})
	if err != nil {
		return nil, err
	}

	return mapCommit(&githubContentResponse.Commit), nil
}
