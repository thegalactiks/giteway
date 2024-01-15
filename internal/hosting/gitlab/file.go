package gitlab

import (
	"context"
	"net/http"
	"strings"

	"github.com/thegalactiks/giteway/hosting"
	"github.com/xanzy/go-gitlab"
)

func mapTreeNode(n *gitlab.TreeNode) *hosting.File {
	file := hosting.File{
		Type: n.Type, // TODO: map type
		Name: n.Name,
		Path: n.Path,
	}

	return &file
}

func mapFile(c *gitlab.File) *hosting.File {
	var size *int
	// if c.GetType() != "dir" {
	// 	size = &c.Size
	// }

	file := hosting.File{
		// Type:     c.GetType(),
		Encoding: &c.Encoding,
		Size:     size,
		Name:     c.FileName,
		Path:     c.FilePath,
	}

	return &file
}

func (h *HostingGitlab) GetFiles(ctx context.Context, repo *hosting.Repository, path string) (*hosting.File, []hosting.File, error) {
	pid := createPid(repo)
	pathWithoutSlash := strings.TrimLeft(path, "/")
	ref := "master"

	fileContent, resp, err := h.client.RepositoryFiles.GetFileMetaData(pid, pathWithoutSlash, &gitlab.GetFileMetaDataOptions{
		Ref: &ref,
	})
	if fileContent != nil {
		return mapFile(fileContent), nil, nil
	} else if err != nil && resp.StatusCode != http.StatusNotFound {
		return nil, nil, err
	}

	escapedPath := gitlab.PathEscape(pathWithoutSlash)
	treeNodes, _, err := h.client.Repositories.ListTree(pid, &gitlab.ListTreeOptions{
		Path:      &escapedPath,
		Recursive: newFalse(),
	})
	if err != nil {
		return nil, nil, err
	}

	var files []hosting.File
	for _, n := range treeNodes {
		hostingFile := mapTreeNode(n)
		files = append(files, *hostingFile)
	}

	return nil, files, nil
}
