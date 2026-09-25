// Package heleket provides a Go client for the Heleket cryptocurrency payment gateway API.
//
// The Heleket API allows you to:
//   - Create payment invoices and accept cryptocurrency payments
//   - Process payouts to cryptocurrency addresses
//   - Manage static wallets for recurring payments
//   - Handle refunds for payments and blocked wallets
//   - Receive and verify webhook notifications
//   - Query balances, exchange rates, and payment history
//
// Basic usage:
//
//	import "github.com/idanyas/heleket-go"
//
//	client, err := heleket.New(
//		&http.Client{},
//		"your-merchant-id",
//		"your-payment-api-key",
//		"your-payout-api-key",
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Create a payment invoice
//	invoice, err := client.CreateInvoice(&heleket.InvoiceRequest{
//		Amount:   "10.50",
//		Currency: "USD",
//		OrderId:  "order-123",
//	})
//
// For more examples, see the examples/ directory in the repository.
package heleket

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	apiUrl     = "https://api.heleket.com/v1"
	timeFormat = "2006-01-02 15:04:05"
)

// Heleket is the main client for interacting with the Heleket API.
// It manages authentication credentials and provides methods for all API operations.
type Heleket struct {
	merchant      string
	paymentApiKey string
	payoutApiKey  string
	client        *http.Client
}

// ErrorResponse represents an error returned by the Heleket API.
// It includes the error state code and descriptive messages.
type ErrorResponse struct {
	State   int8     `json:"state"`
	Message string   `json:"message,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

func (e *ErrorResponse) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("API error (state=%d): %s", e.State, e.Message)
	}
	if len(e.Errors) > 0 {
		return fmt.Sprintf("API error (state=%d): %v", e.State, e.Errors)
	}
	return fmt.Sprintf("API error (state=%d)", e.State)
}

// New creates a new Heleket API client with the provided credentials.
// It validates that all required parameters are non-empty and returns an error if validation fails.
//
// Parameters:
//   - client: HTTP client to use for API requests (must not be nil)
//   - merchant: Your Heleket merchant ID
//   - paymentApiKey: API key for payment operations
//   - payoutApiKey: API key for payout operations
//
// Returns an error if any parameter is invalid.
func New(client *http.Client, merchant, paymentApiKey, payoutApiKey string) (*Heleket, error) {
	if client == nil {
		return nil, errors.New("http client cannot be nil")
	}
	if merchant == "" {
		return nil, errors.New("merchant ID cannot be empty")
	}
	if paymentApiKey == "" {
		return nil, errors.New("payment API key cannot be empty")
	}
	if payoutApiKey == "" {
		return nil, errors.New("payout API key cannot be empty")
	}

	return &Heleket{
		client:        client,
		merchant:      merchant,
		paymentApiKey: paymentApiKey,
		payoutApiKey:  payoutApiKey,
	}, nil
}

func (c *Heleket) fetch(method string, endpoint string, payload any, apiKey string) (*http.Response, error) {
	var body []byte
	var err error

	// For GET requests, payload should be nil. Signature is on an empty string.
	// For POST requests with no parameters, payload should be an empty map or struct, which marshals to "{}".
	if payload != nil {
		body, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request: %w", err)
		}
	} else {
		body = []byte("")
	}

	sign := c.signRequest(apiKey, body)

	var reqBody io.Reader
	if method != http.MethodGet {
		reqBody = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, apiUrl+endpoint, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("merchant", c.merchant)
	req.Header.Set("sign", sign)
	
	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Check HTTP status code
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		defer res.Body.Close()
		
		// Try to parse error response
		var errResp ErrorResponse
		if decodeErr := json.NewDecoder(res.Body).Decode(&errResp); decodeErr == nil {
			return nil, &errResp
		}
		
		// Fallback to generic error
		return nil, fmt.Errorf("API request failed with status %d", res.StatusCode)
	}

	return res, nil
}
