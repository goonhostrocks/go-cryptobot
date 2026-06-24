// Package gocryptobot provides a Go client for the Telegram Crypto Pay API.
// It allows developers to easily integrate crypto payments, create invoices,
// transfer coins, and manage app balances.
package gocryptobot

import "net/http"

// Client represents the Crypto Pay API client.
type Client struct {
	APIToken   string
	BaseURL    string
	httpClient *http.Client
}

// NewClient initializes a new Crypto Pay API client.
// apiToken is your application token (e.g., "123456:AAzQ...").
// baseURL is the API endpoint (e.g., "https://pay.crypt.bot/api" for mainnet or "https://testnet-pay.crypt.bot/api" for testnet).
func NewClient(apiToken string, baseURL string) *Client {
	return &Client{
		APIToken:   apiToken,
		BaseURL:    baseURL,
		httpClient: &http.Client{},
	}
}
