package gocryptobot

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) sendPost(endpoint string, payload interface{}, response interface{}) error {
	var bodyBuf bytes.Buffer
	if payload != nil {
		if err := json.NewEncoder(&bodyBuf).Encode(payload); err != nil {
			return err
		}
	}

	req, err := http.NewRequest("POST", c.BaseURL+endpoint, &bodyBuf)
	if err != nil {
		return err
	}
	req.Header.Set("Crypto-Pay-API-Token", c.APIToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, response)
}

func (c *Client) GetMe() (*GetMeResponse, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/getMe", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Crypto-Pay-API-Token", c.APIToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response GetMeResponse
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) CreateInvoice(invoice *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	var response CreateInvoiceResponse
	err := c.sendPost("/createInvoice", invoice, &response)
	return &response, err
}

func (c *Client) DeleteInvoice(invoice *DeleteInvoiceRequest) (*DeleteInvoiceResponse, error) {
	var response DeleteInvoiceResponse
	err := c.sendPost("/deleteInvoice", invoice, &response)
	return &response, err
}
