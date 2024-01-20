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
	SHA       *string      `json:"sha,omitempty"`
	Tree      CommitTree   `json:"tree,omitempty"`
	Author    CommitAuthor `json:"author"`
	Committer CommitAuthor `json:"committer"`
	Message   string       `json:"message"`
	Date      time.Time    `json:"date"`
}

type GetCommitsOpts struct {
	Ref *string `json:"ref,omitempty"`
}
