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

func TestGetBalance_Success(t *testing.T) {
	client := NewClient("58888:AAPYodtS0zKoIbYe7RDzKGnvMdtxTc6BPWy", "https://testnet-pay.crypt.bot/api")

	res, err := client.GetBalance()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !res.Ok {
		t.Errorf("expected response status 'ok' to be true")
	}

	t.Logf("Retrieved %d balance records", len(res.Result))
}

func TestGetExchangeRates_Success(t *testing.T) {
	client := NewClient("58888:AAPYodtS0zKoIbYe7RDzKGnvMdtxTc6BPWy", "https://testnet-pay.crypt.bot/api")

	res, err := client.GetExchangeRates()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !res.Ok {
		t.Errorf("expected response status 'ok' to be true")
	}

	t.Logf("Retrieved %d exchange rates", len(res.Result))
}

func TestGetCurrencies_Success(t *testing.T) {
	client := NewClient("58888:AAPYodtS0zKoIbYe7RDzKGnvMdtxTc6BPWy", "https://testnet-pay.crypt.bot/api")

	res, err := client.GetCurrencies()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !res.Ok {
		t.Errorf("expected response status 'ok' to be true")
	}

	t.Logf("Retrieved %d currencies", len(res.Result))
}

func TestGetStats_Success(t *testing.T) {
	client := NewClient("58888:AAPYodtS0zKoIbYe7RDzKGnvMdtxTc6BPWy", "https://testnet-pay.crypt.bot/api")

	res, err := client.GetStats(&GetStatsRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !res.Ok {
		t.Errorf("expected response status 'ok' to be true")
	}

	t.Logf("Retrieved stats, volume: %s", res.Result.Volume)
}

func TestGetInvoices_Success(t *testing.T) {
	client := NewClient("58888:AAPYodtS0zKoIbYe7RDzKGnvMdtxTc6BPWy", "https://testnet-pay.crypt.bot/api")

	// Test with no parameters
	res, err := client.GetInvoices(&GetInvoicesRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !res.Ok {
		t.Errorf("expected response status 'ok' to be true")
	}

	t.Logf("Retrieved %d invoices", len(res.Result.Items))
}

func TestCheck_Lifecycle(t *testing.T) {
	client := NewClient("58888:AAPYodtS0zKoIbYe7RDzKGnvMdtxTc6BPWy", "https://testnet-pay.crypt.bot/api")

	// CreateCheck
	res, err := client.CreateCheck(&CreateCheckRequest{
		Asset:  "TON",
		Amount: "0.1",
	})
	if err != nil {
		t.Fatalf("CreateCheck returned error: %v", err)
	}

	if !res.Ok {
		t.Logf("CreateCheck returned ok=false (expected if insufficient balance), skipping rest of check lifecycle test.")
	} else {
		checkID := res.Result.CheckID
		t.Logf("Successfully created check %d", checkID)

		// GetChecks
		resGet, err := client.GetChecks(&GetChecksRequest{
			Asset:    "TON",
			Status:   "active",
			CheckIDs: []int{checkID}, // This hits MarshalJSON
		})
		if err != nil {
			t.Fatalf("GetChecks returned error: %v", err)
		}
		if !resGet.Ok {
			t.Fatalf("GetChecks returned ok=false")
		}
		t.Logf("Successfully got %d checks", len(resGet.Result.Items))

		// DeleteCheck
		resDel, err := client.DeleteCheck(&DeleteCheckRequest{CheckID: checkID})
		if err != nil {
			t.Fatalf("DeleteCheck returned error: %v", err)
		}
		if !resDel.Ok {
			t.Fatalf("DeleteCheck returned ok=false")
		}
		t.Logf("Successfully deleted check %d", checkID)
	}
}

func TestTransfer_Live(t *testing.T) {
	client := NewClient("58888:AAPYodtS0zKoIbYe7RDzKGnvMdtxTc6BPWy", "https://testnet-pay.crypt.bot/api")

	res, err := client.Transfer(&TransferRequest{
		UserID:  1, // Invalid user ID or system ID for testnet
		Asset:   "TON",
		Amount:  "0.1",
		SpendID: "test-spend-integration",
	})
	if err != nil {
		t.Fatalf("Transfer returned error: %v", err)
	}
	if !res.Ok {
		t.Logf("Transfer returned ok=false (expected for random test ID / insufficient balance)")
	} else {
		t.Logf("Transfer returned ok=true, TransferID: %d", res.Result.TransferID)
	}

	// GetTransfers
	resGet, err := client.GetTransfers(&GetTransfersRequest{
		Asset:       "TON",
		TransferIDs: []int{1, 2}, // This hits MarshalJSON
	})
	if err != nil {
		t.Fatalf("GetTransfers returned error: %v", err)
	}
	if !resGet.Ok {
		t.Fatalf("GetTransfers returned ok=false")
	}
	t.Logf("Successfully got %d transfers", len(resGet.Result.Items))
}

func TestGetInvoices_WithIDs(t *testing.T) {
	client := NewClient("58888:AAPYodtS0zKoIbYe7RDzKGnvMdtxTc6BPWy", "https://testnet-pay.crypt.bot/api")

	// This hits MarshalJSON
	res, err := client.GetInvoices(&GetInvoicesRequest{
		Asset:      "USDT",
		InvoiceIDs: []int{1, 2},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !res.Ok {
		t.Errorf("expected response status 'ok' to be true")
	}
	t.Logf("Retrieved %d invoices via ID", len(res.Result.Items))
}
