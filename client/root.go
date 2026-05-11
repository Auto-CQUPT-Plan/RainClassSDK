package client

import "net/http"

type Client struct {
	client *http.Client
}

func (r *Client) setupHttpClient() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	r.client = client
}

func NewClient() *Client {
	s := &Client{}
	s.setupHttpClient()

	return s
}
