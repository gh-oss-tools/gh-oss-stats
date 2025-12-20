package ossstats

import (
	"net/http"
	"time"
)

// Client represents a GitHub OSS stats client.
// It is safe for concurrent use by multiple goroutines.
type Client struct {
	// Authentication
	token string

	// Configuration options
	includeLOC       bool
	includePRDetails bool
	minStars         int
	maxPRs           int
	timeout          time.Duration

	// HTTP client
	httpClient *http.Client

	// Logger
	logger Logger
}

// New creates a new Client with the provided options.
// The client is configured with sensible defaults that can be overridden
// using functional options.
//
// Example:
//
//	client := ossstats.New(
//	    ossstats.WithToken(token),
//	    ossstats.WithMinStars(100),
//	    ossstats.WithVerbose(),
//	)
func New(opts ...Option) *Client {
	// Create client with default values
	client := &Client{
		includeLOC:       true,            // Include LOC metrics by default
		includePRDetails: false,           // Don't include PR details by default
		minStars:         0,               // No star filtering by default
		maxPRs:           500,             // Limit to 500 PRs by default
		timeout:          5 * time.Minute, // 5 minute timeout by default
		httpClient:       &http.Client{},  // Default HTTP client
		logger:           defaultLogger{}, // No-op logger by default
	}

	// Apply all provided options
	for _, opt := range opts {
		opt(client)
	}

	// Configure HTTP client timeout if not already set
	if client.httpClient.Timeout == 0 {
		client.httpClient.Timeout = client.timeout
	}

	return client
}
