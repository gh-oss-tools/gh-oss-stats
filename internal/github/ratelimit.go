package github

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

const (
	// RateLimitRemainingHeader is the header containing remaining requests
	RateLimitRemainingHeader = "X-RateLimit-Remaining"

	// RateLimitResetHeader is the header containing the rate limit reset time
	RateLimitResetHeader = "X-RateLimit-Reset"

	// SearchAPIDelay is the delay between search API calls (30 requests/minute)
	SearchAPIDelay = 2 * time.Second

	// MaxBackoffAttempts is the maximum number of retry attempts for rate limiting
	MaxBackoffAttempts = 5

	// InitialBackoffDelay is the initial delay for exponential backoff
	InitialBackoffDelay = 1 * time.Second
)

// RateLimitInfo contains rate limit information from response headers.
type RateLimitInfo struct {
	Remaining int
	Reset     time.Time
}

// ParseRateLimitHeaders extracts rate limit information from HTTP response headers.
func ParseRateLimitHeaders(headers http.Header) (*RateLimitInfo, error) {
	remainingStr := headers.Get(RateLimitRemainingHeader)
	resetStr := headers.Get(RateLimitResetHeader)

	if remainingStr == "" || resetStr == "" {
		return nil, fmt.Errorf("rate limit headers not found")
	}

	remaining, err := strconv.Atoi(remainingStr)
	if err != nil {
		return nil, fmt.Errorf("parsing remaining: %w", err)
	}

	resetUnix, err := strconv.ParseInt(resetStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parsing reset: %w", err)
	}

	return &RateLimitInfo{
		Remaining: remaining,
		Reset:     time.Unix(resetUnix, 0),
	}, nil
}

// IsRateLimited checks if a response indicates rate limiting.
func IsRateLimited(resp *http.Response) bool {
	return resp.StatusCode == http.StatusTooManyRequests || // 429
		resp.StatusCode == http.StatusForbidden // 403 can also indicate rate limiting
}

// HandleRateLimit implements exponential backoff for rate-limited requests.
// Returns an error if max attempts are reached or context is cancelled.
func HandleRateLimit(ctx context.Context, resp *http.Response, attempt int) error {
	if attempt >= MaxBackoffAttempts {
		return fmt.Errorf("max retry attempts (%d) reached for rate limiting", MaxBackoffAttempts)
	}

	// Try to get rate limit info from headers
	info, err := ParseRateLimitHeaders(resp.Header)
	if err == nil && info.Remaining == 0 {
		// Wait until reset time
		waitDuration := time.Until(info.Reset)
		if waitDuration > 0 {
			// Add a small buffer
			waitDuration += 5 * time.Second

			select {
			case <-time.After(waitDuration):
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	// Fall back to exponential backoff
	backoffDelay := calculateBackoff(attempt)

	select {
	case <-time.After(backoffDelay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// calculateBackoff calculates the exponential backoff delay.
// Formula: InitialBackoffDelay * 2^attempt
func calculateBackoff(attempt int) time.Duration {
	delay := float64(InitialBackoffDelay) * math.Pow(2, float64(attempt))
	return time.Duration(delay)
}

// WaitForSearchAPI implements the required delay between search API calls.
// GitHub's search API has stricter limits (30 requests/minute).
func WaitForSearchAPI(ctx context.Context) error {
	select {
	case <-time.After(SearchAPIDelay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CheckRateLimit checks if we're approaching rate limits and logs a warning.
func CheckRateLimit(info *RateLimitInfo, threshold int) bool {
	return info != nil && info.Remaining <= threshold
}

// GetResetTime returns a human-readable reset time.
func GetResetTime(info *RateLimitInfo) string {
	if info == nil {
		return "unknown"
	}
	return info.Reset.Format(time.RFC3339)
}

// ShouldRetry determines if a request should be retried based on the response.
func ShouldRetry(resp *http.Response) bool {
	if resp == nil {
		return false
	}

	// Retry on rate limiting
	if IsRateLimited(resp) {
		return true
	}

	// Retry on temporary server errors
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		return true
	}

	return false
}
