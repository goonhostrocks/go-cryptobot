package webhook

import gocryptobot "oss.goonhost.rocks/go-cryptobot"

type Update struct {
	UpdateID    int                 `json:"update_id"`    // Unique ID for this update.
	UpdateType  string              `json:"update_type"`  // Type of the update, can be “invoice”, “transfer” or “check”.
	RequestDate string              `json:"request_date"` // Date the update was sent in ISO 8601 format.
	Payload     gocryptobot.Invoice `json:"payload"`      // Payload of the update, can be an invoice, transfer or check.
}
