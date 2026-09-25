package heleket

import (
	"encoding/json"
	"errors"
)

const (
	resendWebhookEndpoint      = "/payment/resend"
	testPaymentWebhookEndpoint = "/test-webhook/payment"
	testPayoutWebhookEndpoint  = "/test-webhook/payout"
	testWalletWebhookEndpoint  = "/test-webhook/wallet"
)

type WebhookConvert struct {
	ToCurrency string `json:"to_currency"`
	Commission string `json:"commission"`
	Rate       string `json:"rate"`
	Amount     string `json:"amount"`
}

type Webhook struct {
	Type              string          `json:"type"`
	UUID              string          `json:"uuid"`
	OrderId           string          `json:"order_id"`
	Amount            string          `json:"amount"`
	PaymentAmount     string          `json:"payment_amount"`
	PaymentAmountUSD  string          `json:"payment_amount_usd"`
	MerchantAmount    string          `json:"merchant_amount"`
	Commission        string          `json:"commission"`
	IsFinal           bool            `json:"is_final"`
	Status            string          `json:"status"`
	From              string          `json:"from"`
	WalletAddressUUID string          `json:"wallet_address_uuid"`
	Network           string          `json:"network"`
	Currency          string          `json:"currency"`
	PayerCurrency     string          `json:"payer_currency"`
	AdditionalData    string          `json:"additional_data"`
	Convert           *WebhookConvert `json:"convert"`
	TxId              string          `json:"txid"`
	Sign              string          `json:"sign"`
}

type ResendWebhookRequest struct {
	PaymentUUID string `json:"uuid,omitempty"`
	OrderId     string `json:"order_id,omitempty"`
}

type resendWebhookRawResponse struct {
	Result []string `json:"result"`
	State  int8     `json:"state"`
}

type TestWebhookRequest struct {
	UrlCallback string `json:"url_callback"`
	Currency    string `json:"currency"`
	Network     string `json:"network"`
	UUID        string `json:"uuid,omitempty"`
	OrderId     string `json:"order_id,omitempty"`
	Status      string `json:"status"`
}

type TestWebhookResponse struct {
	Result []string `json:"result"`
	State  int8     `json:"state"`
}

// ParseWebhook parses and optionally verifies a webhook notification from Heleket.
// Webhooks are sent when payment status changes or transactions complete.
//
// Parameters:
//   - reqBody: Raw request body bytes from the webhook HTTP request
//   - verifySign: Whether to verify the webhook signature (recommended: true)
//
// Returns the parsed Webhook object. If verifySign is true, returns an error if signature is invalid.
func (c *Heleket) ParseWebhook(reqBody []byte, verifySign bool) (*Webhook, error) {
	var apiKey string
	response := &Webhook{}

	err := json.Unmarshal(reqBody, response)
	if err != nil {
		return nil, err
	}

	switch response.Type {
	case "payment", "wallet":
		apiKey = c.paymentApiKey
	case "payout":
		apiKey = c.payoutApiKey
	default:
		return nil, errors.New("unknown webhook type")
	}

	if verifySign {
		err = c.VerifySign(apiKey, reqBody)
		if err != nil {
			return nil, err
		}
	}

	return response, err
}

// ResendWebhook requests that Heleket resend a webhook notification for a specific payment.
// You must provide either PaymentUUID or OrderId in the request.
//
// Returns true if the webhook was successfully queued for resending.
func (c *Heleket) ResendWebhook(resendRequest *ResendWebhookRequest) (bool, error) {
	if resendRequest.PaymentUUID == "" && resendRequest.OrderId == "" {
		return false, errors.New("you must provide one of: [PaymentUUID, OrderId]")
	}

	res, err := c.fetch("POST", resendWebhookEndpoint, resendRequest, c.paymentApiKey)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	response := &resendWebhookRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return false, err
	}

	return len(response.Result) == 0, nil
}

// TestPaymentWebhook sends a test webhook notification to your callback URL.
// This is useful for testing your webhook handler implementation.
//
// The test webhook will contain simulated payment data with the specified currency, network, and status.
func (c *Heleket) TestPaymentWebhook(testRequest *TestWebhookRequest) (*TestWebhookResponse, error) {
	res, err := c.fetch("POST", testPaymentWebhookEndpoint, testRequest, c.paymentApiKey)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &TestWebhookResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

// TestPayoutWebhook sends a test payout webhook notification to your callback URL.
// This is useful for testing your webhook handler implementation.
//
// The test webhook will contain simulated payout data with the specified currency, network, and status.
func (c *Heleket) TestPayoutWebhook(testRequest *TestWebhookRequest) (*TestWebhookResponse, error) {
	res, err := c.fetch("POST", testPayoutWebhookEndpoint, testRequest, c.payoutApiKey)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &TestWebhookResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

// TestWalletWebhook sends a test static wallet webhook notification to your callback URL.
// This is useful for testing your webhook handler implementation.
//
// The test webhook will contain simulated wallet payment data with the specified currency, network, and status.
func (c *Heleket) TestWalletWebhook(testRequest *TestWebhookRequest) (*TestWebhookResponse, error) {
	res, err := c.fetch("POST", testWalletWebhookEndpoint, testRequest, c.paymentApiKey)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &TestWebhookResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
