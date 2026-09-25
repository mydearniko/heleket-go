package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseWebhookWithoutVerification(t *testing.T) {
	samplePayload := []byte(`{
		"type": "payment",
		"uuid": "test-uuid",
		"order_id": "test-order",
		"amount": "10.00",
		"payment_amount": "10.00",
		"payment_amount_usd": "10.00",
		"merchant_amount": "9.80",
		"commission": "0.20",
		"is_final": true,
		"status": "paid",
		"from": "sender-address",
		"wallet_address_uuid": "",
		"network": "tron",
		"currency": "USDT",
		"payer_currency": "USDT",
		"additional_data": "",
		"txid": "test-txid",
		"sign": "test-signature"
	}`)

	webhook, err := TestHeleket.ParseWebhook(samplePayload, false)
	require.NoError(t, err)
	require.NotNil(t, webhook)
	require.Equal(t, "payment", webhook.Type)
	require.Equal(t, "test-uuid", webhook.UUID)
	require.Equal(t, "test-order", webhook.OrderId)
}

func TestParseWebhookInvalidJSON(t *testing.T) {
	invalidPayload := []byte(`{invalid json}`)

	webhook, err := TestHeleket.ParseWebhook(invalidPayload, false)
	require.Error(t, err)
	require.Nil(t, webhook)
}

func TestParseWebhookUnknownType(t *testing.T) {
	unknownTypePayload := []byte(`{
		"type": "unknown",
		"uuid": "test-uuid",
		"sign": "test-signature"
	}`)

	webhook, err := TestHeleket.ParseWebhook(unknownTypePayload, false)
	require.Error(t, err)
	require.Nil(t, webhook)
	require.Contains(t, err.Error(), "unknown webhook type")
}
