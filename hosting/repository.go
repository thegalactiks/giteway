package hosting

import "time"

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
