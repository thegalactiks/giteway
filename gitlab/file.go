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
		ID:   n.ID,
		Type: n.Type,
		Name: n.Name,
		Path: n.Path,
	}

	return &file
}

func mapFile(f *gitlab.File) *hosting.File {
	var encoding hosting.Encoding
	switch f.Encoding {
	case "base64":
		encoding = hosting.Base64Encoding
	default:
		encoding = hosting.TextEncoding
	}

	file := hosting.File{
		ID:       f.BlobID,
		Encoding: &encoding,
		Size:     &f.Size,
		Name:     f.FileName,
		Path:     f.FilePath,
	}

	return &file
}

func formatPath(path string) string {
	return strings.TrimLeft(path, "/")
}

func (h *GitlabService) GetFiles(ctx context.Context, repo *hosting.Repository, path string) (*hosting.File, []hosting.File, error) {
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

func (h *GitlabService) GetRawFile(ctx context.Context, repo *hosting.Repository, path string, opts *hosting.GetFileOpts) ([]byte, error) {
	pid := createPid(repo)

	file, _, err := h.client.RepositoryFiles.GetRawFile(pid, formatPath(path), &gitlab.GetRawFileOptions{
		Ref: opts.Ref,
	})

	return file, err
}

func (h *GitlabService) CreateFile(ctx context.Context, repo *hosting.Repository, file *hosting.File, opts *hosting.CreateFileOpts) (*hosting.File, *hosting.Commit, error) {
	pid := createPid(repo)
	branch := opts.Branch
	if opts.Ref != nil {
		branch = opts.Ref
	}

	gitlabFile, _, err := h.client.RepositoryFiles.CreateFile(pid, formatPath(file.Path), &gitlab.CreateFileOptions{
		Branch:        branch,
		Encoding:      file.GetEncoding(),
		Content:       file.Content,
		CommitMessage: opts.Commit.Message,
	})
	if err != nil {
		return nil, nil, err
	}

	commit, err := h.GetCommits(ctx, repo, &hosting.GetCommitsOpts{Ref: opts.Ref})
	if err != nil {
		return nil, nil, err
	}

	createdFile := hosting.File{
		Path: gitlabFile.FilePath,
	}
	lastCommit := commit[len(commit)-1]

	return &createdFile, &lastCommit, nil
}

func (h *GitlabService) UpdateFile(ctx context.Context, repo *hosting.Repository, file *hosting.File, opts *hosting.UpdateFileOpts) (*hosting.File, *hosting.Commit, error) {
	pid := createPid(repo)
	branch := opts.Branch
	if opts.Ref != nil {
		branch = opts.Ref
	}

	gitlabFile, _, err := h.client.RepositoryFiles.UpdateFile(pid, formatPath(file.Path), &gitlab.UpdateFileOptions{
		Branch:        branch,
		Encoding:      file.GetEncoding(),
		Content:       file.Content,
		CommitMessage: opts.Commit.Message,
	})
	if err != nil {
		return nil, nil, err
	}

	commit, err := h.GetCommits(ctx, repo, &hosting.GetCommitsOpts{Ref: opts.Ref})
	if err != nil {
		return nil, nil, err
	}

	createdFile := hosting.File{
		Path: gitlabFile.FilePath,
	}
	lastCommit := commit[len(commit)-1]

	return &createdFile, &lastCommit, nil
}

func (h *GitlabService) DeleteFile(ctx context.Context, repo *hosting.Repository, path string, opts *hosting.DeleteFileOpts) (*hosting.Commit, error) {
	pid := createPid(repo)
	branch := opts.Branch
	if opts.Ref != nil {
		branch = opts.Ref
	}

	_, err := h.client.RepositoryFiles.DeleteFile(pid, formatPath(path), &gitlab.DeleteFileOptions{
		Branch:        branch,
		CommitMessage: opts.Commit.Message,
	})
	if err != nil {
		return nil, err
	}

	commit, err := h.GetCommits(ctx, repo, &hosting.GetCommitsOpts{Ref: opts.Ref})
	if err != nil {
		return nil, err
	}

	lastCommit := commit[len(commit)-1]

	return &lastCommit, nil
}
