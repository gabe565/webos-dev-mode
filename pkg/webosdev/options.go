package webosdev

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
