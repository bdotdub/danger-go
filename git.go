package danger

type Git struct {
	ModifiedFiles []string `json:"modified_files"`
	CreatedFiles  []string `json:"created_files"`
	DeletedFiles  []string `json:"deleted_files"`
}

type GitCommit struct {
	SHA       string          `json:"sha"`
	Author    GitCommitAuthor `json:"author"`
	Committer GitCommitAuthor `json:"committer"`
	Message   string          `json:"message"`
	Parents   []string        `json:"parents"`
	URL       string          `json:"url"`
}

type GitCommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}
