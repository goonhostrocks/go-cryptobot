package gocryptobot

import (
	"testing"
)

func TestGetMe_Success(t *testing.T) {
	client := NewClient("58888:AAPYodtS0zKoIbYe7RDzKGnvMdtxTc6BPWy", "https://testnet-pay.crypt.bot/api")

	res, err := client.GetMe()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !res.Ok {
		t.Errorf("expected response status 'ok' to be true, got false")
	}

	if res.Result.Name == "" {
		t.Errorf("expected app name to not be empty")
	}

	t.Logf("App ID: %d", res.Result.AppID)
	t.Logf("App Name: %s", res.Result.Name)
	t.Logf("Bot Username: %s", res.Result.PaymentProcessingBotUsername)
}

func TestInvoice_Lifecycle(t *testing.T) {
	// Initialize client using the live testnet
	client := NewClient("58888:AAPYodtS0zKoIbYe7RDzKGnvMdtxTc6BPWy", "https://testnet-pay.crypt.bot/api")

	var createdInvoiceID int

	// 1. TEST: CreateInvoice
	t.Run("CreateInvoice_Success", func(t *testing.T) {
		allowComments := true
		req := &CreateInvoiceRequest{
			CurrencyType:  "crypto",
			Asset:         "TON",
			Amount:        "0.1",
			Description:   "Go library integration test",
			AllowComments: &allowComments,
		}

		res, err := client.CreateInvoice(req)
		if err != nil {
			t.Fatalf("Failed to create invoice: %v", err)
		}

		if !res.Ok {
			t.Fatalf("API response marked 'ok' as false")
		}

		if res.Result.InvoiceID == 0 {
			t.Errorf("Expected a valid invoice_id, got 0")
		}

		if res.Result.Asset != "TON" {
			t.Errorf("Expected asset 'TON', got '%s'", res.Result.Asset)
		}

		if res.Result.Amount != "0.1" {
			t.Errorf("Expected amount '0.1', got '%s'", res.Result.Amount)
		}

		// Save the ID to delete it in the next test block
		createdInvoiceID = res.Result.InvoiceID
		t.Logf("Successfully created invoice with ID: %d", createdInvoiceID)
	})

	// 2. TEST: DeleteInvoice
	t.Run("DeleteInvoice_Success", func(t *testing.T) {
		if createdInvoiceID == 0 {
			t.Skip("Skipping delete test because invoice creation failed or returned 0 ID")
		}

		req := &DeleteInvoiceRequest{
			InvoiceID: createdInvoiceID,
		}

		res, err := client.DeleteInvoice(req)
		if err != nil {
			t.Fatalf("Failed to delete invoice: %v", err)
		}

		if !res.Ok {
			t.Errorf("API response marked 'ok' as false on deletion")
		}

		if !res.Result {
			t.Errorf("Expected result to be true, indicating successful deletion")
		}

		t.Logf("Successfully deleted invoice with ID: %d", createdInvoiceID)
	})
}
