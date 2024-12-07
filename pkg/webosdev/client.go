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

type Client struct {
	client  *http.Client
	baseURL string
	token   string
}

type Response struct {
	Result       string `json:"result"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMsg"`
}

var ErrRequestFailed = errors.New("request failed")

func checkResponse(res *http.Response) (*Response, *http.Response, error) {
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

func (c *Client) request(ctx context.Context, p string) (*http.Response, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, p)
	q := u.Query()
	q.Set("sessionToken", c.token)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *Client) ExtendSession(ctx context.Context) (*Response, *http.Response, error) {
	res, err := c.request(ctx, "/secure/ResetDevModeSession.dev")
	if err != nil {
		return nil, res, err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()

	return checkResponse(res)
}

var ErrInvalidTimestamp = errors.New("invalid timestamp")

func (c *Client) CheckExpiration(ctx context.Context) (time.Duration, *http.Response, error) {
	res, err := c.request(ctx, "/secure/CheckDevModeSession.dev")
	if err != nil {
		return 0, res, err
	}
	defer func() {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()
	}()

	decoded, res, err := checkResponse(res)
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
