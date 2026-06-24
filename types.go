package gocryptobot

// Invoice represents a Crypto Pay invoice.
type Invoice struct {
	InvoiceID      int    `json:"invoice_id"`      // Unique ID for this invoice.
	Hash           string `json:"hash"`            // Hash of the invoice.
	CurrencyType   string `json:"currency_type"`   // Type of the price, can be “crypto” or “fiat”.
	Asset          string `json:"asset"`           // Optional. Cryptocurrency code. Available only if the value of the field currency_type is “crypto”. Currently, can be “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet).
	Fiat           string `json:"fiat"`            // ​Optional. Fiat currency code. Available only if the value of the field currency_type is “fiat”. Currently one of “USD”, “EUR”, “RUB”, “BYN”, “UAH”, “GBP”, “CNY”, “KZT”, “UZS”, “GEL”, “TRY”, “AMD”, “THB”, “INR”, “BRL”, “IDR”, “AZN”, “AED”, “PLN” and “ILS". Lol, fiat, probably the dev has a rusty old 500.
	Amount         string `json:"amount"`          // Amount of the invoice for which the invoice was created.
	PaidAsset      string `json:"paid_asset"`      // ​Optional. Cryptocurrency alphabetic code for which the invoice was paid. Available only if currency_type is “fiat” and status is “paid”.
	PaidAmount     string `json:"paid_amount"`     // ​Optional. Amount of the invoice for which the invoice was paid. Available only if currency_type is “fiat” and status is “paid”.
	PaidFiatRate   string `json:"paid_fiat_rate"`  // Optional. The rate of the paid_asset valued in the fiat currency. Available only if the value of the field currency_type is “fiat” and the value of the field status is “paid”.
	AcceptedAssets []string `json:"accepted_assets"` // Optional. List of assets which can be used to pay the invoice. Available only if currency_type is “fiat”. Currently, can be “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet).
	FeeAsset       string   `json:"fee_asset"`       // Optional. Asset of service fees charged when the invoice was paid. Available only if status is “paid”.
	FeeAmount      string   `json:"fee_amount"`      // Optional. Amount of service fees charged when the invoice was paid. Available only if status is “paid”.

	// Deprecated: Use FeeAmount or FeeAsset instead.
	// This field is only available in webhook payloads.
	Fee string `json:"fee"`

	// Deprecated: URL should be provided to the user to pay the invoice. Use BotInvoiceURL instead.
	PayURL string `json:"pay_url"`

	BotInvoiceURL     string `json:"bot_invoice_url"`      // URL should be provided to the user to pay the invoice.
	MiniAppInvoiceURL string `json:"mini_app_invoice_url"` // Use this URL to pay an invoice to the Telegram Mini App version.
	WebAppInvoiceURL  string `json:"webapp_invoice_url"`   // Use this URL to pay an invoice to the Web version of Crypto Bot.
	Description       string `json:"description"`          // Optional. Description for this invoice.
	Status            string `json:"status"`               // Status of the transfer, can be “active”, “paid” or “expired”.
	SwapTo            string `json:"swap_to"`              // Optional. The asset that will be attempted to be swapped into after the user makes a payment (the swap is not guaranteed). Supported assets: "USDT", "TON", "TRX", "ETH", "SOL", "BTC", "LTC".
	IsSwapped         bool   `json:"is_swapped"`           // ​Optional. For invoices with the "paid" status, this flag indicates whether the swap was successful (only applicable if swap_to is set).
	SwappedUID        string `json:"swapped_uid"`          // Optional. If is_swapped is true, stores the unique identifier of the swap.
	SwappedTo         string `json:"swapped_to"`           // ​Optional. If is_swapped is true, stores the asset into which the swap was made.
	SwappedRate       string `json:"swapped_rate"`         // ​Optional. If is_swapped is true, stores the exchange rate at which the swap was executed.
	SwappedOutput     string `json:"swapped_output"`       // Optional. If is_swapped is true, stores the amount received as a result of the swap (in the swapped_to asset).
	SwappedUSDAmount  string `json:"swapped_usd_amount"`   // Optional. If is_swapped is true, stores the resulting swap amount in USD.
	SwappedUSDRate    string `json:"swapped_usd_rate"`     // ​Optional. If is_swapped is true, stores the USD exchange rate of the currency from swapped_to.
	CreatedAt         string `json:"created_at"`           // Date the invoice was created in ISO 8601 format.
	PaidUSDRate       string `json:"paid_usd_rate"`        // Optional. Price of the asset in USD. Available only if status is “paid”.

	// Deprecated: Use PaidUSDRate instead. Optional. Price of the asset in USD. Available only in the Webhook update payload.
	USDRate string `json:"usd_rate"`

	AllowComments   bool   `json:"allow_comments"`   // True, if the user can add comment to the payment.
	AllowAnonymous  bool   `json:"allow_anonymous"`  // ​True, if the user can pay the invoice anonymously.
	ExpirationDate  string `json:"expiration_date"`  // ​Optional. Date the invoice expires in ISO 8601 format.
	PaidAt          string `json:"paid_at"`          // ​Optional. Date the invoice was paid in ISO 8601 format.
	PaidAnonymously bool   `json:"paid_anonymously"` // True, if the invoice was paid anonymously.
	Comment         string `json:"comment"`          // Optional. Comment to the payment from the user.
	HiddenMessage   string `json:"hidden_message"`   // Optional. Text of the hidden message for this invoice.
	Payload         string `json:"payload"`          // ​Optional. Previously provided data for this invoice.
	PaidBtnName     string `json:"paid_btn_name"`    // Optional. Label of the button, can be “viewItem”, “openChannel”, “openBot” or “callback”.
	PaidBtnURL      string `json:"paid_btn_url"`     // ​Optional. URL opened using the button.
}

// Transfer represents a Crypto Pay transfer.
type Transfer struct {
	TransferID  int    `json:"transfer_id"`  // Unique ID for this transfer.
	SpendID     string `json:"spend_id"`     // Unique UTF-8 string.
	UserID      int    `json:"user_id"`      // Telegram user ID the transfer was sent to.
	Asset       string `json:"asset"`        // Cryptocurrency alphabetic code. Currently, can be “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet).
	Amount      string `json:"amount"`       // Amount of the transfer in float.
	Status      string `json:"status"`       // Status of the transfer, can only be “completed”.
	CompletedAt string `json:"completed_at"` // Date the transfer was completed in ISO 8601 format.
	Comment     string `json:"comment"`      // ​Optional. Comment for this transfer.
}

// Check represents a Crypto Pay check.
type Check struct {
	CheckID     int    `json:"check_id"`      // Unique ID for this check.
	Hash        string `json:"hash"`          // Hash of the check.
	Asset       string `json:"asset"`         // Cryptocurrency alphabetic code. Currently, can be “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet).
	Amount      string `json:"amount"`        // Amount of the check in float.
	BotCheckURL string `json:"bot_check_url"` // URL should be provided to the user to activate the check.
	Status      string `json:"status"`        // Status of the check, can be “active” or “activated”.
	CreatedAt   string `json:"created_at"`    // Date the check was created in ISO 8601 format.
	ActivatedAt string `json:"activated_at"`  // Date the check was activated in ISO 8601 format.
}

// Balance represents an app balance.
type Balance struct {
	CurrencyCode string `json:"currency_code"` // Cryptocurrency alphabetic code. Currently, can be “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet).
	Available    string `json:"available"`     // Total available amount in float.
	OnHold       string `json:"on_hold"`       // Unavailable amount currently is on hold in float.
}

// ExchangeRate represents an exchange rate for a currency.
type ExchangeRate struct {
	IsValid  bool   `json:"is_valid"`  // ​True, if the received rate is up-to-date.
	IsCrypto bool   `json:"is_crypto"` // True, if the source is the cryptocurrency.
	IsFiat   bool   `json:"is_fiat"`   // ​True, if the source is the fiat currency.
	Source   string `json:"source"`    // Cryptocurrency alphabetic code. Currently, can be “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC”.
	Target   string `json:"target"`    // Fiat currency code. Currently, can be “USD”, “EUR”, “RUB”, “BYN”, “UAH”, “GBP”, “CNY”, “KZT”, “UZS”, “GEL”, “TRY”, “AMD”, “THB”, “INR”, “BRL”, “IDR”, “AZN”, “AED”, “PLN” and “ILS".
	Rate     string `json:"rate"`      // The current rate of the source asset valued in the target currency.
}

// AppStats represents app usage statistics.
type AppStats struct {
	Volume              string  `json:"volume"`                // Total volume of paid invoices in USD.
	Conversion          string  `json:"conversion"`            // Conversion of all created invoices.
	UniqueUsersCount    int     `json:"unique_users_count"`    // The unique number of users who have paid the invoice.
	CreatedInvoiceCount int     `json:"created_invoice_count"` // Total created invoice count.
	PaidInvoiceCount    int     `json:"paid_invoice_count"`    // Total paid invoice count.
	StartAt             string  `json:"start_at"`              // The date on which the statistics calculation was started in ISO 8601 format.
	EndAt               string  `json:"end_at"`                // The date on which the statistics calculation was ended in ISO 8601 format.
}

// CreateInvoiceRequest contains parameters for creating an invoice.
type CreateInvoiceRequest struct {
	CurrencyType   string `json:"currency_type,omitempty"`   // Optional. "crypto" or "fiat". Defaults to crypto.
	Asset          string `json:"asset,omitempty"`           // Optional. Required if currency_type is "crypto".
	Fiat           string `json:"fiat,omitempty"`            // Optional. Required if currency_type is "fiat".
	AcceptedAssets string `json:"accepted_assets,omitempty"` // Optional. Comma-separated list of assets.
	Amount         string `json:"amount"`                    // Required. Amount of the invoice in float string.
	SwapTo         string `json:"swap_to,omitempty"`         // Optional. Asset to attempt to swap into.
	Description    string `json:"description,omitempty"`     // Optional. Up to 1024 characters.
	HiddenMessage  string `json:"hidden_message,omitempty"`  // Optional. Up to 2048 characters.
	PaidBtnName    string `json:"paid_btn_name,omitempty"`   // Optional. viewItem, openChannel, openBot, callback.
	PaidBtnURL     string `json:"paid_btn_url,omitempty"`    // Optional. Required if paid_btn_name is specified.
	Payload        string `json:"payload,omitempty"`         // Optional. Internal metadata up to 4kb.
	AllowComments  *bool  `json:"allow_comments,omitempty"`  // Optional. Use pointer to handle explicit false vs omitempty.
	AllowAnonymous *bool  `json:"allow_anonymous,omitempty"` // Optional. Use pointer to handle explicit false vs omitempty.
	ExpiresIn      int    `json:"expires_in,omitempty"`      // Optional. Payment time limit in seconds.
}

// APIError represents an error returned by the Crypto Pay API.
type APIError struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

// DeleteInvoiceRequest contains parameters for deleting an invoice.
type DeleteInvoiceRequest struct {
	InvoiceID int `json:"invoice_id"` // Invoice ID to be deleted.
}

// CreateCheckRequest contains parameters for creating a check.
type CreateCheckRequest struct {
	Asset         string `json:"asset"`                     // Cryptocurrency alphabetic code. Supported assets: “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet).
	Amount        string `json:"amount"`                    // Amount of the check in float. For example: 125.50
	PinToUserID   int    `json:"pin_to_user_id,omitempty"`  // ​Optional. ID of the user who will be able to activate the check.
	PinToUsername string `json:"pin_to_username,omitempty"` // Optional. A user with the specified username will be able to activate the check.
}

// DeleteCheckRequest contains parameters for deleting a check.
type DeleteCheckRequest struct {
	CheckID int `json:"check_id"` // Check ID to be deleted.
}

// TransferRequest contains parameters for transferring coins to a user.
type TransferRequest struct {
	UserID                  int    `json:"user_id"`                             // User ID in Telegram. User must have previously used @CryptoBot (@CryptoTestnetBot for testnet).
	Asset                   string `json:"asset"`                               // Cryptocurrency alphabetic code. Supported assets: “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet).
	Amount                  string `json:"amount"`                              // Amount of the transfer in float. The minimum and maximum amount limits for each of the supported assets roughly correspond to 1-25000 USD. Use getExchangeRates to convert amounts. For example: 125.50
	SpendID                 string `json:"spend_id"`                            // Random UTF-8 string unique per transfer for idempotent requests. The same spend_id can be accepted only once from your app. Up to 64 symbols.
	Comment                 string `json:"comment,omitempty"`                   // Optional. Comment for the transfer. Users will see this comment in the notification about the transfer. Up to 1024 symbols.
	DisableSendNotification *bool  `json:"disable_send_notification,omitempty"` // Optional. Pass true to not send to the user the notification about the transfer. Defaults to false.
}

// GetInvoicesRequest contains parameters for retrieving invoices.
type GetInvoicesRequest struct {
	Asset      string `json:"asset,omitempty"`  // Optional. Cryptocurrency alphabetic code. Supported assets: “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet). Defaults to all currencies.
	Fiat       string `json:"fiat,omitempty"`   // ​Optional. Fiat currency code. Supported fiat currencies: “USD”, “EUR”, “RUB”, “BYN”, “UAH”, “GBP”, “CNY”, “KZT”, “UZS”, “GEL”, “TRY”, “AMD”, “THB”, “INR”, “BRL”, “IDR”, “AZN”, “AED”, “PLN” and “ILS". Defaults to all currencies.
	InvoiceIDs []int  `json:"-"`                // ​Optional. List of invoice IDs separated by comma.
	Status     string `json:"status,omitempty"` // Optional. Status of invoices to be returned. Available statuses: “active” and “paid”. Defaults to all statuses.
	Offset     *int   `json:"offset,omitempty"` // ​Optional. Offset needed to return a specific subset of invoices. Defaults to 0.
	Count      *int   `json:"count,omitempty"`  // ​Optional. Number of invoices to be returned. Values between 1-1000 are accepted. Defaults to 100.
}

// CreateInvoiceResponse is the response from CreateInvoice.
type CreateInvoiceResponse struct {
	Ok     bool    `json:"ok"`
	Result Invoice `json:"result"`
}

type DeleteInvoiceResponse struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

type CreateCheckResponse struct {
	Ok     bool  `json:"ok"`
	Result Check `json:"result"`
}

type DeleteCheckResponse struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

type TransferResponse struct {
	Ok     bool     `json:"ok"`
	Result Transfer `json:"result"`
}

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Ok    bool     `json:"ok"`
	Error APIError `json:"error"`
}

type GetMeResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		AppID                        int    `json:"app_id"`
		Name                         string `json:"name"`
		PaymentProcessingBotUsername string `json:"payment_processing_bot_username"`
	} `json:"result"`
}

type GetInvoicesResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		Items []Invoice `json:"items"`
	} `json:"result"`
}

type GetChecksRequest struct {
	Asset    string `json:"asset,omitempty"`  // Optional. Cryptocurrency alphabetic code. Supported assets: “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet). Defaults to all currencies.
	CheckIDs []int  `json:"-"`                // Optional. List of check IDs separated by comma.
	Status   string `json:"status,omitempty"` // Optional. Status of check to be returned. Available statuses: “active” and “activated”. Defaults to all statuses.
	Offset   *int   `json:"offset,omitempty"` // Optional. Offset needed to return a specific subset of check. Defaults to 0.
	Count    *int   `json:"count,omitempty"`  // Optional. Number of check to be returned. Values between 1-1000 are accepted. Defaults to 100.
}

type GetChecksResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		Items []Check `json:"items"`
	} `json:"result"`
}

type GetTransfersRequest struct {
	Asset       string `json:"asset,omitempty"`    // Optional. Cryptocurrency alphabetic code. Supported assets: “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet). Defaults to all currencies.
	TransferIDs []int  `json:"-"`                  // Optional. List of transfer IDs separated by comma.
	SpendID     string `json:"spend_id,omitempty"` // Optional. Unique UTF-8 transfer string.
	Offset      *int   `json:"offset,omitempty"`   // Optional. Offset needed to return a specific subset of transfers. Defaults to 0.
	Count       *int   `json:"count,omitempty"`    // Optional. Number of transfers to be returned. Values between 1-1000 are accepted. Defaults to 100.
}

type GetTransfersResponse struct {
	Ok     bool `json:"ok"`
	Result struct {
		Items []Transfer `json:"items"`
	} `json:"result"`
}

type GetBalanceResponse struct {
	Ok     bool      `json:"ok"`
	Result []Balance `json:"result"`
}

type GetExchangeRatesResponse struct {
	Ok     bool           `json:"ok"`
	Result []ExchangeRate `json:"result"`
}

type Currency struct {
	IsBlockchain bool   `json:"is_blockchain"`
	IsStablecoin bool   `json:"is_stablecoin"`
	IsFiat       bool   `json:"is_fiat"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	URL          string `json:"url,omitempty"`
	Decimals     int    `json:"decimals"`
}

type GetCurrenciesResponse struct {
	Ok     bool       `json:"ok"`
	Result []Currency `json:"result"`
}

type GetStatsRequest struct {
	StartAt string `json:"start_at,omitempty"` // Optional. Date from which start calculating statistics in ISO 8601 format. Defaults is current date minus 24 hours.
	EndAt   string `json:"end_at,omitempty"`   // Optional. The date on which to finish calculating statistics in ISO 8601 format. Defaults is current date.
}

type GetStatsResponse struct {
	Ok     bool     `json:"ok"`
	Result AppStats `json:"result"`
}
