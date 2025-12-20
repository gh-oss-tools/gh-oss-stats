package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	// GitHubAPIBaseURL is the base URL for GitHub's REST API v3
	GitHubAPIBaseURL = "https://api.github.com"

	// APIVersion is the GitHub API version header value
	APIVersion = "2022-11-28"
)

// APIClient is a low-level GitHub API client.
type APIClient struct {
	httpClient *http.Client
	token      string
	baseURL    string
}

// NewAPIClient creates a new GitHub API client.
func NewAPIClient(httpClient *http.Client, token string) *APIClient {
	return &APIClient{
		httpClient: httpClient,
		token:      token,
		baseURL:    GitHubAPIBaseURL,
	}
}

// doRequest performs an HTTP request with proper authentication and headers.
func (c *APIClient) doRequest(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	url := c.baseURL + path

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Set required headers
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", APIVersion)

	// Add authentication if token is provided
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	return resp, nil
}

// get performs a GET request and decodes the JSON response.
func (c *APIClient) get(ctx context.Context, path string, result interface{}) (*http.Response, error) {
	resp, err := c.doRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return resp, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	// Decode JSON response
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return resp, fmt.Errorf("decoding response: %w", err)
		}
	}

	return resp, nil
}

// SearchIssues searches for issues/PRs matching the given query.
func (c *APIClient) SearchIssues(ctx context.Context, query string, page, perPage int) (*SearchIssuesResponse, *http.Response, error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("page", fmt.Sprintf("%d", page))
	params.Set("per_page", fmt.Sprintf("%d", perPage))
	params.Set("sort", "updated")
	params.Set("order", "desc")

	path := "/search/issues?" + params.Encode()

	var result SearchIssuesResponse
	resp, err := c.get(ctx, path, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPullRequest fetches detailed information about a pull request.
func (c *APIClient) GetPullRequest(ctx context.Context, owner, repo string, number int) (*PullRequest, *http.Response, error) {
	path := fmt.Sprintf("/repos/%s/%s/pulls/%d", owner, repo, number)

	var result PullRequest
	resp, err := c.get(ctx, path, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetRepository fetches information about a repository.
func (c *APIClient) GetRepository(ctx context.Context, owner, repo string) (*Repository, *http.Response, error) {
	path := fmt.Sprintf("/repos/%s/%s", owner, repo)

	var result Repository
	resp, err := c.get(ctx, path, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetRateLimit fetches the current rate limit status.
func (c *APIClient) GetRateLimit(ctx context.Context) (*RateLimitResponse, error) {
	path := "/rate_limit"

	var result RateLimitResponse
	_, err := c.get(ctx, path, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ParseRepoURL extracts owner and repo name from a repository URL.
// Supports URLs like "https://api.github.com/repos/owner/repo"
func ParseRepoURL(repoURL string) (owner, repo string, err error) {
	// Remove trailing slash
	repoURL = strings.TrimSuffix(repoURL, "/")

	// Split by slashes
	parts := strings.Split(repoURL, "/")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid repository URL: %s", repoURL)
	}

	// Get last two parts (owner/repo)
	repo = parts[len(parts)-1]
	owner = parts[len(parts)-2]

	return owner, repo, nil
}

// ParseLinkHeader extracts pagination URLs from the Link header.
// Returns a map of rel -> URL (e.g., "next" -> "https://...")
func ParseLinkHeader(linkHeader string) map[string]string {
	links := make(map[string]string)

	if linkHeader == "" {
		return links
	}

	// Split by comma to get individual links
	parts := strings.Split(linkHeader, ",")
	for _, part := range parts {
		// Each part looks like: <https://...>; rel="next"
		sections := strings.Split(strings.TrimSpace(part), ";")
		if len(sections) != 2 {
			continue
		}

		// Extract URL (remove < and >)
		urlPart := strings.TrimSpace(sections[0])
		urlPart = strings.TrimPrefix(urlPart, "<")
		urlPart = strings.TrimSuffix(urlPart, ">")

		// Extract rel value
		relPart := strings.TrimSpace(sections[1])
		relPart = strings.TrimPrefix(relPart, "rel=\"")
		relPart = strings.TrimSuffix(relPart, "\"")

		links[relPart] = urlPart
	}

	return links
}
