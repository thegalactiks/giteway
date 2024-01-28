package hosting

import "time"

type CommitAuthor struct {
	Date  time.Time `json:"date"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type CommitTree struct {
	SHA string `json:"sha"`
}

type Commit struct {
	SHA       string        `json:"sha"`
	Tree      *CommitTree   `json:"tree,omitempty"`
	Author    *CommitAuthor `json:"author,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
	Message   *string       `json:"message,omitempty"`
	Date      *time.Time    `json:"date,omitempty"`
}

type GetCommitsOpts struct {
	Ref *string `json:"ref,omitempty"`
}
