package gocryptobot

import "net/http"

type Client struct {
	APIToken   string
	BaseURL    string
	httpClient *http.Client
}

func NewClient(apiToken string, baseURL string) *Client {
	return &Client{
		APIToken:   apiToken,
		BaseURL:    baseURL,
		httpClient: &http.Client{},
	}
}
