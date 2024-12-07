package webosdev

import "time"

type Option func(c *Client)

func WithSessionToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.client.Timeout = timeout
	}
}
