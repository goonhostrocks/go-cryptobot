package gocryptobot

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (r GetInvoicesRequest) MarshalJSON() ([]byte, error) {
	// Define an alias type to avoid infinite recursion loops during marshaling
	type Alias GetInvoicesRequest

	// Create a temporary struct that replaces []int with a string representation
	var idsStr string
	if len(r.InvoiceIDs) > 0 {
		strIDs := make([]string, len(r.InvoiceIDs))
		for i, id := range r.InvoiceIDs {
			strIDs[i] = strconv.Itoa(id)
		}
		idsStr = strings.Join(strIDs, ",")
	}

	return json.Marshal(&struct {
		Alias
		InvoiceIDs string `json:"invoice_ids,omitempty"`
	}{
		Alias:      Alias(r),
		InvoiceIDs: idsStr,
	})
}

func (r GetChecksRequest) MarshalJSON() ([]byte, error) {
	type Alias GetChecksRequest

	var idsStr string
	if len(r.CheckIDs) > 0 {
		strIDs := make([]string, len(r.CheckIDs))
		for i, id := range r.CheckIDs {
			strIDs[i] = strconv.Itoa(id)
		}
		idsStr = strings.Join(strIDs, ",")
	}

	return json.Marshal(&struct {
		Alias
		CheckIDs string `json:"check_ids,omitempty"`
	}{
		Alias:    Alias(r),
		CheckIDs: idsStr,
	})
}

func (r GetTransfersRequest) MarshalJSON() ([]byte, error) {
	type Alias GetTransfersRequest

	var idsStr string
	if len(r.TransferIDs) > 0 {
		strIDs := make([]string, len(r.TransferIDs))
		for i, id := range r.TransferIDs {
			strIDs[i] = strconv.Itoa(id)
		}
		idsStr = strings.Join(strIDs, ",")
	}

	return json.Marshal(&struct {
		Alias
		TransferIDs string `json:"transfer_ids,omitempty"`
	}{
		Alias:       Alias(r),
		TransferIDs: idsStr,
	})
}
