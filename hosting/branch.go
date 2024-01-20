package hosting

type Branch struct {
	Name   string  `json:"name"`
	Commit *Commit `json:"commit,omitempty"`
}

type CreateBranchOpts struct {
	Ref    *string `json:"ref,omitempty"`
	SHA    *string `json:"sha,omitempty"`
	Branch *string `json:"branch,omitempty"`
}
