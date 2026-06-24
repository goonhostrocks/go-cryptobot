package gocryptobot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_MockedMethods(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/createCheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok": true, "result": {"check_id": 1, "hash": "abc", "asset": "TON", "amount": "10", "bot_check_url": "https://t.me/CryptoBot?start=abc", "status": "active"}}`))
	})
	mux.HandleFunc("/deleteCheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok": true, "result": true}`))
	})
	mux.HandleFunc("/getChecks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok": true, "result": {"items": [{"check_id": 1, "hash": "abc", "asset": "TON", "amount": "10", "bot_check_url": "https://t.me/CryptoBot?start=abc", "status": "active"}]}}`))
	})
	mux.HandleFunc("/transfer", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok": true, "result": {"transfer_id": 1, "user_id": 123, "asset": "TON", "amount": "5", "status": "completed"}}`))
	})
	mux.HandleFunc("/getTransfers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok": true, "result": {"items": [{"transfer_id": 1, "user_id": 123, "asset": "TON", "amount": "5", "status": "completed"}]}}`))
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client := NewClient("mock-token", server.URL)

	// CreateCheck
	_, err := client.CreateCheck(&CreateCheckRequest{Asset: "TON", Amount: "10"})
	if err != nil {
		t.Fatalf("CreateCheck error: %v", err)
	}

	// DeleteCheck
	_, err = client.DeleteCheck(&DeleteCheckRequest{CheckID: 1})
	if err != nil {
		t.Fatalf("DeleteCheck error: %v", err)
	}

	// GetChecks
	_, err = client.GetChecks(&GetChecksRequest{Asset: "TON", Status: "active"})
	if err != nil {
		t.Fatalf("GetChecks error: %v", err)
	}

	// Transfer
	_, err = client.Transfer(&TransferRequest{UserID: 123, Asset: "TON", Amount: "5", SpendID: "spend1"})
	if err != nil {
		t.Fatalf("Transfer error: %v", err)
	}

	// GetTransfers
	_, err = client.GetTransfers(&GetTransfersRequest{Asset: "TON"})
	if err != nil {
		t.Fatalf("GetTransfers error: %v", err)
	}
}

func TestGetInvoicesRequest_MarshalJSON(t *testing.T) {
	req := GetInvoicesRequest{
		Asset:      "USDT",
		InvoiceIDs: []int{123, 456},
		Status:     "paid",
	}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"asset":"USDT","status":"paid","invoice_ids":"123,456"}`
	if string(b) != expected {
		t.Errorf("expected %s, got %s", expected, string(b))
	}
}

func TestGetChecksRequest_MarshalJSON(t *testing.T) {
	req := GetChecksRequest{
		Asset:    "TON",
		CheckIDs: []int{789},
	}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"asset":"TON","check_ids":"789"}`
	if string(b) != expected {
		t.Errorf("expected %s, got %s", expected, string(b))
	}
}

func TestGetTransfersRequest_MarshalJSON(t *testing.T) {
	req := GetTransfersRequest{
		Asset:       "BTC",
		TransferIDs: []int{111, 222},
	}
	b, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"asset":"BTC","transfer_ids":"111,222"}`
	if string(b) != expected {
		t.Errorf("expected %s, got %s", expected, string(b))
	}
}
