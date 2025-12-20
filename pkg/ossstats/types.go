package ossstats

import (
	"fmt"
	"time"
)

// Stats represents the complete statistics for a GitHub user's
// open source contributions to external repositories.
type Stats struct {
	Username      string         `json:"username"`
	GeneratedAt   time.Time      `json:"generatedAt"`
	Summary       Summary        `json:"summary"`
	Contributions []Contribution `json:"contributions"`
}

// Summary contains aggregate statistics across all contributions.
type Summary struct {
	TotalProjects  int `json:"totalProjects"`
	TotalPRsMerged int `json:"totalPRsMerged"`
	TotalCommits   int `json:"totalCommits"`
	TotalAdditions int `json:"totalAdditions"`
	TotalDeletions int `json:"totalDeletions"`
}

// Contribution represents a user's contribution to a single external repository.
type Contribution struct {
	Repo              string    `json:"repo"`              // Full repo name (owner/repo)
	Owner             string    `json:"owner"`             // Repository owner
	RepoName          string    `json:"repoName"`          // Repository name
	Description       string    `json:"description"`       // Repository description
	RepoURL           string    `json:"repoURL"`           // Full GitHub URL
	Stars             int       `json:"stars"`             // Repository star count
	PRsMerged         int       `json:"prsMerged"`         // Number of merged PRs
	Commits           int       `json:"commits"`           // Total commits across PRs
	Additions         int       `json:"additions"`         // Lines added
	Deletions         int       `json:"deletions"`         // Lines deleted
	FirstContribution time.Time `json:"firstContribution"` // First PR merged date
	LastContribution  time.Time `json:"lastContribution"`  // Most recent PR merged date
}

// ErrRateLimited indicates that GitHub's rate limit has been exceeded.
type ErrRateLimited struct {
	ResetAt time.Time
	Message string
}

func (e *ErrRateLimited) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("rate limited: %s (resets at %s)", e.Message, e.ResetAt.Format(time.RFC3339))
	}
	return fmt.Sprintf("rate limited (resets at %s)", e.ResetAt.Format(time.RFC3339))
}

// ErrAuthentication indicates an authentication failure with the GitHub API.
type ErrAuthentication struct {
	Message string
}

func (e *ErrAuthentication) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("authentication failed: %s", e.Message)
	}
	return "authentication failed"
}

// ErrNotFound indicates that the specified GitHub user was not found.
type ErrNotFound struct {
	Username string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("user not found: %s", e.Username)
}

// ErrPartialResults indicates that the operation completed with partial results
// due to errors encountered during processing (e.g., rate limiting).
type ErrPartialResults struct {
	Stats   *Stats
	Errors  []error
	Message string
}

func (e *ErrPartialResults) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("partial results: %s (%d errors encountered)", e.Message, len(e.Errors))
	}
	return fmt.Sprintf("partial results (%d errors encountered)", len(e.Errors))
}
