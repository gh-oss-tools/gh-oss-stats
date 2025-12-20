package github

import "time"

// SearchIssuesResponse represents the response from GitHub's search/issues API.
// Used to find merged PRs authored by a specific user.
type SearchIssuesResponse struct {
	TotalCount        int     `json:"total_count"`
	IncompleteResults bool    `json:"incomplete_results"`
	Items             []Issue `json:"items"`
}

// Issue represents a GitHub issue or pull request from the search API.
type Issue struct {
	Number        int             `json:"number"`
	Title         string          `json:"title"`
	State         string          `json:"state"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	ClosedAt      *time.Time      `json:"closed_at"`
	PullRequest   *PullRequestRef `json:"pull_request"`
	RepositoryURL string          `json:"repository_url"`
	HTMLURL       string          `json:"html_url"`
	User          User            `json:"user"`
}

// PullRequestRef contains references to a pull request's URLs.
type PullRequestRef struct {
	URL      string     `json:"url"`
	HTMLURL  string     `json:"html_url"`
	DiffURL  string     `json:"diff_url"`
	PatchURL string     `json:"patch_url"`
	MergedAt *time.Time `json:"merged_at"`
}

// User represents a GitHub user.
type User struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
	Type  string `json:"type"`
}

// PullRequest represents a GitHub pull request with detailed information.
type PullRequest struct {
	Number       int        `json:"number"`
	State        string     `json:"state"`
	Title        string     `json:"title"`
	User         User       `json:"user"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	ClosedAt     *time.Time `json:"closed_at"`
	MergedAt     *time.Time `json:"merged_at"`
	Merged       bool       `json:"merged"`
	Commits      int        `json:"commits"`
	Additions    int        `json:"additions"`
	Deletions    int        `json:"deletions"`
	ChangedFiles int        `json:"changed_files"`
	HTMLURL      string     `json:"html_url"`
}

// Repository represents a GitHub repository with metadata.
type Repository struct {
	Name            string     `json:"name"`
	FullName        string     `json:"full_name"`
	Owner           User       `json:"owner"`
	Description     string     `json:"description"`
	HTMLURL         string     `json:"html_url"`
	Fork            bool       `json:"fork"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	PushedAt        *time.Time `json:"pushed_at"`
	StargazersCount int        `json:"stargazers_count"`
	Language        string     `json:"language"`
	ForksCount      int        `json:"forks_count"`
	OpenIssuesCount int        `json:"open_issues_count"`
	DefaultBranch   string     `json:"default_branch"`
}

// RateLimitResponse represents the rate limit information from GitHub's API.
type RateLimitResponse struct {
	Resources RateLimitResources `json:"resources"`
}

// RateLimitResources contains rate limit information for different API resources.
type RateLimitResources struct {
	Core   RateLimit `json:"core"`
	Search RateLimit `json:"search"`
}

// RateLimit represents rate limit information for a specific resource.
type RateLimit struct {
	Limit     int   `json:"limit"`
	Remaining int   `json:"remaining"`
	Reset     int64 `json:"reset"` // Unix timestamp
	Used      int   `json:"used"`
}
