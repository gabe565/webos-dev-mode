// Package webosdev provides a client for interacting with the webOS developer API.
package webosdev

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

const defaultBaseURL = "https://developer.lge.com"

// NewClient creates a Client for interacting with the webOS developer API.
// Optional configuration can be provided via Option functions.
func NewClient(opts ...Option) *Client {
	c := &Client{
		client:  &http.Client{},
		baseURL: defaultBaseURL,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Client represents a client for the webOS developer API.
type Client struct {
	client  *http.Client
	baseURL string
	token   string
}

// Response represents the structure of a webOS developer API response.
type Response struct {
	Result       string `json:"result"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMsg"`
}

// ErrRequestFailed indicates that an API request returned an error.
var ErrRequestFailed = errors.New("request failed")

func (c *Client) request(ctx context.Context, p string) (*Response, *http.Response, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, nil, err
	}
	u.Path = path.Join(u.Path, p)
	q := u.Query()
	q.Set("sessionToken", c.token)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, res, err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("%w: %s", ErrRequestFailed, res.Status)
	}

	var decoded *Response
	if err := json.NewDecoder(res.Body).Decode(&decoded); err != nil {
		return decoded, res, err
	}

	if decoded.Result != "success" {
		return decoded, res, fmt.Errorf("%w: %s", ErrRequestFailed, decoded.ErrorMessage)
	}

	return decoded, res, nil
}

// ExtendSession extends the current webOS developer mode session by making an API call.
// It returns the decoded response, the raw HTTP response, and any error encountered.
func (c *Client) ExtendSession(ctx context.Context) (*Response, *http.Response, error) {
	return c.request(ctx, "/secure/ResetDevModeSession.dev")
}

// ErrInvalidTimestamp indicates that the timestamp returned by the API could not be parsed.
var ErrInvalidTimestamp = errors.New("invalid timestamp")

// CheckExpiration checks the remaining time in the current webOS developer session.
// It parses the response timestamp and returns the remaining duration, the raw HTTP response, and any error encountered.
func (c *Client) CheckExpiration(ctx context.Context) (time.Duration, *http.Response, error) {
	decoded, res, err := c.request(ctx, "/secure/CheckDevModeSession.dev")
	if err != nil {
		return 0, res, err
	}

	parts := strings.Split(decoded.ErrorMessage, ":")
	if len(parts) != 3 {
		return 0, res, fmt.Errorf("%w: %s", ErrInvalidTimestamp, decoded.ErrorMessage)
	}

	var expiration time.Duration
	for i, part := range parts {
		v, err := strconv.Atoi(part)
		if err != nil {
			return 0, res, err
		}
		switch i {
		case 0:
			expiration += time.Duration(v) * time.Hour
		case 1:
			expiration += time.Duration(v) * time.Minute
		case 2:
			expiration += time.Duration(v) * time.Second
		}
	}

	return expiration, res, nil
}
