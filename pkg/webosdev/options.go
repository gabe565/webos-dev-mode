package webosdev

import (
	"crypto/tls"
	"net/http"
	"time"
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
		if baseURL == "" {
			c.baseURL = DefaultBaseURL
		} else {
			c.baseURL = baseURL
		}
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
		c.userAgent = userAgent
	}
}

// WithInsecureSkipVerify returns an Option that toggles TLS verification for HTTP requests made by a Client.
func WithInsecureSkipVerify(insecureSkipVerify bool) Option {
	return func(c *Client) {
		if c.client.Transport == nil {
			c.client.Transport = http.DefaultTransport.(*http.Transport).Clone()
		}
		//nolint:gosec
		c.client.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecureSkipVerify}
	}
}
