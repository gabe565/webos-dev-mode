package webosdev

import (
	"time"

	"gabe565.com/utils/httpx"
)

// Option defines a function that configures a Client instance.
type Option func(c *Client)

// WithSessionToken returns an Option that sets the session token for a Client.
// This token is used to authenticate API requests.
func WithSessionToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

// WithBaseURL returns an Option that sets a custom base URL for a Client.
// The default base URL is "https://developer.lge.com".
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithTimeout returns an Option that sets a timeout for HTTP requests made by a Client.
// The timeout specifies the maximum duration for the request, including connection, processing, and reading the response.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.client.Timeout = timeout
	}
}

// WithUserAgent returns an Option that sets the User-Agent string to be used for HTTP requests made by a Client.
func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.client.Transport = httpx.NewUserAgentTransport(nil, userAgent)
	}
}
