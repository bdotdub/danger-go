package danger

type GitHub struct {
	Issue              GitHubIssue       `json:"issue"`
	PullRequest        GitHubPullRequest `json:"pr"`
	Commits            []GitHubCommit    `json:"commits"`
	Reviews            []GitHubReview    `json:"reviews"`
	RequestedReviewers []GitHubUser      `json:"requested_reviewers"`
}

type GitHubPullRequest struct {
	Number             int            `json:"number"`
	Title              string         `json:"title"`
	Body               string         `json:"body"`
	User               GitHubUser     `json:"user"`
	Assignee           GitHubUser     `json:"assignee"`
	Assignees          []GitHubUser   `json:"assignees"`
	CreatedAt          string         `json:"created_at"`
	UpdatedAt          string         `json:"updated_at"`
	ClosedAt           string         `json:"closed_at"`
	MergedAt           string         `json:"merged_at"`
	Head               GitHubMergeRef `json:"head"`
	Base               GitHubMergeRef `json:"base"`
	State              string         `json:"state"`
	IsLocked           bool           `json:"locked"`
	IsMerged           bool           `json:"merged"`
	CommitCount        int            `json:"commit_count"`
	CommentCount       int            `json:"comment_count"`
	ReviewCommentCount int            `json:"review_comments"`
	Additions          int            `json:"additions"`
	Deletions          int            `json:"deletions"`
	ChangedFiles       int            `json:"changed_files"`
}

type GitHubUser struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	UserType string `json:"type"`
}

type GitHubMergeRef struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	SHA   string     `json:"sha"`
	User  GitHubUser `json:"user"`
	Repo  GitHubRepo `json:"repo"`
}

type GitHubRepo struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	FullName    string     `json:"full_name"`
	Owner       GitHubUser `json:"owner"`
	IsPrivate   bool       `json:"private"`
	Description string     `json:"description"`
	IsFork      bool       `json:"fork"`
	HTMLURL     string     `json:"html_url"`
}

type GitHubReview struct {
	User     GitHubUser `json:"user"`
	ID       int        `json:"id"`
	CommitID string     `json:"commit_id"`
	State    string     `json:"state"`
}

type GitHubCommit struct {
	sha       string     `json:"sha"`
	commit    GitCommit  `json:"commit"`
	URL       string     `json:"url"`
	Author    GitHubUser `json:"author"`
	Committer GitHubUser `json:"committer"`
}

type GitHubIssue struct {
	Number       int                `json:"number"`
	Title        string             `json:"title"`
	User         GitHubUser         `json:"user"`
	State        string             `json:"state"`
	IsLocked     bool               `json:"locked"`
	Body         string             `json:"body"`
	CommentCount int                `json:"comments"`
	Assignee     GitHubUser         `json:"assignee"`
	Assignees    []GitHubUser       `json:"assignees"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
	ClosedAt     string             `json:"closed_at"`
	Labels       []GitHubIssueLabel `json:"labels"`
}

type GitHubIssueLabel struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
