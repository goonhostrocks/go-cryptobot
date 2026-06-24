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

// GetMe tests your app's authentication token and returns basic information about the app.
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

// CreateInvoice creates a new invoice for a user to pay.
func (c *Client) CreateInvoice(invoice *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	var response CreateInvoiceResponse
	err := c.sendPost("/createInvoice", invoice, &response)
	return &response, err
}

// DeleteInvoice deletes an invoice created by your app.
func (c *Client) DeleteInvoice(invoice *DeleteInvoiceRequest) (*DeleteInvoiceResponse, error) {
	var response DeleteInvoiceResponse
	err := c.sendPost("/deleteInvoice", invoice, &response)
	return &response, err
}

// CreateCheck creates a new check.
func (c *Client) CreateCheck(check *CreateCheckRequest) (*CreateCheckResponse, error) {
	var response CreateCheckResponse
	err := c.sendPost("/createCheck", check, &response)
	return &response, err
}

// DeleteCheck deletes a check created by your app.
func (c *Client) DeleteCheck(check *DeleteCheckRequest) (*DeleteCheckResponse, error) {
	var response DeleteCheckResponse
	err := c.sendPost("/deleteCheck", check, &response)
	return &response, err
}

// Transfer sends coins from your app's balance to a user.
// This method must first be enabled in the security settings of your app.
func (c *Client) Transfer(transfer *TransferRequest) (*TransferResponse, error) {
	var response TransferResponse
	err := c.sendPost("/transfer", transfer, &response)
	return &response, err
}

// GetInvoices retrieves invoices created by your app.
func (c *Client) GetInvoices(invoices *GetInvoicesRequest) (*GetInvoicesResponse, error) {
	var response GetInvoicesResponse
	err := c.sendPost("/getInvoices", invoices, &response)
	return &response, err
}

// GetChecks retrieves checks created by your app.
func (c *Client) GetChecks(checks *GetChecksRequest) (*GetChecksResponse, error) {
	var response GetChecksResponse
	err := c.sendPost("/getChecks", checks, &response)
	return &response, err
}

// GetTransfers retrieves transfers created by your app.
func (c *Client) GetTransfers(transfers *GetTransfersRequest) (*GetTransfersResponse, error) {
	var response GetTransfersResponse
	err := c.sendPost("/getTransfers", transfers, &response)
	return &response, err
}

// GetBalance retrieves the balances of your app.
func (c *Client) GetBalance() (*GetBalanceResponse, error) {
	var response GetBalanceResponse
	err := c.sendPost("/getBalance", nil, &response)
	return &response, err
}

// GetExchangeRates retrieves exchange rates of supported currencies.
func (c *Client) GetExchangeRates() (*GetExchangeRatesResponse, error) {
	var response GetExchangeRatesResponse
	err := c.sendPost("/getExchangeRates", nil, &response)
	return &response, err
}

// GetCurrencies retrieves a list of supported fiat and cryptocurrencies.
func (c *Client) GetCurrencies() (*GetCurrenciesResponse, error) {
	var response GetCurrenciesResponse
	err := c.sendPost("/getCurrencies", nil, &response)
	return &response, err
}

// GetStats retrieves app statistics.
func (c *Client) GetStats(stats *GetStatsRequest) (*GetStatsResponse, error) {
	var response GetStatsResponse
	err := c.sendPost("/getStats", stats, &response)
	return &response, err
}
