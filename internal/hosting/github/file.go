package github

import (
	"context"

	"github.com/google/go-github/v58/github"
	"github.com/thegalactiks/giteway/hosting"
)

func mapFile(c *github.RepositoryContent) *hosting.File {
	var size *int
	if c.GetType() != "dir" {
		size = c.Size
	}

	file := hosting.File{
		Type:     c.GetType(),
		Encoding: c.Encoding,
		Size:     size,
		Name:     c.GetName(),
		Path:     c.GetPath(),
	}

	return &file
}

func (h *HostingGithub) GetFiles(ctx context.Context, repo *hosting.Repository, path string) (*hosting.File, []hosting.File, error) {
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
