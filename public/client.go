package public

import (
	"net/http"
	"time"
)

type Client struct {
	host   string
	client *http.Client
}

func New(host string, timeout time.Duration) *Client {
	return &Client{
		host: host,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}
