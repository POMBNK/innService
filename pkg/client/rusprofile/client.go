package rusprofile

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	host     = "rusprofile.ru"
	basePath = "search"
)

type Client struct {
	Client *http.Client
}

func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

func (c *Client) ParsePage(inn string) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   basePath,
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("bad GET request %w", err)
	}

	q := url.Values{}
	q.Add("query", inn)
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't get response %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	page, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error on response: %w", err)
	}

	return page, nil
}
